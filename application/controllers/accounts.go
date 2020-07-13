/*
@Time : 2020/7/8 19:10
@Author : zxr
@File : accounts
@Software: GoLand
@desc :  用户帐户相关操作
*/
package controllers

import (
	"net/http"
	"pix/application/logic"
	"pix/application/services"
	"pix/configs"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserLogin struct {
	UserName string `form:"username" binding:"required"`
	PassWord string `form:"password" binding:"required"`
	Next     string `form:"next" json:"next"`
	FormToken
}
type FormToken struct {
	Token string `form:"token" binding:"required" json:"token"`
}

var loginToken string

//登录 - 显示登录模板
func Login(c *gin.Context) {
	loginToken = logic.GenLoginToken()
	nextUrl := c.DefaultQuery("next", "/")
	c.HTML(http.StatusOK, "login.html", gin.H{
		"frontDomain": configs.STATIC_DOMAIN,
		"cdnDomain":   configs.STATIC_CDN_DOMAIN,
		"token":       loginToken,
		"nextUrl":     nextUrl,
	})
}

//登录
func LoginDo(c *gin.Context) {
	var (
		login          UserLogin
		formToken      FormToken
		accountService *services.Account
		account        *services.Account
		err            error
	)
	if err = c.ShouldBind(&login); err != nil || login.Token != loginToken {
		logrus.Error("LoginDo err:", err)
		loginToken = logic.GenLoginToken()
		formToken.Token = loginToken
		c.JSON(http.StatusOK, ResponseErr(formToken, logic.LoginError))
		return
	}
	accountService = &services.Account{
		UserName: strings.TrimSpace(login.UserName),
		PassWord: logic.Md5(strings.TrimSpace(login.PassWord)),
	}
	if account, err = accountService.Login(); err != nil {
		loginToken = logic.GenLoginToken()
		formToken.Token = loginToken
		c.JSON(http.StatusOK, ResponseErr(formToken, logic.LoginError))
		return
	}
	c.SetCookie("token", account.Token, 3600, "/", configs.COOKIEDOMAIN, false, true)
	c.JSON(http.StatusOK, ResponseSucc(&UserLogin{
		Next: login.Next,
	}, logic.LoginSuccess))
}

//注销
func Logout(c *gin.Context) {

}

//注册
func Register(c *gin.Context) {

}

//评论
func Comment(c *gin.Context) {

}

//喜欢
func Like(c *gin.Context) {

}

//收藏
func Collect(c *gin.Context) {

}

//关注
func Follow(c *gin.Context) {
	parAuthorId := c.Param("authorId")
	authorId, err := strconv.Atoi(parAuthorId)
	if err != nil {
		c.Redirect(http.StatusFound, "/404")
	}
	services.NewAccount().Follow(0, authorId)
}

//关注列表
func Following(c *gin.Context) {

}