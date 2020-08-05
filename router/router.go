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

	h.engine.GET("/accounts/login", controllers.Login)            //登录的HTML
	h.engine.POST("/accounts/loginDo", controllers.LoginDo)       //执行登录
	h.engine.GET("/accounts/logout", controllers.Logout)          //注销
	h.engine.GET("/accounts/register", controllers.Register)      //注册
	h.engine.POST("/accounts/registerDo", controllers.RegisterDo) //注册

	h.engine.GET("/accounts/settings/", controllers.Settings)               //设置个人资料HTML
	h.engine.POST("/accounts/settingsDo/", controllers.SettingsDo)          //设置个人资料
	h.engine.POST("/accounts/change_password/", controllers.ChangePassword) //修改密码

	h.engine.GET("/accounts/following/", controllers.Following)      //关注列表
	h.engine.POST("/accounts/follow/:authorId/", controllers.Follow) //关注

	h.engine.POST("/accounts/collect/:imgId/:cNum", controllers.Collect)    //收藏
	h.engine.GET("/accounts/favorites", controllers.Favorites)              //收藏列表
	h.engine.POST("/accounts/delfavorite/:imgId/", controllers.DelFavorite) //删除收藏
	h.engine.POST("/images/like/:imgId/:lNum", controllers.Like)            //喜欢
	h.engine.POST("/images/comment/:imgId", controllers.Comment)            //提交评论

	h.engine.GET("/accounts/upload", controllers.Upload) //上传图片
	h.engine.GET("/accounts/media", controllers.Media)   //我的图片

	h.engine.GET("/users/:user/", controllers.Profile)          //我的资料
	h.engine.GET("/photo/search/", controllers.PhotoSearch)     //搜索
	h.engine.GET("/photo/search/:key", controllers.PhotoSearch) //搜索
	h.engine.GET("/image/search/:key", controllers.PhotoSearch) //搜索

	//h.engine.Any()

	h.engine.GET("/404", controllers.NotFound) //404页面
}
