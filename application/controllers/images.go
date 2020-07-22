/*
@Time : 2020/7/3 18:57
@Author : zxr
@File : images
@Software: GoLand
*/
package controllers

import (
	"fmt"
	"net/http"
	"net/url"
	"pix/application/logic"
	"pix/application/services"
	"pix/configs"
	"strconv"

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

//喜欢
func Like(c *gin.Context) {
	var (
		account *services.AccountAuth
		imgId   int
		err     error
	)
	imgIdStr := c.Param("imgId")
	lNumStr := c.Param("lNum")
	if imgId, err = strconv.Atoi(imgIdStr); err != nil {
		c.String(http.StatusOK, fmt.Sprintf("<script>alert('非法请求')</script><i class='icon icon_like_filled'></i><b>%s</b>", lNumStr))
		return
	}
	if account, err = getUser(c); err != nil || account.UserId == 0 {
		c.String(http.StatusOK, fmt.Sprintf("<script>alert('请先登录')</script><i class='icon icon_like_filled'></i><b>%s</b>", lNumStr))
		return
	}
	if _, err = services.NewPicService().Like(account.UserId, imgId, 1); err != nil {
		c.String(http.StatusOK, fmt.Sprintf("<script>alert('%s')</script><i class='icon icon_like_filled'></i><b>%s</b>", err.Error(), lNumStr))
		return
	}
	num, _ := strconv.Atoi(lNumStr)
	c.String(http.StatusOK, fmt.Sprintf("script:$('.like_button').addClass('pure-button-disabled').find('b').html('%d');", num+1))
}
