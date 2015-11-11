// main.go
package main

import (
	"github.com/astaxie/beego"
	"github.com/toophy/login/app"
	_ "github.com/toophy/login/routers"
)

// Gogame framework version.
const (
	VERSION = "0.0.2"
)

func main() {

	if app.GetApp().Start(100) {
		//
		go beego.Run(beego.AppConfig.String("Listen"))
		// 主协程
		go app.Main_go()
		// 等待结束
		app.GetApp().WaitExit()
	}
}
