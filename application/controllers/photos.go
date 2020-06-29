/*
@Time : 2020/6/29 15:28
@Author : zxr
@File : photos
@Software: GoLand
*/
package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//图片详情页
func Detail(c *gin.Context) {
	var (
		pxId int
		err  error
	)
	idStr := c.Param("id")
	if pxId, err = strconv.Atoi(idStr); err != nil || pxId == 0 {
		c.Redirect(http.StatusFound, "/404")
		return
	}

}
