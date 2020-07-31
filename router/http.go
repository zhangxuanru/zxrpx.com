/*
@Time : 2020/6/23 14:31
@Author : zxr
@File : http
@Software: GoLand
*/
package router

import (
	"context"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"pix/application/logic"
	"pix/configs"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	server *http.Server
	engine *gin.Engine
}

func init() {
	gin.SetMode(gin.DebugMode)
}

func NewHttpServer() *HttpServer {
	server := &HttpServer{
		engine: gin.Default(),
	}
	server.setMiddleware()
	server.settingRouter()
	server.LoadTemplateFunc()
	server.LoadHtml()
	server.LoadStatic()
	return server
}

//启动服务
func (h *HttpServer) StartServer() {
	h.server = &http.Server{
		Addr:           configs.HTTP_PORT,
		Handler:        h.engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		logrus.Println("启动HTTP服务，端口为:", configs.HTTP_PORT)
		if err := h.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatal("ListenAndServe err:", err)
		}
	}()
	logrus.Println("进程ID=", os.Getpid())
}

//监听信号关闭服务
func (h *HttpServer) MonitorServer() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT)
	<-quit
	logrus.Println("start shutdown server ...")
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	if err := h.server.Shutdown(ctx); err != nil {
		logrus.Fatal("server Shutdown:", err)
	}
	logrus.Println("server exiting....")
}

//载入模板文件
func (h *HttpServer) LoadHtml() {
	h.engine.LoadHTMLGlob("application/views/tpl/**/*")
}

//载入静态文件
func (h *HttpServer) LoadStatic() {
	h.engine.StaticFS("/static", http.Dir("public"))
}

//模板函数
func (h *HttpServer) LoadTemplateFunc() {
	h.engine.SetFuncMap(template.FuncMap{
		"ViewWH":           logic.ViewWHAttr,
		"ViewSearchWH":     logic.ViewSearchWHAttr,
		"ViewImageAddr":    logic.ViewImageAddr,
		"ViewHeadPortrait": logic.ViewHeadPortrait,
		"ViewPicByHeight":  logic.ViewPicByHeight,
		"ViewPicAddr":      logic.ViewPicAddr,
		"ReplaceUrlParam":  logic.ReplaceUrlParam,
		"RemoveUrlParam":   logic.RemoveUrlParam,
	})
}
