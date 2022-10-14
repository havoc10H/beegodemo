package main

import (
	_ "quickstart/routers"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	// beego.Run()
	web.Run()
}
