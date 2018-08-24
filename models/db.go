package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/zhengjianwen/utils/log"
	"github.com/go-xorm/xorm"
	"fmt"
	"github.com/astaxie/beego"
)


var Orm  *xorm.Engine

func InitMySQL() (err error) {
	Orm, err = ConnOrm()
	if err != nil {
		log.Println("[mysql][rongyun]",err)
		return err
	}
	err = Orm.Sync2(UserInfo{},Article{},Images{},Classify{})
	init_class()
	return err
}

func ConnOrm() (*xorm.Engine,error){
	mysqladdr := beego.AppConfig.String("mysql")
	if mysqladdr == ""{
		return nil,fmt.Errorf("数据库地址配置不正确")
	}
	name := beego.AppConfig.String("sqlname")
	if name == ""{
		name = "pingdaren"
	}
	addr := fmt.Sprintf("%s/%s?charset=utf8&loc=Asia%%2FShanghai",mysqladdr,name)
	log.Println("数据库地址：",addr)

	db, err := xorm.NewEngine("mysql", addr)
	if err != nil {
		log.Println("数据库连接失败",err)
	}
	if err = db.Ping(); err != nil {
		return nil,fmt.Errorf("链接数据库失败")
	}
	db.ShowSQL(beego.AppConfig.DefaultBool("sqlshow",true))
	db.SetMaxIdleConns(beego.AppConfig.DefaultInt("sqlidle",5))
	db.SetMaxOpenConns(beego.AppConfig.DefaultInt("sqlmax",1000))
	return db,nil
}
