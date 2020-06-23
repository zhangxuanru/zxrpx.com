/*
@Time : 2020/6/22 19:08
@Author : zxr
@File : main
@Software: GoLand
*/
package main

import (
	_ "pix/initial"
	"pix/router"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	server := router.NewHttpServer()
	server.StartServer()
	defer server.MonitorServer()
}
