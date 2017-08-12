# Linux系统监视悬浮窗   
用QML+Golang写的系统资源监视悬浮框，前后端分离，可监控主机  
纯QML实现在[这里](https://github.com/myml/qml-monitor)  

# 命令行参数
```
//启动前后端，监控本机  
./go-monitor-linux 
 
//启动后端监控，无界面,不指定端口则随机
./go-monitor-linux 127.0.0.1:

//启动前端界面,监控远程主机（http开头）
./go-monitor-linux http://127.0.0.1:43587
```

# 编译教程
## 环境
* Qt5.9.1
* Golang1.9bate2  

## Linux编译过程
```
git clone https://github.com/myml/go-monitor.git
cd go-monitor
go build -buildmode=c-archive -o libMonitor.a main.go
cd go-monitor-linux
qmake
make
./go-monitor-linux
```
# 说明
* 界面参考 [软媒魔方](http://mofang.ruanmei.com/)
* 实现参考 [深度系统监视器原理剖析](http://www.jianshu.com/p/deb0ed35c1c2?from=jiantop.com)
* 实现细节 [Wiki](https://github.com/myml/qml-monitor/wiki)
* 开发环境 Qt5.9.1
* 发布环境 Qt5.9.1
* 测试环境 Deepin15.4.1
* 开源协议 [WTFPL](https://github.com/myml/go-monitor/blob/master/LICENSE)
* 截图  
 ![截图](s.png)