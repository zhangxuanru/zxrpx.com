/*
@Time : 2020/6/23 15:23
@Author : zxr
@File : index
@Software: GoLand
*/
package controllers

import (
	"math"
	"net/http"
	"pix/application/services"
	"pix/configs"
	"strconv"

	"github.com/gin-gonic/gin"
)

//首页
func Index(c *gin.Context) {
	var (
		limit   = 100
		account *services.AccountAuth
		page    int
		err     error
	)
	pageStr := c.DefaultQuery("page", "1")
	if page, err = strconv.Atoi(pageStr); err != nil || page > 500 || page < 1 {
		page = 1
	}
	picList, count := services.NewPicService().SetPageParams(page, limit).GetPicListByAddTimeOrder()
	totalPage := int(math.Ceil(float64(count) / float64(limit)))
	isNextPage := totalPage-page >= 1

	account, _ = getUser(c)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"frontDomain": configs.STATIC_DOMAIN,
		"picList":     picList,
		"count":       count,
		"page":        page,
		"prevPage":    page - 1,
		"nextPage":    page + 1,
		"totalPage":   totalPage,
		"isNextPage":  isNextPage,
		"account":     account,
		"baseUrl":     "/?",
		"searchUrl":   "/photo/search/?",
	})
}
