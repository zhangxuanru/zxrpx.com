/*
@Time : 2020/6/23 15:23
@Author : zxr
@File : index
@Software: GoLand
*/
package controllers

import (
	"pix/application/services"
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
	result := services.NewPicService(page, 10).GetPicListByAddTimeOrder()

	//context.HTML(http.StatusOK, "index.html", gin.H{"frontDomain": configs.STATIC_DOMAIN})
}
