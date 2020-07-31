/*
@Time : 2020/7/23 14:45
@Author : zxr
@File : search
@Software: GoLand
*/
package controllers

import (
	"math"
	"net/http"
	"pix/application/logic"
	"pix/application/services"
	"pix/configs"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

//搜索
func PhotoSearch(c *gin.Context) {
	var (
		limit  = 60
		err    error
		search services.SearchPhoto
	)
	if err = c.ShouldBindQuery(&search); err != nil {
		logrus.Error("search ShouldBindQuery error:", err)
	}
	if key := c.Param("key"); len(key) > 0 {
		search.KeyWord = key
	}
	if len(search.KeyWord) > 10 {
		search.KeyWord = string([]rune(search.KeyWord)[:10])
	}
	if search.Page > 100 || search.Page < 1 {
		search.Page = 1
	}
	offset := (search.Page - 1) * limit
	list, tagList, total, err := services.NewElastic().SearchQuery(&search, offset, limit)
	template := "search.html"
	if err != nil || total == 0 {
		template = "not.html"
	}
	totalPage := int(math.Ceil(float64(total) / float64(limit)))
	isNextPage := totalPage-search.Page >= 1
	account, _ := getUser(c)
	baseUrl := logic.BuildSearchUrl(&search)
	c.HTML(http.StatusOK, template, gin.H{
		"search":      search,
		"list":        list,
		"tagList":     tagList,
		"total":       total,
		"account":     account,
		"frontDomain": configs.STATIC_DOMAIN,
		"query":       search.KeyWord,
		"page":        search.Page,
		"prevPage":    search.Page - 1,
		"nextPage":    search.Page + 1,
		"totalPage":   totalPage,
		"isNextPage":  isNextPage,
		"baseUrl":     baseUrl,
		"searchUrl":   baseUrl,
	})
}
