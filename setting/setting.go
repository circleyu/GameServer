package setting

import (
	"fmt"
	"os"
	"strconv"
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
var log2Files = false
var dbServer = "localhost"
var dbPort = "27017"
var dbUser = "root"
var dbPassword = "123456"

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
	if n := configuration.SelectElement("Log2Files"); n != nil {
		b, err := strconv.ParseBool(n.InnerText())
		if err == nil {
			log2Files = b
		}
	}
	if n := configuration.SelectElement("DBServer"); n != nil {
		dbServer = n.InnerText()
	}
	if n := configuration.SelectElement("DBPort"); n != nil {
		dbPort = n.InnerText()
	}
	if n := configuration.SelectElement("DBUser"); n != nil {
		dbUser = n.InnerText()
	}
	if n := configuration.SelectElement("DBPassword"); n != nil {
		dbPassword = n.InnerText()
	}
}

// ServiceName 獲得server名稱
func ServiceName() string {
	return serverName
}

// Host 獲得主機位置
func Host() string {
	return host
}

// Port 獲得Port
func Port() string {
	return port
}

// RegisterServer 獲得etcd主機位置
func RegisterServer() string {
	return registerServer
}

// HmacKeyPath 獲得HmacKey檔案位置
func HmacKeyPath() string {
	return hmacKeyPath
}

// Log2Files 是否寫入log檔案
func Log2Files() bool {
	return log2Files
}

// DBServer 資料庫位置
func DBServer() string {
	return dbServer
}

// DBPort 資料庫Port
func DBPort() string {
	return dbPort
}

// DBUser 資料庫使用者名稱
func DBUser() string {
	return dbUser
}

// DBPassword 資料庫使用者密碼
func DBPassword() string {
	return dbPassword
}

// GetLogPath log檔案位置
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
