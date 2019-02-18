package main

import (
	"net/http"
	"tianwei.pro/sam-agent"
	_ "tianwei.pro/sam-web/routers"
	"tianwei.pro/sam-web/rpc"
	_ "tianwei.pro/sam-web/rpc"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.ErrorHandler("404", pageNotFound)

	defer rpc.CloseClient()

	sam_agent.StartAgent()
	defer sam_agent.StopAgent()

	beego.SetStaticPath("/static", "static")

	beego.InsertFilter("/*", beego.BeforeRouter, sam_agent.SamFilter)
	beego.Run()
}


func pageNotFound(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusNotFound)
	rw.Write([]byte(`{"msg": "api not found"}`))
}