/*
@Time : 2020/7/8 19:10
@Author : zxr
@File : accounts
@Software: GoLand
@desc :  用户帐户相关操作
*/
package controllers

import (
	"fmt"
	"net/http"
	"net/url"
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
	nextUrl, _ = url.QueryUnescape(nextUrl)
	if account.UserId > 0 {
		c.Redirect(http.StatusFound, nextUrl)
	}
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
	account := &services.AccountAuth{}
	c.HTML(http.StatusOK, "logout.html", gin.H{
		"frontDomain": configs.STATIC_DOMAIN,
		"account":     account,
		"website":     configs.WEBSITE,
	})
}

//注册
func Register(c *gin.Context) {
	delUserCookie(c)
	account := &services.AccountAuth{}
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

//设置个人信息 HTML
func Settings(c *gin.Context) {
	var (
		account     *services.AccountAuth
		userService *services.UserService
		err         error
	)
	if account, err = getUser(c); err != nil {
		query := url.QueryEscape("/accounts/settings/")
		c.Redirect(http.StatusFound, "/accounts/login/?next="+query)
	}
	userService = services.NewUserService()
	extend := userService.GetUserExtend(account.UserId)
	userInfo := userService.GetUserInfoByUid(account.UserId)
	formToken = logic.GenLoginFormToken()
	c.HTML(http.StatusOK, "settings.html", gin.H{
		"frontDomain": configs.STATIC_DOMAIN,
		"cdnDomain":   configs.STATIC_CDN_DOMAIN,
		"account":     account,
		"token":       formToken,
		"extend":      extend,
		"userInfo":    userInfo,
	})
}

//设置个人信息
func SettingsDo(c *gin.Context) {
	var (
		setting services.SettingUser
		account *services.AccountAuth
		form    Form
		err     error
	)
	if account, err = getUser(c); err != nil {
		c.JSON(http.StatusOK, Response(NotLoginCode, form, logic.BeforLogin))
		return
	}
	if err = c.ShouldBind(&setting); err != nil || setting.Token != formToken {
		logrus.Error("SettingsDo err:", err)
		formToken = logic.GenLoginFormToken()
		form.Token = formToken
		c.JSON(http.StatusOK, ResponseErr(form, logic.SettingErr))
		return
	}
	if err = services.NewAccount().SettingUser(&setting, account.UserId); err != nil {
		logrus.Error("SettingUser err:", err)
		formToken = logic.GenLoginFormToken()
		form.Token = formToken
		c.JSON(http.StatusOK, ResponseErr(form, err.Error()))
		return
	}
	formToken = logic.GenLoginFormToken()
	form.Token = formToken
	c.JSON(http.StatusOK, ResponseSucc(form, logic.SettingSuccess))
}

//修改密码
func ChangePassword(c *gin.Context) {
	var (
		account  *services.AccountAuth
		settPass services.UserChangePassword
		form     Form
		err      error
	)
	if account, err = getUser(c); err != nil {
		c.JSON(http.StatusOK, Response(NotLoginCode, form, logic.BeforLogin))
		return
	}
	if err = c.ShouldBind(&settPass); err != nil {
		logrus.Error("ChangePassword err:", err)
		c.JSON(http.StatusOK, ResponseErr(form, logic.ChangePassErr))
		return
	}
	settPass.NewPassword = logic.EncryptPassword(settPass.NewPassword)
	settPass.OldPassword = logic.EncryptPassword(settPass.OldPassword)
	settPass.ConfPassword = logic.EncryptPassword(settPass.ConfPassword)

	if err = services.NewAccount().ChangePassword(account.UserId, settPass); err != nil {
		c.JSON(http.StatusOK, ResponseErr(form, err.Error()))
		return
	}
	delUserCookie(c)
	c.JSON(http.StatusOK, ResponseSucc(form, logic.SettingSuccess))
}

//关注列表
func Following(c *gin.Context) {
	var (
		account *services.AccountAuth
		err     error
	)
	if account, err = getUser(c); err != nil {
		query := url.QueryEscape("/accounts/following/")
		c.Redirect(http.StatusFound, "/accounts/login/?next="+query)
	}
	list := services.NewAccount().Following(account.UserPxId)

	c.HTML(http.StatusOK, "following.html", gin.H{
		"frontDomain": configs.STATIC_DOMAIN,
		"account":     account,
		"list":        list.List,
	})
}

//关注
func Follow(c *gin.Context) {
	var (
		account  *services.AccountAuth
		authorId int
		err      error
	)
	parAuthorId := c.Param("authorId")
	if authorId, err = strconv.Atoi(parAuthorId); err != nil {
		c.String(http.StatusOK, "<script>alert('非法请求')</script><i class='check'>✓</i> 关注</span>")
		return
	}
	if account, err = getUser(c); err != nil {
		c.String(http.StatusOK, "<script>alert('请先登录')</script> <i class='check'>✓</i> 关注</span>")
		return
	}
	if status, err := services.NewAccount().Follow(account.UserId, authorId); err != nil || status == false {
		c.String(http.StatusOK, "<script>alert('"+err.Error()+"')</script> <i class='check'>✓</i> 关注</span>")
		return
	}
	c.String(http.StatusOK, "script:$('#follow_button').addClass('pure-button-disabled');")
}

//收藏
func Collect(c *gin.Context) {
	var (
		account *services.AccountAuth
		imgId   int
		status  bool
		err     error
	)
	imgIdStr := c.Param("imgId")
	cNumStr := c.Param("cNum")
	if imgId, err = strconv.Atoi(imgIdStr); err != nil {
		c.String(http.StatusOK, fmt.Sprintf("<script>alert('非法请求')</script><i class='icon icon_favorite_filled'></i><b>%s</b>", cNumStr))
		return
	}
	if account, err = getUser(c); err != nil {
		c.String(http.StatusOK, fmt.Sprintf("<script>alert('请先登录')</script><i class='icon icon_favorite_filled'></i><b>%s</b>", cNumStr))
		return
	}
	if status, err = services.NewAccount().Collect(account.UserId, imgId); err != nil || status == false {
		c.String(http.StatusOK, fmt.Sprintf("<script>alert('"+err.Error()+"')</script><i class='icon icon_favorite_filled'></i><b>%s</b>", cNumStr))
		return
	}
	num, _ := strconv.Atoi(cNumStr)
	c.String(http.StatusOK, fmt.Sprintf("script:$('.favorite_button').addClass('pure-button-disabled').find('b').html(%d);", num+1))
}

//收藏列表
func Favorites(c *gin.Context) {
	var (
		account *services.AccountAuth
		err     error
	)
	if account, err = getUser(c); err != nil {
		query := url.QueryEscape("/accounts/favorites")
		c.Redirect(http.StatusFound, "/accounts/login/?next="+query)
	}
	list := services.NewAccount().Favorites(account.UserPxId)
	c.HTML(http.StatusOK, "favorites.html", gin.H{
		"frontDomain": configs.STATIC_DOMAIN,
		"account":     account,
		"list":        list,
	})
}

//删除收藏
func DelFavorite(c *gin.Context) {
	var (
		account *services.AccountAuth
		imgId   int
		err     error
	)
	imgIdStr := c.Param("imgId")
	if imgId, err = strconv.Atoi(imgIdStr); err != nil {
		c.String(http.StatusOK, fmt.Sprintf("<script>alert('非法请求')</script>"))
		return
	}
	if account, err = getUser(c); err != nil {
		c.String(http.StatusOK, fmt.Sprintf("<script>alert('请先登录')</script>"))
		return
	}
	if status, err := services.NewAccount().DelFavorite(account.UserId, imgId); status == false || err != nil {
		c.String(http.StatusOK, fmt.Sprintf("<script>alert('"+err.Error()+"')</script>"))
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("script:$('.photo-%d').removeClass('has_favorited').find('.icon_favorite').closest('.ajax').html('<i class='icon icon_favorite'></i> 609');", imgId))
}

//评论
func Comment(c *gin.Context) {

}

//我的图片
func Media(c *gin.Context) {

}

//上传
func Upload(c *gin.Context) {

}
