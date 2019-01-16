package models

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
)

func init() {
	day := time.Now().Day()
	config := fmt.Sprintf(`{"filename":"logs/%s/user.log"}`,day)
	beego.SetLogger("file", config)
}

type User struct {
	Id            int       `json:"-" pk:"auto" orm:"column(account_id)"`
	NickName      string    `json:"nick_name" orm:"column(account_name);size(64)"`
	Content       string    `json:"content" orm:"column(account_content);type(text)"`
	Head          string    `json:"Head" orm:"column(head);type(text)"`
	Password      string    `json:"passwd" orm:"null;size(64)"`
	Email         string    `json:"mail" orm:"null;size(64)"`
	Ip            string    `json:"ip" orm:"null;size(15)"`
	Qq            string    `json:"qq" orm:"null;size(64)"`
	Weichat       string    `json:"weichat" orm:"null;size(64)"`
	Weibo         string    `json:"weibo" orm:"null;size(64)"`
	Mobile        string    `json:"mobile" orm:"size(20)"`
	LastLoginTime time.Time `json:"time_update" orm:"null;type(datetime)"`
	CreateTime    time.Time `json:"time_create" orm:"default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateTime    time.Time `orm:"auto_now_add;null;type(datetime)"`
	LoginType     int 		`json:"login_type"`
}

func (u *User) GetTableName() string{
	return "user"
}
//初始化
func NewAdmin() *User {
	return new(User)
}

//初始化列表
func (u *User) newMakeDataArr() []User {
	return make([]User, 0)
}

func (u *User) AddPersonnel() {

}

func (u *User) Insert() error {
	o := orm.NewOrm() //new一个Orm，默认使用名为default的数据库
	//o.Using("myApp") // 你可以使用Using函数指定其他数据库
	_, err := o.Insert(u)
	//fmt.Println("err:",err.Error())
	return err
}

func (u *User)Update()error {
	o := orm.NewOrm()
	_,err := o.Update(u)
	return err
}

func(u *User) QueryByMobile() error {
	o := orm.NewOrm() //new一个Orm，默认使用名为default的数据库
	//o.Using("wenwo") // 你可以使用Using函数指定其他数据库
	err:= o.QueryTable(u.GetTableName()).Filter("mobile",u.Mobile).One(u)
	//err := o.Raw("SELECT  account_name FROM user WHERE mobile = ?", u.Mobile).QueryRow(&u)
	if err != nil{
		beego.Warn()
		return err
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		beego.Warn("Not row found where mobile = %s",u.Mobile)
		return orm.ErrNoRows
	}
	return nil
}
