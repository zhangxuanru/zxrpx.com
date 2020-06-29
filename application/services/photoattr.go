/*
@Time : 2020/6/29 17:09
@Author : zxr
@File : photoattr
@Software: GoLand
*/
package services

import "pix/application/models"

//根据图片ID查询图片属性信息
func (p *PicService) GetPicAttrListByIds(idList []int) (result attrMapRes) {
	var attrList []models.PictureAttr
	if attrList = models.NewPictureAttr().GetListByPicIds(idList); len(attrList) == 0 {
		return
	}
	result = make(attrMapRes)
	for _, attr := range attrList {
		result[attr.PicId] = append(result[attr.PicId], attr)
	}
	return
}
