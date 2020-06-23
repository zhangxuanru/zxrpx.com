/*
@Time : 2020/6/23 15:23
@Author : zxr
@File : index
@Software: GoLand
*/
package controllers

import (
	"pix/application/services"

	"github.com/gin-gonic/gin"
)

//首页
func Index(context *gin.Context) {
	services.NewPicService(1).GetPicListByAddTimeOrder()
	//context.HTML(http.StatusOK, "index.html", gin.H{"frontDomain": configs.STATIC_DOMAIN})
}
