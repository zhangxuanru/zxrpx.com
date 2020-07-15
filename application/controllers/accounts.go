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

//用户登录表单
type UserLogin struct {
	UserName string `form:"username" binding:"required"`
	PassWord string `form:"password" binding:"required"`
	Next     string `form:"next" json:"next"`
	Form
}

type Form struct {
	Token string `form:"token" binding:"required" json:"token"`
}

//用户注册表单
type UserRegister struct {
	UserLogin
	Email string `form:"email" binding:"required"`
}

var formToken string

//登录 - 显示登录模板
func Login(c *gin.Context) {
	account, _ := getUser(c)
	formToken = logic.GenLoginFormToken()
	nextUrl := c.DefaultQuery("next", "/")
	c.HTML(http.StatusOK, "login.html", gin.H{
		"frontDomain": configs.STATIC_DOMAIN,
		"cdnDomain":   configs.STATIC_CDN_DOMAIN,
		"token":       formToken,
		"nextUrl":     nextUrl,
		"account":     account,
	})
}

//执行登录
func LoginDo(c *gin.Context) {
	var (
		login          UserLogin
		form           Form
		accountService *services.Account
		account        *services.Account
		err            error
	)
	if err = c.ShouldBind(&login); err != nil || login.Token != formToken {
		logrus.Error("LoginDo err:", err)
		formToken = logic.GenLoginFormToken()
		form.Token = formToken
		c.JSON(http.StatusOK, ResponseErr(form, logic.LoginError))
		return
	}
	accountService = &services.Account{
		UserName: strings.TrimSpace(login.UserName),
		PassWord: logic.EncryptPassword(login.PassWord),
	}
	if account, err = accountService.Login(); err != nil {
		formToken = logic.GenLoginFormToken()
		form.Token = formToken
		c.JSON(http.StatusOK, ResponseErr(form, logic.LoginError))
		return
	}
	setUserCookie(c, account.Token)
	formToken = ""
	c.JSON(http.StatusOK, ResponseSucc(&UserLogin{
		Next: login.Next,
	}, logic.LoginSuccess))

}

//注销
func Logout(c *gin.Context) {
	delUserCookie(c)
	account, _ := getUser(c)
	c.HTML(http.StatusOK, "logout.html", gin.H{
		"frontDomain": configs.STATIC_DOMAIN,
		"account":     account,
		"website":     configs.WEBSITE,
	})
}

//注册
func Register(c *gin.Context) {
	account, _ := getUser(c)
	formToken = logic.GenLoginFormToken()
	c.HTML(http.StatusOK, "register.html", gin.H{
		"frontDomain": configs.STATIC_DOMAIN,
		"cdnDomain":   configs.STATIC_CDN_DOMAIN,
		"token":       formToken,
		"account":     account,
	})
}

//执行注册
func RegisterDo(c *gin.Context) {
	var (
		register       UserRegister
		form           Form
		accountService *services.Account
		account        *services.Account
		err            error
	)
	if err = c.ShouldBind(&register); err != nil || register.Token != formToken {
		logrus.Error("RegisterDo err:", err)
		formToken = logic.GenLoginFormToken()
		form.Token = formToken
		c.JSON(http.StatusOK, ResponseErr(form, logic.RegisterError))
		return
	}
	accountService = &services.Account{
		UserName: strings.TrimSpace(register.UserName),
		PassWord: logic.EncryptPassword(register.PassWord),
		Email:    strings.TrimSpace(register.Email),
	}
	if account, err = accountService.Register(); err != nil {
		logrus.Error("Register error:", err)
		formToken = logic.GenLoginFormToken()
		form.Token = formToken
		c.JSON(http.StatusOK, ResponseErr(form, err.Error()))
		return
	}
	setUserCookie(c, account.Token)
	formToken = ""
	c.JSON(http.StatusOK, ResponseSucc(&UserLogin{
		Next: register.Next,
	}, logic.RegisterSuccess))
}

//我的图片
func Media(c *gin.Context) {

}

//设置个人信息
func Settings(c *gin.Context) {

}

//消息
func Messages(c *gin.Context) {

}

//收藏
func Collect(c *gin.Context) {

}

//收藏列表
func Favorites(c *gin.Context) {

}

//上传
func Upload(c *gin.Context) {

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

//评论
func Comment(c *gin.Context) {

}

//喜欢
func Like(c *gin.Context) {

}
