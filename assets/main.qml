import QtQuick 2.0
import QtQuick.Controls 1.2
import QtQuick.Layouts 1.0

ApplicationWindow {
  id: rootWindow
  visible: true
  property int margin: 5
  width: 450
  height: 215
  title: qsTr("Youtube Comment Giveaway")
  minimumWidth: width
  minimumHeight: height
  //minimumWidth: mainLayout.Layout.minimumWidth + 2 * margin
  //minimumHeight: mainLayout.Layout.minimumHeight + 2 * margin
  maximumWidth: width
  maximumHeight: height
  //flags: Qt.FramelessWindowHint | Qt.Window
  flags: Qt.Window

    ColumnLayout {
        id: mainLayout
        height: 200
        anchors.fill: parent
        anchors.margins: margin
        GroupBox {
            id: rowBox
            visible: true
            flat: false
            Layout.fillHeight: false
            Layout.alignment: Qt.AlignLeft | Qt.AlignTop
            title: "Youtube Video ID"
            Layout.fillWidth: true

            RowLayout {
                id: rowLayout
                spacing: 5
                anchors.top: parent.top
                anchors.topMargin: 0
                anchors.bottom: parent.bottom
                anchors.bottomMargin: 0
                anchors.left: parent.left
                anchors.leftMargin: 0
                anchors.right: parent.right
                anchors.rightMargin: 0
                TextField {
                    id: youtubeUrlID
                    text: "7zv4ndak8Mo"
                    inputMask: ""
                    placeholderText: "7zv4ndak8Mo"
                    Layout.fillWidth: true
                }
                Button {
                    id: startButton
                    text: "Start"
                    onClicked: {
                        youtuberUI.start(youtubeUrlID.text, regExpPatternID.text)
                        labelAuthor.text = youtuberUI.textAuthor
                        labelComment.text = youtuberUI.textComment
                        imageAuthor.source = youtuberUI.imageLinkAuthor
                    }
                }
            }
        }


        GroupBox {
            id: rowBox3
            Layout.fillWidth: true
            RowLayout {
                id: rowLayout1
                anchors.leftMargin: 0
                anchors.top: parent.top
                anchors.topMargin: 0
                anchors.bottomMargin: 0
                TextField {
                    id: regExpPatternID
                    text: "(?i)ЭргозДайХаляву|Эргоз, дай халяву|Эргоз,дайхаляву|Эргоз, дайзаляву"
                    Layout.fillWidth: true
                    inputMask: ""
                    placeholderText: ""
                }
                spacing: 5
                anchors.left: parent.left
                anchors.bottom: parent.bottom
                anchors.right: parent.right
                anchors.rightMargin: 0
            }
            Layout.alignment: Qt.AlignLeft | Qt.AlignTop
            visible: true
            Layout.fillHeight: false
            title: "Youtube Comment RegExp"
            flat: false
        }



        GroupBox {
            id: rowBox2
            antialiasing: true
            Layout.fillWidth: true
            RowLayout {
                id: rowLayout2
                anchors.top: parent.top
                anchors.topMargin: 0
                anchors.bottom: parent.bottom
                anchors.bottomMargin: 0
                anchors.left: parent.left
                anchors.leftMargin: 0
                anchors.right: parent.right
                anchors.rightMargin: 0
                spacing: 5

                Image {
                    id: imageAuthor
                    width: 48
                    height: 48
                    sourceSize.height: 48
                    sourceSize.width: 48
                    source: "qrc:///assets/photo.jpg"
                }

                ColumnLayout {
                    id: columnLayout1
                    width: 100
                    height: 100
                    clip: false
                    Layout.fillHeight: true
                    Layout.fillWidth: true
                    Layout.alignment: Qt.AlignLeft | Qt.AlignTop

                    Label {
                        id: labelAuthor
                        color: "#0458ff"
                        text: "<html><a href=\"#\">Author</a></html>"
                        Layout.fillHeight: false
                        Layout.maximumWidth: 365
                        font.bold: true
                        verticalAlignment: Text.AlignTop
                        horizontalAlignment: Text.AlignLeft
                        Layout.fillWidth: false
                        Layout.alignment: Qt.AlignLeft | Qt.AlignTop
                        //onLinkActivated: Qt.openUrlExternally(text)
                        MouseArea {
                            id: mousearea
                            anchors.fill: parent
                            cursorShape: Qt.PointingHandCursor
                            onClicked: {
                                console.log("Going to winner Youtube page :)")
                                Qt.openUrlExternally(youtuberUI.linkAuthor)
                            }
                        }
                    }

                    Label {
                        id: labelComment
                        text: qsTr("Comment")
                        Layout.maximumWidth: 365
                        Layout.fillHeight: true
                        verticalAlignment: Text.AlignTop
                        horizontalAlignment: Text.AlignLeft
                        Layout.alignment: Qt.AlignLeft | Qt.AlignTop
                        wrapMode: Text.WordWrap
                        Layout.fillWidth: true
                    }
                }
            }
            Layout.alignment: Qt.AlignLeft | Qt.AlignTop
            visible: true
            Layout.fillHeight: false
            title: "Youtube Winner"
            flat: false
        }



    }

    Label {
        id: labelAuthor1
        x: 360
        y: 200
        color: "#0458ff"
        text: "<html><a href=\"http://ergoz.ru\">by ErgoZ Vaper</a></html>"
        Layout.fillWidth: false
        MouseArea {
            id: mousearea1
            anchors.fill: parent
            cursorShape: Qt.PointingHandCursor
            onClicked: {
                Qt.openUrlExternally("http://ergoz.ru")
            }
        }
        Layout.alignment: Qt.AlignLeft | Qt.AlignTop
        Layout.maximumWidth: 365
        Layout.fillHeight: false
        font.bold: true
        verticalAlignment: Text.AlignTop
        horizontalAlignment: Text.AlignLeft
    }
}
