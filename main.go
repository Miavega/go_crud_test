package main

import (
	_ "github.com/Miavega/go_crud_test/routers"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
)

func main() {
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:admin123@127.0.0.1/postgres?sslmode=disable" )  //beego.AppConfig.String("sqlconn")
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

