package setting

import (
	"fmt"
	"os"
	"time"

	"github.com/antchfx/xmlquery"
)

var configuration *xmlquery.Node

var serverName = "login_service"
var host = "localhost"
var port = "50001"
var registerServer = "http://localhost:2379"
var hmacKeyPath = "./hmacKey"
var logPath = "./Logs"

func init() {

	f, err := os.Open("./setting.config")
	if err != nil {
		panic(err)
	}

	doc, err := xmlquery.Parse(f)
	if err != nil {
		panic(err)
	}
	configuration = xmlquery.FindOne(doc, "//configuration")
	if n := configuration.SelectElement("ServiceName"); n != nil {
		serverName = n.InnerText()
	}
	if n := configuration.SelectElement("Host"); n != nil {
		host = n.InnerText()
	}
	if n := configuration.SelectElement("Port"); n != nil {
		port = n.InnerText()
	}
	if n := configuration.SelectElement("RegisterServer"); n != nil {
		registerServer = n.InnerText()
	}
	if n := configuration.SelectElement("HmacKeyPath"); n != nil {
		hmacKeyPath = n.InnerText()
	}
	if n := configuration.SelectElement("LogPath"); n != nil {
		logPath = n.InnerText()
	}
}

func ServiceName() string {
	return serverName
}

func Host() string {
	return host
}

func Port() string {
	return port
}

func RegisterServer() string {
	return registerServer
}

func HmacKeyPath() string {
	return hmacKeyPath
}

func GetLogPath() string {
	_, err := os.Stat(logPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(logPath, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}
	}

	return fmt.Sprintf("%v/%v.log", logPath, time.Now().Format("20060102150405"))
}
