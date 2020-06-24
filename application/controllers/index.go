/*
@Time : 2020/6/23 15:23
@Author : zxr
@File : index
@Software: GoLand
*/
package controllers

import (
	"net/http"
	"pix/application/services"
	"pix/configs"
	"strconv"

	"github.com/gin-gonic/gin"
)

//首页
func Index(c *gin.Context) {
	var (
		page int
		err  error
	)
	pageStr := c.DefaultQuery("page", "1")
	if page, err = strconv.Atoi(pageStr); err != nil {
		page = 1
	}
	picList, count := services.NewPicService(page, 10).GetPicListByAddTimeOrder()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"frontDomain": configs.STATIC_DOMAIN,
		"picList":     picList,
		"count":       count,
	})
}

func Test(c *gin.Context) {
	c.HTML(http.StatusOK, "src.html", gin.H{
		"frontDomain": configs.STATIC_DOMAIN})
}
