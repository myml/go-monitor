import QtQuick 2.5
import QtQuick.Controls 1.4
import QtQuick.Window 2.0

ApplicationWindow {
    property int w: 50
    property string addr: ""
    id: root
    visible: true
    flags: Qt.Tool | Qt.FramelessWindowHint | Qt.WindowStaysOnTopHint
    color: "transparent"

    height: w
    width: data.count * w
    x: Screen.width
    y: Screen.height
    ListModel {
        id: data
    }
    ListView {
        anchors.fill: parent
        orientation: ListView.Horizontal
        model: data
        delegate: R {
        }
    }
    function getInfo() {
        var xhr = new XMLHttpRequest()
        xhr.open("GET", addr, false)
        xhr.send()
        return JSON.parse(xhr.responseText)
    }
    Timer {
        repeat: true
        running: true
        onTriggered: {
            var info = getInfo()
            for (var i in info) {
                data.set(i, info[i])
            }
        }
    }
    Component.onCompleted: {


        //        var info = getInfo()
        //        for (var k in info) {
        //            data.append(info[k])
        //        }
    }

    MouseArea {
        anchors.fill: parent
        property int mx: 0
        property int my: 0
        onPressed: {
            mx = mouseX
            my = mouseY
        }
        onPositionChanged: {
            root.x += mouseX - mx
            root.y += mouseY - my
        }
    }
}
