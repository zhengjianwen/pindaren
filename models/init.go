package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
	"github.com/zhengjianwen/utils/log"
)

var Orm *xorm.Engine

func init()  {
	log.Debugf("[mysql][init] start")
	maxIdle := 30
	maxConn := 30
	addr := getMysqlUrl()
	log.Debugf("[mysql][init] %s",addr)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", addr,maxIdle,maxConn)
	orm.DefaultTimeLoc = time.UTC
	orm.RegisterModel(new(AuthUser),new(UserInfo),new(Classify),new(Member),new(Comment),new(Collection),
		new(News),new(Praise),new(tags),new(Images),new(Article))

	name := "default"

	err := orm.RunSyncdb(name, false, true)
	if err != nil{
		log.Errorf("[mysql][init] %v",err)
	}
}

func getMysqlUrl() string {
	user := beego.AppConfig.DefaultString("mysqluser","root")
	pasw := beego.AppConfig.DefaultString("mysqlpass","123456")
	addr := beego.AppConfig.DefaultString("mysqlurls","127.0.0.1")
	port := beego.AppConfig.DefaultInt64("mysqlport",3306)
	dbname := beego.AppConfig.DefaultString("mysqldb","pindaren")
	url := "%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=Asia%%2FShanghai"
	return fmt.Sprintf(url,user,pasw,addr,port,dbname)
}
