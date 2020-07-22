/*
@Time : 2020/6/29 15:28
@Author : zxr
@File : photos
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

//图片详情页
func Detail(c *gin.Context) {
	var (
		pxId       int
		pxIdStr    string
		err        error
		picData    []*services.PhotoResult
		account    *services.AccountAuth
		commentNum = 3
		isLike     bool //是否喜欢
		isCollect  bool //是否收藏
		isFollow   bool //是否关注
	)
	pxIdStr = c.Param("id")
	if pxId, err = strconv.Atoi(pxIdStr); err != nil || pxId == 0 {
		c.Redirect(http.StatusFound, "/404")
	}
	//获取图片基础信息
	if picData = services.NewPicService().SetPageParams(1, 1).GetPhotosDetailByPxId(pxId); len(picData) == 0 {
		c.Redirect(http.StatusFound, "/404")
	}
	photo := picData[0]
	//获取评论
	commentsList, commentCount := services.NewComments().GetCommentsByPicId(pxId, 1, commentNum)
	//获取近期图像
	recentImages := services.NewUserService().GetRecentImagesByUid(photo.User.PxUid, 12, pxId)
	//获取推荐标签
	tagList := services.NewTagService().GetRandOffsetTagList(16)
	account, _ = getUser(c)
	//获取喜欢，收藏，关注信息
	if account.UserId > 0 {
		isCollect = services.NewAccount().ExistsCollect(account.UserId, pxId)
		isFollow = services.NewAccount().ExistsFollow(account.UserId, photo.User.PxUid)
		isLike = services.NewPicService().ExistsLike(account.UserId, pxId)
	}
	//更新浏览量
	go func() {
		services.NewPicService().UpdateUserPhotoViewNumByPicId(pxId)
	}()

	c.HTML(http.StatusOK, "photos.html", gin.H{
		"frontDomain":  configs.STATIC_DOMAIN,
		"baseUrl":      "/photos/" + pxIdStr,
		"photo":        photo,
		"commentsList": commentsList,
		"commentCount": commentCount - commentNum,
		"commentNum":   commentNum,
		"recentImages": recentImages,
		"tagList":      tagList,
		"account":      account,
		"isLike":       isLike,
		"isCollect":    isCollect,
		"isFollow":     isFollow,
		"pxIdStr":      pxIdStr,
	})
}
