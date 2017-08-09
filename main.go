package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"
)

func main() {

	//enable high dpi scaling
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	//use the material style for qml controls
	//"universal" is also available
	quickcontrols2.QQuickStyle_SetStyle("material")

	//create a qml application
	view := qml.NewQQmlApplicationEngine(nil)

	//load the main qml file
	view.Load(core.NewQUrl3("qml/main.qml", 0))

	//enter the main event loop
	gui.QGuiApplication_Exec()
}
