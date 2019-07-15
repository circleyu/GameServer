package setting

import (
	"fmt"
	"os"
	"time"

	"github.com/antchfx/xmlquery"
	"github.com/apsdehal/go-logger"
)

var configuration *xmlquery.Node

var serverName = "login_service"
var host = "localhost"
var port = "50001"
var registerServer = "http://localhost:2379"
var hmacKeyPath = "./hmacKey"
var logPath = "./Logs"

func Init() {

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

func CreateLogger(moduleName string) (*logger.Logger, func(), error) {
	_, err := os.Stat(logPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(logPath, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}
	}

	logFilePath := fmt.Sprintf("%v/%v.log", logPath, time.Now().Format("20060102150405"))

	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	log, err := logger.New(moduleName, 0, file)
	if err != nil {
		panic(err) // Check for error
	}

	// Show warning with format message
	log.SetFormat("[%{time}] [%{level}] [%{module}] %{message} (in %{filename}:%{line})")

	return log, func() { file.Close() }, err
}
