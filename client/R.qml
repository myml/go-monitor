import QtQuick 2.0

Rectangle {
	width: w
	height: w
	opacity: 0.7
	color:"#fff"
    function toGB(bytes) {
        if (bytes < 1024) {
            return bytes + "B"
        }
        if (bytes / 1024 < 1024) {
            return parseInt((bytes / 1024)) + "K"
        }
        return (bytes / 1024 / 1024).toFixed(1) + "M"
    }
	Rectangle {
        width: parent.width
        height: parent.height * Rate/100
		Behavior on height {
			NumberAnimation {
				duration: 300
			}
		}
        anchors.bottom: parent.bottom
		color: Rate<90?"#eee":"#f11"
        opacity: 0.9
    }
    Text {
        anchors.centerIn: parent
        anchors.verticalCenterOffset: -10
		text:Rate==-1?toGB(Up):Name
    }
    Text {
        anchors.centerIn: parent
        anchors.verticalCenterOffset: 10
		text:Rate==-1?toGB(Down):Rate.toFixed(1) + "%"
    }
}
