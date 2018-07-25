package main

import (
	_ "github.com/zhengjianwen/pindaren/routers"
	_"github.com/zhengjianwen/pindaren/models"
	_"github.com/zhengjianwen/pindaren/initial"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

)

func main() {
	beego.Run()
	orm.RunCommand()
}

