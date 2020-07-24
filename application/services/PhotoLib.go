/*
@Time : 2020/6/29 17:10
@Author : zxr
@File : photoutil
@Software: GoLand
*/
package services

import (
	"fmt"
	"pix/application/models"
)

//组合图片属性和图片tag 数据
func (p *PicService) combinePicAttrTagData(picList []models.Picture) (result []*PhotoResult) {
	var (
		attrRes    attrMapRes
		tagRes     tagMapRes
		tagService = NewTagService()
		idList     = getPxIdByPicList(picList)
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
			Id:             pic.Id,
			Uid:            pic.Uid,
			CategoryId:     pic.CategoryId,
			PxImgId:        pic.PxImgId,
			ImageType:      pic.ImageType,
			ImageFormat:    pic.ImageFormat,
			ImageFormatStr: getImgExt(pic.ImageFormat),
			LikeNum:        pic.LikeNum,
			FavoritesNum:   pic.FavoritesNum,
			CommentsNum:    pic.CommentsNum,
			ViewNum:        pic.ViewNum,
			DownloadsNum:   pic.DownloadsNum,
			AddTime:        pic.AddTime,
			Attr:           attrRes[pic.PxImgId],
			Tags:           tagRes[pic.PxImgId],
			TagStr:         tagService.GetTagStrByTagModels(tagRes[pic.PxImgId]),
		}
	}
	return result
}

//组合图片属性和图片tag和user数据
func (p *PicService) comBindPicUserResult(picList []models.Picture) (result []*PhotoResult) {
	result = p.combinePicAttrTagData(picList)
	//添加用户信息
	uidList := getUserByPicList(picList)
	if len(uidList) == 0 {
		return
	}
	userData := NewUserService().GetUserDataByUidList(uidList)
	for _, item := range result {
		if _, ok := userData[item.Uid]; ok {
			item.User = userData[item.Uid]
		} else {
			item.User = nil
		}
	}
	return
}

//更新图片表下载量
func (p *PicService) incrByDownNum(picId, num int) (affected int64, err error) {
	if picId == 0 {
		return
	}
	affected, err = models.NewPicture().IncrByDownNum(picId, num)
	return
}

//更新浏览量
func (p *PicService) incrByViewNum(picId, num int) (affected int64, err error) {
	if picId == 0 {
		return
	}
	affected, err = models.NewPicture().IncrByViewNum(picId, num)
	return
}

//拿用户的第一张图片
func (p *PicService) GetUserFirstPhotoUrl(uid int) (ImageURL string, FileName string) {
	var attr models.PictureAttr
	query := &models.QueryParams{
		Offset:  0,
		Limit:   1,
		Fields:  "px_img_id",
		Where:   fmt.Sprintf("px_uid=%d", uid),
		IsCount: false,
	}
	picList, _ := models.NewPicture().GetPicListByParams(query)
	if len(picList) > 0 {
		picData := picList[0]
		attr = models.NewPictureAttr().GetAttrByPidWidth(picData.PxImgId, 486)
	}
	return attr.ImageURL, attr.FileName
}

//根据图片ID获取图片信息
func (p *PicService) GetPicListByIds(idList []int) (result []*PhotoResult) {
	var (
		attrRes    attrMapRes
		tagRes     tagMapRes
		tagService = NewTagService()
		picList    []models.Picture
	)
	result = make([]*PhotoResult, len(idList))
	picList = models.NewPicture().GetPicListByIds(idList)
	tagRes = tagService.GetTagListByPicIds(idList)
	attrRes = p.GetPicAttrListByIds(idList)
	for key, pic := range picList {
		if _, ok := attrRes[pic.PxImgId]; !ok || len(attrRes[pic.PxImgId]) == 0 {
			continue
		}
		result[key] = &PhotoResult{
			Id:             pic.Id,
			Uid:            pic.Uid,
			CategoryId:     pic.CategoryId,
			PxImgId:        pic.PxImgId,
			ImageType:      pic.ImageType,
			ImageFormat:    pic.ImageFormat,
			ImageFormatStr: getImgExt(pic.ImageFormat),
			LikeNum:        pic.LikeNum,
			FavoritesNum:   pic.FavoritesNum,
			CommentsNum:    pic.CommentsNum,
			ViewNum:        pic.ViewNum,
			DownloadsNum:   pic.DownloadsNum,
			AddTime:        pic.AddTime,
			Attr:           attrRes[pic.PxImgId],
			Tags:           tagRes[pic.PxImgId],
			TagStr:         tagService.GetTagStrByTagModels(tagRes[pic.PxImgId]),
		}
	}
	return result
}
