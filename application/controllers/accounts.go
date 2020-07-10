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

	"github.com/gin-gonic/gin"
)

type UserLogin struct {
	UserName string `form:"username"`
	PassWord string `form:"password"`
	Token    string `form:"token"`
}

var loginToken string

//显示登录模板
func Login(c *gin.Context) {
	loginToken = logic.Md5(logic.GetRandomString(6))
	c.HTML(http.StatusOK, "login.html", gin.H{
		"frontDomain": configs.STATIC_DOMAIN,
		"cdnDomain":   configs.STATIC_CDN_DOMAIN,
		"token":       loginToken,
	})
}

//登录
func LoginDo(c *gin.Context) {
	var (
		login UserLogin
		err   error
	)
	if err = c.ShouldBind(&login); err != nil {

	}
	c.JSON(http.StatusOK, login)
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
