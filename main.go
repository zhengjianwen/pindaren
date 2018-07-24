package main

import (
	_ "github.com/zhengjianwen/pindaren/routers"
	"github.com/astaxie/beego"
	"github.com/zhengjianwen/pindaren/models"
)

func main() {
	models.MysqlInit()
	beego.Run()
}

