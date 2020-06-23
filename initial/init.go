/*
@Time : 2020/6/23 14:56
@Author : zxr
@File : init
@Software: GoLand
*/
package initial

import (
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func init() {
	initLog()
}

func initLog() {
	logrus.SetFormatter(&nested.Formatter{
		TimestampFormat: time.RFC3339,
	})
	logrus.SetLevel(logrus.DebugLevel)
	//logrus.SetReportCaller(true)

	//日志切割
	path := "./logs/error.log"
	writer, err := rotatelogs.New(
		path+".%Y%m%d",
		rotatelogs.WithLinkName(path),             // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(5*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		logrus.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}

	pathInfo := "./logs/info.log"
	writerInfo, _ := rotatelogs.New(
		pathInfo+".%Y%m%d",
		rotatelogs.WithLinkName(pathInfo),         // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(5*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)

	debug := "./logs/debug.log"
	writerDebug, _ := rotatelogs.New(
		debug+".%Y%m%d",
		rotatelogs.WithLinkName(debug),            // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(5*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writerDebug, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writerInfo,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, nil)
	logrus.AddHook(lfHook)
}
