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
	h.engine.GET("/", controllers.Index)                         //首页
	h.engine.GET("/photos/:id", controllers.Detail)              //图片详情页
	h.engine.GET("/images/download/:file", controllers.Download) //查看下载图片

	h.engine.GET("/accounts/comment", controllers.Comment)         //提交评论
	h.engine.GET("/accounts/like", controllers.Like)               //喜欢
	h.engine.GET("/accounts/collect", controllers.Collect)         //收藏
	h.engine.GET("/accounts/follow/:authorId", controllers.Follow) //关注

	h.engine.GET("/accounts/login", controllers.Login)            //登录的HTML
	h.engine.POST("/accounts/loginDo", controllers.LoginDo)       //执行登录
	h.engine.GET("/accounts/logout", controllers.Logout)          //注销
	h.engine.GET("/accounts/register", controllers.Register)      //注册
	h.engine.POST("/accounts/registerDo", controllers.RegisterDo) //注册

	//h.engine.Any()

	h.engine.GET("/404", controllers.NotFound) //404页面
}
