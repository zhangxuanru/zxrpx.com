/*
@Time : 2020/6/11 18:06
@Author : zxr
@File : init
@Software: GoLand
*/
package models

import (
	"fmt"
	"log"
	"os"
	"pix/configs"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

var Db *gorm.DB

//查询参数结构体
type QueryParams struct {
	Offset  int
	Limit   int
	Order   string
	Where   string
	Fields  string
	IsCount bool
}

//初始化MYSQL
func init() {
	var err error
	source := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", configs.DbUser, configs.DbPassWd, configs.DbHost, configs.DbPort, configs.DbDataBase)
	if Db, err = gorm.Open("mysql", source); err != nil {
		logrus.Fatal("init mysql:", err)
		return
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "px_" + defaultTableName
	}

	Db.SingularTable(true)
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
	//path := fmt.Sprintf("./logs/sql-%d-%d-%d.log", time.Now().Year(), time.Now().Month(), time.Now().Day())
	//file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModePerm)
	//if err == nil {
	Db.LogMode(true)
	//Db.SetLogger(log.New(file, "\r\n", 0))
	Db.SetLogger(log.New(os.Stdout, "\r\n", 0))
	//	}
}

//获取DB对象
func GetDB() *gorm.DB {
	return Db
}

//设置查询条件默认值
func setDefaultParams(params *QueryParams) *QueryParams {
	if params.Fields == "" {
		params.Fields = "id"
	}
	if params.Limit == 0 {
		params.Limit = 10
	}
	return params
}
