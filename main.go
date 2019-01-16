package main

import (
	_ "wenwoBlogServer/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"wenwoBlogServer/models"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dbType := beego.AppConfig.String("db_type")
	//连接名称
	dbAlias := beego.AppConfig.String(dbType + "::db_alias")
	//数据库名称
	dbName := beego.AppConfig.String(dbType + "::db_name")
	//数据库连接用户名
	dbUser := beego.AppConfig.String(dbType + "::db_user")
	//数据库连接用户名
	dbPwd := beego.AppConfig.String(dbType + "::db_pwd")
	//数据库IP（域名）
	dbHost := beego.AppConfig.String(dbType + "::db_host")
	//数据库端口
	dbPort := beego.AppConfig.String(dbType + "::db_port")

	dbCharset := beego.AppConfig.String(dbType + "::db_charset")
	orm.RegisterDataBase(dbAlias, dbType, dbUser + ":" + dbPwd + "@tcp(" + dbHost + ":"+
		dbPort+ ")/"+ dbName+ "?charset="+ dbCharset, 30)

	orm.RegisterModel(new(models.User))

	orm.RunSyncdb("default", false, true)

	beego.SetLogger("file", `{"filename":"logs/test.log"}`)

	beego.Error("test")

}

func main() {
	beego.Run()
}
