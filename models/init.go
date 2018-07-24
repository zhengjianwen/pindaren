package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	"github.com/astaxie/beego"
	"github.com/zhengjianwen/utils/log"
	"fmt"
)

var Orm *xorm.Engine

func MysqlInit() {
	log.Println("开始连接数据。。。")
	Orm,_ = ConnOrm(getMysqlUrl())
	if Orm == nil{
		panic("数据库连接失败")
	}
}

func ConnOrm(addr string) (*xorm.Engine,error){
	db, err := xorm.NewEngine("mysql", addr)
	if err != nil {
		log.Errorf("[InitMysql] connect mysql => %s\n %v", addr,err)
		return nil,err
	}
	if db == nil{
		return nil,fmt.Errorf("连接数据库失败")
	}
	show,_ := beego.AppConfig.Bool("showsql")
	db.ShowSQL(show)
	db.Logger().SetLevel(core.LOG_INFO)
	return db,nil
}

func getMysqlUrl() string {
	user := beego.AppConfig.String("mysqluser")
	pasw := beego.AppConfig.String("mysqlpass")
	addr := beego.AppConfig.String("mysqlurls")
	port := beego.AppConfig.String("mysqlport")
	dbname := beego.AppConfig.String("mysqldb")
	url := "%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=Asia%%2FShanghai"
	return fmt.Sprintf(url,user,pasw,addr,port,dbname)
}
