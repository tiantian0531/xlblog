package main

import (
	"fmt"
	"xlblog/core"
	"xlblog/global"
)

func main() {
	global.Viper = core.Viper("./settings.yaml")

	fmt.Println(global.Config.Redis.Password)

	fmt.Println(global.Config.Redis.Password)
}
