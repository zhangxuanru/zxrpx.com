/*
@Time : 2020/6/29 17:10
@Author : zxr
@File : photoutil
@Software: GoLand
*/
package services

import "pix/application/models"

//组合图片属性和图片tag 数据
func (p *PicService) combinePicAttrTagData(picList []models.Picture) (result []*PhotoResult) {
	var (
		attrRes    attrMapRes
		tagRes     tagMapRes
		tagService = NewTagService()
		idList     = p.getPxIdByPicList(picList)
	)
	result = make([]*PhotoResult, len(picList))
	if attrRes = p.GetPicAttrListByIds(idList); len(attrRes) == 0 {
		return
	}
	tagRes = tagService.GetTagListByPicIds(idList)
	for key, pic := range picList {
		if _, ok := attrRes[pic.PxImgId]; !ok || len(attrRes[pic.PxImgId]) == 0 {
			continue
		}
		result[key] = &PhotoResult{
			Id:           pic.Id,
			Uid:          pic.Uid,
			CategoryId:   pic.CategoryId,
			PxImgId:      pic.PxImgId,
			ImageType:    pic.ImageType,
			LikeNum:      pic.LikeNum,
			FavoritesNum: pic.FavoritesNum,
			CommentsNum:  pic.CommentsNum,
			Attr:         attrRes[pic.PxImgId],
			Tags:         tagRes[pic.PxImgId],
			TagStr:       tagService.GetTagStrByTagModels(tagRes[pic.PxImgId]),
		}
	}
	return result
}

//组合图片属性和图片tag和user数据
func (p *PicService) comBindPicUserResult(picList []models.Picture) (result []*PhotoResult) {
	result = p.combinePicAttrTagData(picList)
	//添加用户信息
	uidList := p.GetUserByPicList(picList)
	if len(uidList) == 0 {
		return
	}
	userData := NewUserService().GetUserDataByUidList(uidList)
	for _, item := range result {
		if _, ok := userData[item.Uid]; ok {
			item.User = userData[item.Uid]
		} else {
			item.User = models.User{}
		}
	}
	return
}

//获取图片ID列表
func (p *PicService) getPxIdByPicList(picList []models.Picture) []int {
	idList := make([]int, len(picList))
	for key, pic := range picList {
		idList[key] = pic.PxImgId
	}
	return idList
}

//根据图片列表获取用户资料
func (p *PicService) GetUserByPicList(picList []models.Picture) []int {
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
