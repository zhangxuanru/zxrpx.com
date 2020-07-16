/*
@Time : 2020/7/15 14:06
@Author : zxr
@File : users
@Software: GoLand
*/
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//我的资料
func Profile(c *gin.Context) {
	c.String(http.StatusOK, "hello world")
}
