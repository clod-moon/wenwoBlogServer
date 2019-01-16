package models

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)


//初始化

//下面是统一的表名管理
func TableName(name string) string {
	prefix := beego.AppConfig.String("")
	return prefix + name
}

