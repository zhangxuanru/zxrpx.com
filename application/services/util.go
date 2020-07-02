/*
@Time : 2020/7/2 15:29
@Author : zxr
@File : util
@Software: GoLand
*/
package services

import "pix/application/models"

//获取图片ID列表
func getPxIdByPicList(picList []models.Picture) []int {
	idList := make([]int, len(picList))
	for key, pic := range picList {
		idList[key] = pic.PxImgId
	}
	return idList
}

//根据图片列表获取用户资料
func getUserByPicList(picList []models.Picture) []int {
	uidMap := make(map[int]struct{})
	for _, pic := range picList {
		if _, ok := uidMap[pic.Uid]; ok {
			continue
		}
		uidMap[pic.Uid] = struct{}{}
	}
	uidList := make([]int, len(uidMap))
	i := 0
	for uid, _ := range uidMap {
		uidList[i] = uid
		i++
	}
	return uidList
}

//返回对应数字类型返回文件后缀字符串
func getImgExt(format int) string {
	extMap := map[int]string{
		1: "JPG",
		2: "PNG",
	}
	ext, ok := extMap[format]
	if ok {
		return ext
	}
	return "JPG"
}
