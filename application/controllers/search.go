/*
@Time : 2020/7/23 14:45
@Author : zxr
@File : search
@Software: GoLand
*/
package controllers

import (
	"net/http"
	"pix/application/services"
	"pix/configs"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

//搜索
func PhotoSearch(c *gin.Context) {
	query := strings.TrimSpace(c.Param("text"))
	if len(query) == 0 {
		c.Redirect(http.StatusFound, "/404")
	}
	if len(query) > 10 {
		query = string([]rune(query)[:10])
	}
	account, _ := getUser(c)
	list, tagList, total, err := services.NewElastic().SearchTag(query, 0, 10)
	if err != nil {
		c.Redirect(http.StatusFound, "/404")
	}
	logrus.Infof("list:%+v\n\n", list)
	logrus.Infof("total:%d\n\n", total)
	logrus.Infof("err:%v\n\n", err)

	c.HTML(http.StatusOK, "search.html", gin.H{
		"list":        list,
		"tagList":     tagList,
		"total":       total,
		"account":     account,
		"frontDomain": configs.STATIC_DOMAIN,
		"query":       query,
		"page":        1,
		"totalPage":   2,
		"prevPage":    1,
		"nextPage":    2,
	})
}
