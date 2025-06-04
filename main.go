package main

import (
	"go-blog/core"
	"go-blog/global"
	"go-blog/initialize"
)

func main() {
	global.Config = core.InitConf()
	global.Log = core.InitLogger()
	initialize.OtherInit()
	global.DB = initialize.InitGorm()
	global.Redis = initialize.ConnectRedis()
	global.ESClient = initialize.ConnectES()

	defer global.Redis.Close()

	initialize.InitCron()

	core.RunServer()
}
