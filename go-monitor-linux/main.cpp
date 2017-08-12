#include <QGuiApplication>
#include <QQmlApplicationEngine>
#include <QQmlContext>
#include <QDebug>
#include "../libMonitor.h"

int main(int argc, char *argv[])
{
    QGuiApplication app(argc, argv);
    QString addr;
    if(argc>1){
        addr=argv[1];
        if(!addr.contains("http")){
            GoString gAddr={argv[1],addr.length()};
            GoServer(gAddr,false);
        }
    }else{
        GoString gAddr={"127.0.0.1:",10};
        addr.sprintf("http://%s",GoServer(gAddr,true));
    }
    qDebug(addr.toLatin1());
    QQmlApplicationEngine engine;
    engine.load(QUrl(QStringLiteral("qrc:/main.qml")));
    engine.rootObjects().first()->setProperty("addr",addr);
    if (engine.rootObjects().isEmpty())
        return -1;
    return app.exec();
}
