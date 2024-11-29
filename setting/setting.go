package setting

import (
	"github.com/go-ini/ini"
	"github.com/spf13/viper"
	"log"
	"net"
	"sync"
)

type commonConf struct {
	HttpPort  string
	RPCPort   string
	Cluster   bool
	CryptoKey string
}

var CommonSetting = &commonConf{}

type etcdConf struct {
	Endpoints []string
}

var EtcdSetting = &etcdConf{}

type global struct {
	LocalHost      string //本机内网IP
	ServerList     map[string]string
	ServerListLock sync.RWMutex
}

var GlobalSetting = &global{}

var cfg *ini.File

func Setup() {
	var err error
	viper.SetConfigName("app")
	viper.AddConfigPath("github.com/goravel-socket/config")
	viper.SetConfigType("ini")
	if err = viper.ReadInConfig(); err != nil {
		log.Fatalf("setting.Setup, fail to read 'conf/app.ini': %v", err)
	}
	// 使用ini.Load解析读取到的配置文件内容
	//通过viper获取数据,并解析成ini结构体
	cfg, err = ini.Load(viper.ConfigFileUsed())
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}
	mapTo("common", cfg)
	mapTo("etcd", EtcdSetting)
	GlobalSetting = &global{
		LocalHost:  getIntranetIp(),
		ServerList: make(map[string]string),
	}
}

func Default() {
	CommonSetting = &commonConf{
		HttpPort:  "6000",
		RPCPort:   "7000",
		Cluster:   false,
		CryptoKey: "Adba723b7fe06819",
	}

	GlobalSetting = &global{
		LocalHost:  getIntranetIp(),
		ServerList: make(map[string]string),
	}
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}

// 获取本机内网IP
func getIntranetIp() string {
	addrs, _ := net.InterfaceAddrs()

	for _, addr := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}
