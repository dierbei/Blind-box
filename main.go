package main

import "github.com/dierbei/blind-box/initialize"

func main() {
	initialize.InitLogger()
	initialize.InitConfig()
	engine := initialize.InitRouter()
	engine.Run(":8080")
}
