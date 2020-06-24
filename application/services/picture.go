/*
@Time : 2020/6/23 18:00
@Author : zxr
@File : picture
@Software: GoLand
*/
package services

import (
	"pix/application/models"
)

type PicService struct {
	Page  int
	Limit int
}

type attrMapRes map[int][]models.PictureAttr

type PicResult struct {
	Id           int   `json:"id"`
	Uid          int64 `json:"uid"`
	PxImgId      int   `json:"px_img_id"`
	ImageType    int   `json:"image_type"`
	LikeNum      uint  `json:"like_num"`
	FavoritesNum uint  `json:"favorites_num"`
	CommentsNum  uint  `json:"comments_num"`
	Attr         []models.PictureAttr
	Tags         []models.Tag
}

func NewPicService(page int, limit int) *PicService {
	if page < 1 {
		page = 1
	}
	return &PicService{
		Page:  page,
		Limit: limit,
	}
}

//根据添加时间查询图片列表
func (p *PicService) GetPicListByAddTimeOrder() (result []*PicResult, count int) {
	var (
		picList []models.Picture
	)
	offset := (p.Page - 1) * p.Limit
	params := &models.QueryParams{
		Offset: offset,
		Limit:  p.Limit,
		Fields: "id,uid,px_img_id,like_num,favorites_num,comments_num,image_type",
		Order:  "add_time DESC,like_num DESC",
		Where:  "state=1",
	}
	if picList, count = models.NewPicture().GetPicListByParams(params); len(picList) == 0 {
		return
	}
	return p.CombinePicData(picList), count
}

//组合图片数据
func (p *PicService) CombinePicData(picList []models.Picture) (result []*PicResult) {
	var (
		attrRes attrMapRes
		tagRes  tagMapRes
		idList  []int
	)
	idList = p.GetIdList(picList)
	if attrRes = p.GetPicAttrListByIds(idList); len(attrRes) == 0 {
		return
	}
	tagRes = NewTagService().GetTagListByPicIds(idList)
	result = make([]*PicResult, len(picList))
	for key, pic := range picList {
		result[key] = &PicResult{
			Id:           pic.Id,
			Uid:          pic.Uid,
			PxImgId:      pic.PxImgId,
			ImageType:    pic.ImageType,
			LikeNum:      pic.LikeNum,
			FavoritesNum: pic.FavoritesNum,
			CommentsNum:  pic.CommentsNum,
			Attr:         attrRes[pic.PxImgId],
			Tags:         tagRes[pic.PxImgId],
		}
	}
	return result
}

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

//获取图片ID列表
func (p *PicService) GetIdList(picList []models.Picture) []int {
	idList := make([]int, len(picList))
	for key, pic := range picList {
		idList[key] = pic.PxImgId
	}
	return idList
}
