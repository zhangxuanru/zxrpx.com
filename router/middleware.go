/*
@Time : 2020/6/23 15:56
@Author : zxr
@File : middleware
@Software: GoLand
*/
package router

import (
	"github.com/gin-gonic/gin"
)

//设置中间件
func (h *HttpServer) setMiddleware() {
	//h.engine.Use(h.JWTAuth())
}

//获取当前登录用户信息
func (h *HttpServer) JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
