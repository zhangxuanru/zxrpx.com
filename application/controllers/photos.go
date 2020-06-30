/*
@Time : 2020/6/29 15:28
@Author : zxr
@File : photos
@Software: GoLand
*/
package controllers

import (
	"fmt"
	"net/http"
	"pix/application/services"
	"pix/configs"
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
	data := services.NewPicService(1, 1).GetPhotosDetailByPxId(pxId)
	if len(data) == 0 {
		c.Redirect(http.StatusFound, "/404")
		return
	}
	photo := data[0]

	fmt.Printf("photo:%+v\n\n", photo)
	fmt.Printf("User:%+v\n\n", photo.User)

	c.HTML(http.StatusOK, "photos.html", gin.H{
		"frontDomain": configs.STATIC_DOMAIN,
		"baseUrl":     "/photos/" + idStr,
		"photo":       photo,
	})
}
