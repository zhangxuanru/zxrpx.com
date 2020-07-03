/*
@Time : 2020/7/3 18:57
@Author : zxr
@File : images
@Software: GoLand
*/
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//下载图片
func Download(c *gin.Context) {
	//如果query中有attachment，则表示是下载， 如果没有query  则表示是下载
	file := c.Param("file")
	str := "file=" + file
	c.String(http.StatusOK, str)
}
