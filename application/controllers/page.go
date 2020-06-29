/*
@Time : 2020/6/29 16:23
@Author : zxr
@File : page
@Software: GoLand
*/
package controllers

import (
	"net/http"
	"pix/configs"

	"github.com/gin-gonic/gin"
)

//404 page
func NotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", gin.H{
		"frontDomain": configs.STATIC_DOMAIN,
	})
	return
}
