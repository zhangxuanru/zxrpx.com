/*
@Time : 2020/6/23 15:56
@Author : zxr
@File : middleware
@Software: GoLand
*/
package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//设置中间件
func (h *HttpServer) setMiddleware() {
	h.engine.Use(h.JWTAuth())
}

//验证登录
func (h *HttpServer) JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		token := c.Request.Header.Get("token")
		if len(token) > 10 {
			fmt.Println("add1 token")
			c.Request.Header.Set("token", token)
			c.Next()
		}
		if token, err = c.Cookie("token"); err != nil {
			c.Next()
		}
		//todo 还有问题
		if len(token) > 10 {
			fmt.Println("add2 token")
			//c.SetCookie("token", "", -1, "/", configs.COOKIEDOMAIN, false, true)
			c.Request.Header.Set("token", token)
			c.Header("token", token)
			c.Next()
		}
		c.Next()
	}
}
