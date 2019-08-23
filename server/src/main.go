package main

import (
	"flag"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"src/db"
	"src/log"
	"src/routers"
	"src/utils/config"
)

var configFile = flag.String("config", "../config/config.yaml", "配置文件路径")

func init() {
	flag.Parse()

	config.InitConfig(*configFile) // 初始化配置
	log.InitLogrus()               // 初始化日志
	db.InitDB()                    // 初始化数据库
}

func main() {
	routers.InitRouter() //初始化路由
}
