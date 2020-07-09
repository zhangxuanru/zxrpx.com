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
	"pix/application/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

}

func Logout(c *gin.Context) {

}

func Register(c *gin.Context) {

}

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
