package main

import (
	"go-blog/core"
	"go-blog/global"
)

func main() {
	global.Config = core.InitConf()
	global.Log = core.InitLogger()

	core.RunServer()
}
