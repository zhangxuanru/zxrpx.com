/*
@Time : 2020/6/23 18:00
@Author : zxr
@File : picture  图片接口服务
@Software: GoLand
*/
package services

import (
	"fmt"
	"pix/application/models"
)

type PicService struct {
	Page  int
	Limit int
}

type attrMapRes map[int][]models.PictureAttr

type PhotoResult struct {
	Id           int  `json:"id"`
	Uid          int  `json:"uid"`
	PxImgId      int  `json:"px_img_id"`
	CategoryId   int  `json:"category_id"`
	ImageType    int  `json:"image_type"`
	LikeNum      uint `json:"like_num"`
	FavoritesNum uint `json:"favorites_num"`
	CommentsNum  uint `json:"comments_num"`
	TagStr       string
	Attr         []models.PictureAttr
	Tags         []models.Tag
	User         *UserStat
}

var Fields = "id,uid,px_img_id,category_id,like_num,favorites_num,comments_num,image_type"

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
func (p *PicService) GetPicListByAddTimeOrder() (result []*PhotoResult, count int) {
	var picList []models.Picture
	offset := (p.Page - 1) * p.Limit
	params := &models.QueryParams{
		Offset:  offset,
		Limit:   p.Limit,
		IsCount: true,
		Fields:  Fields,
		Order:   "sort ASC,like_num DESC,add_time DESC",
		Where:   "state=1",
	}
	if picList, count = models.NewPicture().GetPicListByParams(params); len(picList) == 0 {
		return
	}
	return p.combinePicAttrTagData(picList), count
}

//根据图片ID获取图片详情
func (p *PicService) GetPhotosDetailByPxId(pxId int) (result []*PhotoResult) {
	var picList []models.Picture
	params := &models.QueryParams{
		Offset:  0,
		Limit:   1,
		IsCount: false,
		Fields:  Fields,
		Where:   fmt.Sprintf("px_img_id=%d AND state=1", pxId),
	}
	if picList, _ = models.NewPicture().GetPicListByParams(params); len(picList) == 0 {
		result = make([]*PhotoResult, 0)
		return
	}
	result = p.comBindPicUserResult(picList)
	return result
}
