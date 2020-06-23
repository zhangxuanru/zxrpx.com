/*
@Time : 2020/6/23 11:56
@Author : zxr
@File : router
@Software: GoLand
*/
package router

import (
	"pix/application/controllers"
)

//设置路由
func (h *HttpServer) settingRouter() {
	h.engine.GET("/", controllers.Index)
}
