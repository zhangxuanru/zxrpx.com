/*
@Time : 2020/7/23 14:45
@Author : zxr
@File : search
@Software: GoLand
*/
package controllers

import (
	"fmt"
	"math"
	"net/http"
	"pix/application/services"
	"pix/configs"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//搜索
func PhotoSearch(c *gin.Context) {
	var (
		limit = 5
		page  int
		err   error
	)
	query := strings.TrimSpace(c.Param("text"))
	if len(query) == 0 {
		c.Redirect(http.StatusFound, "/404")
	}
	if len(query) > 10 {
		query = string([]rune(query)[:10])
	}
	pageStr := c.DefaultQuery("page", "1")
	if page, err = strconv.Atoi(pageStr); err != nil || page > 500 || page < 1 {
		page = 1
	}
	offset := (page - 1) * limit
	list, tagList, total, err := services.NewElastic().SearchTag(query, offset, limit)
	if err != nil {
		c.Redirect(http.StatusFound, "/404")
	}
	//todo .....
	for _, v := range list {
		fmt.Printf("%+v\n\n", v)
	}
	totalPage := int(math.Ceil(float64(total) / float64(limit)))
	isNextPage := totalPage-page >= 1
	account, _ := getUser(c)
	c.HTML(http.StatusOK, "search.html", gin.H{
		"list":        list,
		"tagList":     tagList,
		"total":       total,
		"account":     account,
		"frontDomain": configs.STATIC_DOMAIN,
		"query":       query,
		"page":        page,
		"prevPage":    page - 1,
		"nextPage":    page + 1,
		"totalPage":   totalPage,
		"isNextPage":  isNextPage,
		"baseUrl":     "/photo/search/" + query + "/",
	})
}
