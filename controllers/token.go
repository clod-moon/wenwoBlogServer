package controllers

import (
	"strings"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
	"errors"
	"wenwoBlogServer/models"
)


type TokenController struct {
	beego.Controller
}

// ParseToken parse JWT token in http header.
func (c *TokenController) ParseToken(authString string) *jwt.Token {

	kv := strings.Split(authString, " ")
	fmt.Println(len(kv),kv[0])
	if len(kv) != 2 || kv[0] != "Bearer" {
		beego.Error("AuthString invalid:", authString)
		return nil
	}
	tokenString := kv[1]

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		//// 可选项验证  'aud' claim
		//aud := "https://api.cn.atomintl.com"
		//checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
		//if !checkAud {
		//  return token, errors.New("Invalid audience.")
		//}
		// 必要的验证 'iss' claim
		iss := "https://atomintl.auth0.com/"
		checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
		if !checkIss {
			return token, errors.New("Invalid issuer.")
		}

		return []byte("mykey"), nil
	})
	if err != nil {
		beego.Error("Parse token:", err)
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				// That's not even a token
				return nil
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				return nil
			} else {
				// Couldn't handle this token
				return nil
			}
		} else {
			// Couldn't handle this token
			return nil
		}
	}
	if !token.Valid {
		beego.Error("Token invalid:", tokenString)
		return nil
	}
	beego.Debug("Token:", token)
	return token
}

func (c *TokenController) CreateToken(u *models.User) (string, error) {
	// 带权限创建令牌
	claims := make(jwt.MapClaims)
	claims["username"] = u.NickName
	claims["mobile"] = u.Mobile
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix() //30分钟有效期，过期需要重新登录获取token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用自定义字符串加密 and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte("mykey"))
	if err != nil {
		beego.Error("jwt.SignedString:", err)
		return "jwt.SignedString", err
	}
	return tokenString, nil
}


func (c *TokenController) GetUser(authString string) {
	var query = make(map[string]string)
	//get token
	token:= c.ParseToken(authString)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		beego.Debug("get ParseToken claims error")
		//return
	}
	beego.Debug("claims:", claims)
	var Email string = claims["Email"].(string)
	beego.Debug("Email:", Email)
	query["Uid"] = Email

}