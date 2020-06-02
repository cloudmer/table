package service

import (
	"flag"
	"log"
	"os"
)

var (
	brokers  = ""
	version  = ""
	group    = ""
	topics   = ""
	assignor = ""
	oldest   = true
	verbose  = false
)

var YamlFile string

func FlagChecking()  {
	flag.StringVar(&YamlFile, "f", "", "yaml 配置文件绝对路径")
	flag.Parse()

	if len(YamlFile) == 0 {
		log.Println("yaml 配置文件绝对路径 -f, 或用 -h or -help 来获取帮助")
		os.Exit(0)
	}

	// 检查文件是否存在
	_, err := os.Stat(YamlFile)
	if os.IsNotExist(err) {
		log.Printf("yaml文件不存在 for path: %s", YamlFile)
		os.Exit(0)
	}
}