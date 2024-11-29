package setting

import (
	"flag"
	"github.com/go-ini/ini"
	"log"
	"net"
	"path/filepath"
	"runtime"
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
	// 获取当前文件的绝对路径
	_, currentFilePath, _, _ := runtime.Caller(0)
	configDir := filepath.Join(filepath.Dir(currentFilePath), "../config")

	// 设置默认配置文件路径
	defaultConfigFile := filepath.Join(configDir, "app.ini")

	// 命令行参数
	configFile := flag.String("c", defaultConfigFile, "Path to the configuration file")
	flag.Parse()
	var err error
	cfg, err = ini.Load(*configFile)
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
