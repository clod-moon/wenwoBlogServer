package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"io/ioutil"
	"wenwoBlogServer/models"
	"errors"
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
	"math/rand"
)

type UserController struct {
	TokenController       `json:"-"`
	beego.Controller      `json:"-"`
	RetHeader RetHeader   `json:"header"`
	UserInfo  models.User `json:"user_info"`
	Type      int         `json:"type"`
	Token     string      `json:"token"`
}

func NewUserControlle()*UserController{
	return &UserController{}
}

func  GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}


func (u *UserController) Signup() {
	var retbody string
	var retErr error
	user := new(models.User)
	data, err := ioutil.ReadAll(u.Ctx.Request.Body)
	if err != nil {
		return
	} else {
		err := json.Unmarshal(data, &user)
		if err != nil {
			return
		}
		if user.LoginType == LOGIN_TYPE_GETAUTHCODE {
			retbody,retErr=u.getRetBody(CODE_SUCCESS,MSG_SUCCESS)
			if retErr != nil {
				u.Ctx.WriteString(retbody)
			}else{
				u.Ctx.WriteString(retErr.Error())
			}
		}else if user.LoginType == LOGIN_TYPE_AUTHCODE {

		}else if user.LoginType == LOGIN_TYPE_MOBILE {

		}
		err = user.Insert()
		if err != nil {
			u.Ctx.WriteString(err.Error())
		} else {
			u.Ctx.WriteString("success")
		}
	}
	u.Ctx.WriteString(string(data))
}

func(u *UserController)getRetBody(code int, msg string)(string,error){
	if code != CODE_SUCCESS{
		newobj:=NewUserControlle()
		newobj.RetHeader.Code = code
		newobj.RetHeader.Message = msg
		retbody,err :=json.Marshal(newobj)
		if err != nil{
			beego.Error(err.Error())
			return string(retbody),errors.New(err.Error())
		}
	}else{
		u.RetHeader.Code = code
		u.RetHeader.Message = msg
		retbody,err :=json.Marshal(u)
		if err != nil{
			beego.Error(err.Error())
			return string(retbody),errors.New(err.Error())
		}else {
			return string(retbody),nil
		}
	}
	return "",nil
}

func (u *UserController) Login() {

	var retbody string
	var retErr error
	user := new(models.User)
	data, err := ioutil.ReadAll(u.Ctx.Request.Body)
	if err != nil {
		return
	} else {
		err := json.Unmarshal(data, &user)
		if err != nil {
			return
		}
		if user.LoginType == LOGIN_TYPE_GETAUTHCODE{

		}else if user.LoginType == LOGIN_TYPE_AUTHCODE {

		}else if user.LoginType == LOGIN_TYPE_MOBILE{
			u.UserInfo.Mobile = user.Mobile
			err = u.UserInfo.QueryByMobile()
			if err != nil {
				if err == orm.ErrNoRows{
					retbody,retErr = u.getRetBody(CODE_USER_USERNAME_NOT_EXIST,
						MSG_USER_USERNAME_NOT_EXIST)
				}else{
					retbody,retErr = u.getRetBody(CODE_USER_USERNAME_NOT_EXIST,
						MSG_USER_USERNAME_NOT_EXIST)
				}
				if retErr != nil {
					u.Ctx.WriteString(retbody)
				}else{
					u.Ctx.WriteString(retErr.Error())
				}

			} else {
				if u.UserInfo.Password == user.Password {
					tokenString,err:= u.CreateToken(&u.UserInfo)
					if err == nil{
						u.Token =tokenString
					}
					retbody,retErr =u.getRetBody(CODE_SUCCESS,
						MSG_SUCCESS)
					if retErr != nil {
						u.Ctx.WriteString(retErr.Error())
					}else{
						fmt.Println(retbody)
						u.Ctx.WriteString(retbody)
					}
				}else{
					retbody,retErr =u.getRetBody(CODE_PASSWORD_ERROR,
						MSG_PASSWORD_ERROR)
					if retErr != nil {
						u.Ctx.WriteString(retbody)
					}else{
						u.Ctx.WriteString(retErr.Error())
					}
				}

			}
		}
	}
}

func (u *UserController) GetAuthCode() (int, string) {
	var authCode string
	return 0, authCode
}


func (u *UserController)GetArticle(){
	//var retbody string
	//var retErr error

	data, err := ioutil.ReadAll(u.Ctx.Request.Body)
	if err != nil {
		return
	} else {
		err := json.Unmarshal(data, &u.UserInfo)
		if err != nil {
			return
		}
	}
}


// ParseToken parse JWT token in http header.

