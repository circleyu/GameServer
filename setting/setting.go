package setting

import (
	"github.com/antchfx/xmlquery"
	"os"
)

var configuration *xmlquery.Node

var serverName = "login_service"
var host = "localhost"
var port = "50001"
var registerServer = "http://localhost:2379"
var hmacKeyPath = "./hmacKey"

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
		host= n.InnerText()
	}
	if n := configuration.SelectElement("Port"); n != nil {
		port= n.InnerText()
	}
	if n := configuration.SelectElement("RegisterServer"); n != nil {
		registerServer= n.InnerText()
	}
	if n := configuration.SelectElement("HmacKeyPath"); n != nil {
		hmacKeyPath = n.InnerText()
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
