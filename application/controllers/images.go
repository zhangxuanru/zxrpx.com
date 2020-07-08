/*
@Time : 2020/7/3 18:57
@Author : zxr
@File : images
@Software: GoLand
*/
package controllers

import (
	"net/http"
	"net/url"
	"pix/application/logic"
	"pix/application/services"
	"pix/configs"

	"github.com/gin-gonic/gin"
)

//查看或下载图片 如果query中有attachment，则表示是下载， 如果没有query  则表示是查看
func Download(c *gin.Context) {
	var (
		body []byte
		err  error
	)
	file := c.Param("file")
	_, ok := c.GetQueryArray("attachment")
	srcUrl := configs.STATIC_CDN_DOMAIN + file
	if ok == false {
		//七牛时间戳防盗链
		c.Redirect(http.StatusSeeOther, srcUrl)
	} else {
		body, err = logic.DownloadPic(srcUrl, nil)
	}
	if err != nil {
		c.Redirect(http.StatusFound, "/404")
	}
	// 更新下载量
	go func() {
		services.NewPicService().UpdateUserPhotoDownNumByFile(file)
	}()
	filename := url.QueryEscape(file) // 防止中文乱码
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.Writer.Header().Add("Content-Disposition", "attachment; filename=\""+filename+"\"")
	_, _ = c.Writer.Write(body)
}

//为图片添加评论
func AddComment(c *gin.Context) {

}

//喜欢
func Like(c *gin.Context) {

}

//收藏
func Collect(c *gin.Context) {

}
