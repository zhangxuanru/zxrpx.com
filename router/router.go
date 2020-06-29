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
	h.engine.GET("/", controllers.Index)            //首页
	h.engine.GET("/photos/:id", controllers.Detail) //图片详情页

	h.engine.GET("/404", controllers.NotFound) //404页面
}
