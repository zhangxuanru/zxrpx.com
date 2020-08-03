/*
@Time : 2020/7/15 14:06
@Author : zxr
@File : users
@Software: GoLand
*/
package controllers

import (
	"math"
	"net/http"
	"pix/application/services"
	"pix/configs"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//我的资料
func Profile(c *gin.Context) {
	var (
		account  *services.AccountAuth
		isFollow bool
		getUid   int
		err      error
		page     int
		limit    = 100
	)
	uidStr := strings.TrimLeft(c.Param("user"), "nick-")
	if len(uidStr) == 0 {
		c.Redirect(http.StatusFound, "/404")
	}
	if getUid, err = strconv.Atoi(uidStr); err != nil {
		c.Redirect(http.StatusFound, "/404")
	}
	tab := c.DefaultQuery("tab", "popular")
	pageStr := c.DefaultQuery("page", "1")
	if page, err = strconv.Atoi(pageStr); err != nil || page > 500 || page < 1 {
		page = 1
	}
	userData, photoList, count := services.NewUserService().Profile(getUid, tab, page, limit)
	if userData == nil {
		c.Redirect(http.StatusFound, "/404")
	}
	totalPage := int(math.Ceil(float64(count) / float64(limit)))
	isNextPage := totalPage-page >= 1
	account, _ = getUser(c)
	if account.UserId > 0 {
		isFollow = services.NewAccount().ExistsFollow(account.UserId, userData.PxUid)
	}
	c.HTML(http.StatusOK, "profile.html", gin.H{
		"account":     account,
		"uid":         getUid,
		"userData":    userData,
		"photoList":   photoList,
		"photoLen":    len(photoList),
		"count":       count,
		"page":        page,
		"prevPage":    page - 1,
		"nextPage":    page + 1,
		"totalPage":   totalPage,
		"isNextPage":  isNextPage,
		"isFollow":    isFollow,
		"frontDomain": configs.STATIC_DOMAIN,
		"tab":         tab,
		"baseUrl":     c.Param("user") + "/?tab=" + tab,
	})
}
