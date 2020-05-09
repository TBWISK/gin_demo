package public

import (
	"log"
	"strings"

	"github.com/TBWISK/goconf"
)

func init() {
	ips := GetLocalIPs()
	if len(ips) > 0 {
		LocalIP = ips[0]
	}
}

//GetConfEnv 获取配置环境名
func GetConfEnv() string {
	return confEnv
}

//confEnv 配置环境
var confEnv string

//Destroy 资源销毁
func Destroy() {

}

var cparse *goconf.ConfigParse

//InitModule 配置初始化
func InitModule(configPath string) error {

	cparse = goconf.NewConfigParse(configPath)
	// 设置ip信息，优先设置便于日志打印

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

//公共初始化函数：支持两种方式设置配置文件
//
//函数传入配置文件 Init("./conf/dev/")
//如果配置文件为空，会从命令行中读取 	  -config conf/dev/
func Init(configPath string) error {
	return InitModule(configPath)
}

var DebugMode string

//GetStringConf 获取get配置信息
func GetStringConf(section, key string) string {
	// 	keys := strings.Split(key, ".")
	// 	if len(keys) < 2 {
	// 		return ""
	// 	}
	// 	v, ok := ViperConfMap[keys[0]]
	// 	if !ok {
	// 		return ""
	// 	}
	// 	confString := v.GetString(strings.Join(keys[1:len(keys)], "."))
	// 	return confString
	return cparse.GetConfig().Section(section).Key(key).MustString("")
}

//GetIntConf 获取信息
func GetIntConf(section, key string) int {
	// keys := strings.Split(key, ".")
	// if len(keys) < 2 {
	// 	return 0
	// }
	// v := ViperConfMap[keys[0]]
	// conf := v.GetInt(strings.Join(keys[1:len(keys)], "."))
	// return conf
	return cparse.GetConfig().Section(section).Key(key).MustInt()
}

//GetStringSliceConf 获取get配置信息
func GetStringSliceConf(secKey string, key string) []string {
	items := cparse.GetConfig().Section(secKey).Key(key).Strings(",")
	for idx := range items {
		items[idx] = strings.Replace(items[idx], "\"", "", -1)
	}
	return items
}
