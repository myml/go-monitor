package main

import (
	"C"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/go-macaron/macaron"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	gnet "github.com/shirou/gopsutil/net"
)

var m sync.Map

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}
func start() {
	//cpu
	log.Println("start cpu")
	go func() {
		r := R{
			Name: "CPU",
		}
		for {
			p, err := cpu.Percent(0, false)
			if err != nil {
				log.Panic(err)
			}
			r.Rate = p[0]
			m.Store("cpu", r)
			//			log.Println("cpu:", p[0])
			time.Sleep(time.Second)
		}
	}()

	//memory
	log.Println("start memory")
	go func() {
		r := R{
			Name: "内存",
		}
		for {
			vmem, err := mem.VirtualMemory()
			if err != nil {
				log.Panic(err)
			}
			r.Rate = vmem.UsedPercent
			m.Store("memory", r)
			//			log.Println("memory:", m.UsedPercent)
			time.Sleep(time.Second)
		}
	}()

	//loadavg
	log.Println("start loadavg")
	go func() {
		cpuNum, err := cpu.Counts(true)
		if err != nil {
			log.Panic(err)
		}
		r := R{
			Name: "负载",
		}
		for {
			avg, err := load.Avg()
			if err != nil {
				log.Panic(err)
			}
			r.Rate = avg.Load1 / float64(cpuNum) * 100
			m.Store("loadavg", r)
			//			log.Println("loadavg:", avg.Load1/float64(cpuNum))
			time.Sleep(time.Second)
		}
	}()

	//net
	log.Println("start net")
	go func() {
		r := R{
			Name: "网络",
			Rate: -1,
		}
		var lastNetIO gnet.IOCountersStat
		for {
			netIO, err := gnet.IOCounters(false)
			if err != nil {
				log.Panic(err)
			}
			r.Up = netIO[0].BytesSent - lastNetIO.BytesSent
			r.Down = netIO[0].BytesRecv - lastNetIO.BytesRecv
			m.Store("net", r)
			//			log.Println("net.up", netIO[0].BytesSent-lastNetIO.BytesSent)
			//			log.Println("net.down", netIO[0].BytesRecv-lastNetIO.BytesRecv)
			lastNetIO = netIO[0]
			time.Sleep(time.Second)
		}
	}()

	//disk
	go func() {
		r := R{
			Name: "磁盘",
			Rate: -1,
		}
		var lastDiskIO disk.IOCountersStat
		for {
			disksIO, err := disk.IOCounters("sda", "sdb", "sdc")
			if err != nil {
				log.Panic(err)
			}
			var diskIO disk.IOCountersStat
			for name := range disksIO {
				diskIO.ReadBytes += disksIO[name].ReadBytes
				diskIO.WriteBytes += disksIO[name].WriteBytes
			}

			r.Up = diskIO.ReadBytes - lastDiskIO.ReadBytes
			r.Down = diskIO.WriteBytes - lastDiskIO.WriteBytes
			m.Store("disk", r)
			lastDiskIO = diskIO
			time.Sleep(time.Second)
		}
	}()
}

type R struct {
	Name string
	Rate float64
	Up   uint64
	Down uint64
}

//export GoServer
func GoServer(addr string, isGo bool) *C.char {
	return C.CString(server(addr, isGo))
}
func server(addr string, isGo bool) string {
	start()
	h := macaron.Classic()
	h.Use(macaron.Renderer())
	h.Get("/", func(ctx *macaron.Context) {
		var out []interface{}
		for _, k := range []string{"cpu", "memory", "loadavg", "net", "disk"} {
			if v, ok := m.Load(k); ok {
				out = append(out, v)
			}
		}
		log.Println(out)
		ctx.JSON(200, out)
	})
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Panic(err)
	}
	log.Println(l.Addr())
	if isGo {
		go http.Serve(l, h)
		return l.Addr().String()
	} else {
		err := http.Serve(l, h)
		return err.Error()
	}
}
func main() {
	log.Panic(server("127.0.0.1:", false))
}
