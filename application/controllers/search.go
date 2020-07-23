/*
@Time : 2020/7/23 14:45
@Author : zxr
@File : search
@Software: GoLand
*/
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//搜索 //todo 明天继续  搜索
func PhotoSearch(c *gin.Context) {
	query := c.Param("text")
	c.String(http.StatusOK, query)
}
