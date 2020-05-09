package lib

import (
	"flag"
	"log"
	"net"
	dlog "tbwisk/common/log"
	"time"

	"github.com/TBWISK/goconf"
)

var TimeLocation *time.Location
var TimeFormat = "2006-01-02 15:04:05"
var DateFormat = "2006-01-02"
var LocalIP = net.ParseIP("127.0.0.1")

//公共初始化函数：支持两种方式设置配置文件
//
//函数传入配置文件 Init("./conf/dev/")
//如果配置文件为空，会从命令行中读取 	  -config conf/dev/
func Init(configPath string) error {
	return InitModule(configPath, []string{"base", "mysql", "redis"})
}

var cparse *goconf.ConfigParse

//模块初始化
func InitModule(configPath string, modules []string) error {
	var conf *string
	if len(configPath) > 0 {
		conf = &configPath
	} else {
		conf = flag.String("config", "", "input config file like ./conf/dev/")
		flag.Parse()
	}
	// if *conf == "" {
	// 	flag.Usage()
	// 	os.Exit(1)
	// }
	cparse = goconf.NewConfigParse(configPath)
	log.Println("------------------------------------------------------------------------")
	log.Printf("[INFO]  config=%s\n", *conf)
	log.Printf("[INFO] %s\n", " start loading resources.")

	// 设置ip信息，优先设置便于日志打印
	ips := GetLocalIPs()
	if len(ips) > 0 {
		LocalIP = ips[0]
	}

	// // 设置时区
	// if location, err := time.LoadLocation(cparse.GetConfig().Section("base").Key("time_location").String()); err != nil {
	// 	return err
	// } else {
	// 	TimeLocation = location
	// }
	log.Printf("[INFO] %s\n", " success loading resources.")
	log.Println("------------------------------------------------------------------------")
	return nil
}

//公共销毁函数
func Destroy() {
	log.Println("------------------------------------------------------------------------")
	log.Printf("[INFO] %s\n", " start destroy resources.")
	// CloseDB()
	dlog.Close()
	log.Printf("[INFO] %s\n", " success destroy resources.")
}

func GetLocalIPs() (ips []net.IP) {
	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		return nil
	}
	for _, address := range interfaceAddr {
		ipNet, isValidIpNet := address.(*net.IPNet)
		if isValidIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP)
			}
		}
	}
	return ips
}

func InArrayString(s string, arr []string) bool {
	for _, i := range arr {
		if i == s {
			return true
		}
	}
	return false
}
