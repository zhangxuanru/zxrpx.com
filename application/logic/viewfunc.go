/*
@Time : 2020/6/28 16:52
@Author : zxr
@File : viewfunc
@Software: GoLand
*/
package logic

import (
	"fmt"
	"pix/application/models"
	"pix/configs"
)

const (
	Env           = EnvTest
	EnvTest       = "test"
	DefaultPicUrl = ""
)

//首页模板宽高设置
func ViewWHAttr(index int, style string) int {
	viewAttr := GetViewAttrData()
	attr := viewAttr[index]
	if style == "w" {
		return attr.Width
	}
	return attr.Height
}

//根据图片属性和宽度获取图片地址
func ViewImageAddr(attrList []models.PictureAttr, height int) string {
	if len(attrList) == 0 {
		//这里可以返回一张默认图片
		return DefaultPicUrl
	}
	for _, attr := range attrList {
		if attr.Height == height {
			if Env == EnvTest {
				//test start
				return attr.ImageURL
				//test end
			}
			if attr.IsQiniu == 1 {
				return configs.STATIC_CDN_DOMAIN + attr.FileName
			}
			return attr.ImageURL
		}
	}
	return DefaultPicUrl
}

//显示头像
func ViewHeadPortrait(fileName string, width, height int) string {
	headImgUrl := fmt.Sprintf("%s%s%s%d%s%d%s", configs.STATIC_CDN_DOMAIN, fileName, "?imageView2/1/w/", width, "/h/", height, "/q/80|imageslim")
	return headImgUrl
}
