/*
@Time : 2020/7/13 12:05
@Author : zxr
@File : base
@Software: GoLand
*/
package controllers

import (
	"pix/application/services"
	"pix/configs"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	SuccessCode  = 200
	ErrCode      = -1
	LOGIN_COOKIE = "token" //用户登录的cookie名
)

//获取登录用户
func getUser(c *gin.Context) (account *services.AccountAuth, err error) {
	var (
		token        string
		customClaims *services.CustomClaims
	)
	account = &services.AccountAuth{}
	if token, err = c.Cookie(LOGIN_COOKIE); err != nil {
		return
	}
	if customClaims, err = services.NewJWT().ParseToken(token); err != nil {
		logrus.Error("ParseToken error1:", err)
		return
	}
	//快到期之前就延长生效期
	if customClaims.ExpiresAt <= time.Now().Unix()-500 {
		if token, err = services.NewJWT().RefreshToken(token); err != nil {
			return
		}
		setUserCookie(c, token)
		customClaims, err = services.NewJWT().ParseToken(token)
	}
	if err != nil {
		logrus.Error("ParseToken error2:", err)
		return
	}
	account = &services.AccountAuth{
		UserId:   customClaims.UserId,
		UserPxId: customClaims.UserPxId,
		UserName: customClaims.UserName,
	}
	return
}

//设置用户登录cookie
func setUserCookie(c *gin.Context, token string) {
	c.SetCookie(LOGIN_COOKIE, token, 0, "/", configs.COOKIEDOMAIN, false, true)
}

//删除用户登录cookie
func delUserCookie(c *gin.Context) {
	c.SetCookie(LOGIN_COOKIE, "", -1, "/", configs.COOKIEDOMAIN, false, true)
}

func Response(code int, data interface{}, msg ...string) *response {
	r := &response{
		Code: code,
		Data: data,
	}
	if len(msg) > 0 {
		r.Msg = msg[0]
	}
	return r
}

func ResponseSucc(data interface{}, msg string) *response {
	return Response(SuccessCode, data, msg)
}

func ResponseErr(data interface{}, msg string) *response {
	return Response(ErrCode, data, msg)
}
