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
		return DefaultPicUrl
	}
	fileName := ""
	maxHeight := 0
	for _, attr := range attrList {
		if attr.Height > maxHeight {
			maxHeight = attr.Height
			fileName = attr.FileName
		}
		if attr.Height == height {
			if Env == EnvTest {
				return attr.ImageURL
			}
			if attr.IsQiniu == 1 {
				return configs.STATIC_CDN_DOMAIN + attr.FileName
			}
			return attr.ImageURL
		}
	}
	url := fmt.Sprintf("%s%s%s%d%s", configs.STATIC_CDN_DOMAIN, fileName, "?imageView2/0/h/", height, "/q/85|imageslim")
	return url
}

//显示头像
func ViewHeadPortrait(fileName string, width, height int) string {
	if fileName == "" || len(fileName) < 3 {
		fileName = "timg.jpg"
	}
	headImgUrl := fmt.Sprintf("%s%s%s%d%s%d%s", configs.STATIC_CDN_DOMAIN, fileName, "?imageView2/1/w/", width, "/h/", height, "/q/80|imageslim")
	return headImgUrl
}

//显示图片,只设置高度
func ViewPicByHeight(fileName string, height int) string {
	url := fmt.Sprintf("%s%s%s%d%s", configs.STATIC_CDN_DOMAIN, fileName, "imageView2/0/h/", height, "/q/85|imageslim")
	return url
}

//如果同时有图片源地址和文件名，判断显示是源地址还是CDN文件地址
func ViewPicAddr(srcUrl, fileName string) string {
	if Env == EnvTest && srcUrl != "" {
		return srcUrl
	}
	return configs.STATIC_CDN_DOMAIN + fileName
}
