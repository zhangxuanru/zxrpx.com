/*
@Time : 2020/6/28 16:52
@Author : zxr
@File : viewfunc
@Software: GoLand
*/
package logic

import (
	"fmt"
	"net/url"
	"pix/application/models"
	"pix/configs"
	"strings"
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

//搜索页模板宽高设置
func ViewSearchWHAttr(index int, style string) int {
	viewAttr := GetSearchViewAttr()
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

//替换URl中具体的参数和值
func ReplaceUrlParam(urlStr string, key string, val string) string {
	u, err := url.Parse(urlStr)
	if err != nil {
		return urlStr
	}
	par := u.Query().Get(key)
	old := fmt.Sprintf("%s=%s", key, par)
	new := fmt.Sprintf("%s=%s", key, val)
	if par == "" {
		urlStr += "&" + new
	} else {
		urlStr = strings.Replace(urlStr, old, new, 1)
	}
	return urlStr
}

//删除URL参数中某一个
func RemoveUrlParam(urlStr string, key string) string {
	u, err := url.Parse(urlStr)
	if err != nil {
		return urlStr
	}
	par := u.Query().Get(key)
	if par != "" {
		old := fmt.Sprintf("%s=%s", key, par)
		urlStr = strings.Replace(urlStr, old, "", 1)
	}
	return urlStr
}
