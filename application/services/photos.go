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

var Fields = "id,uid,px_img_id,category_id,like_num,view_num,downloads_num,favorites_num,comments_num,image_type,image_format,add_time"

func NewPicService() *PicService {
	return &PicService{}
}

//设置分页参数
func (p *PicService) SetPageParams(page int, limit int) *PicService {
	if page < 1 {
		page = 1
	}
	p.Page = page
	p.Limit = limit
	return p
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

//按添加时间和uid获取图片列表
func (p *PicService) GetLatestPhotoList(page int, limit int, uid int, order string) (result []*PhotoResult, count int) {
	var picList []models.Picture
	offset := (page - 1) * limit
	params := &models.QueryParams{
		Offset:  offset,
		Limit:   limit,
		IsCount: true,
		Fields:  Fields,
		Order:   order,
		Where:   "state=1",
	}
	if uid > 0 {
		params.Where += fmt.Sprintf(" AND px_uid=%d", uid)
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

//根据图片名更新图片表与用户统计表的下载量
func (p *PicService) UpdateUserPhotoDownNumByFile(file string) (err error) {
	picId := p.GetPxIdByFile(file)
	if _, err = p.incrByDownNum(picId, 1); err != nil {
		return
	}
	photo := models.NewPicture().GetPxUidByPicId(picId)
	_, err = NewUserService().IncrByUserDownNum(int(photo.PxUid), 1)
	return
}

//根据图片ID更新图片表与用户统计表的浏览量
func (p *PicService) UpdateUserPhotoViewNumByPicId(picId int) (err error) {
	if _, err = p.incrByViewNum(picId, 1); err != nil {
		return
	}
	photo := models.NewPicture().GetPxUidByPicId(picId)
	_, err = models.NewUserStat().IncrByUserViewNum(int(photo.PxUid), 1)
	return
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

//根据文件名获取图片picId
func (p *PicService) GetPxIdByFile(file string) int {
	attr := models.NewPictureAttr().GetPxIdByFile(file)
	return attr.PicId
}

//喜欢操作
func (p *PicService) Like(userId, imgId, num int) (status bool, err error) {
	if _, err = models.NewUserLike().AddLike(userId, imgId); err != nil {
		return false, err
	}
	if _, err = models.NewPicture().EditByLikeNum(imgId, num); err != nil {
		return false, err
	}
	pic := models.NewPicture().GetPxUidByPicId(imgId)
	if pic.PxUid > 0 {
		uid := int(pic.PxUid)
		_, _ = models.NewUserStat().IncrByUserLikeNum(uid, 1)
	}
	return true, nil
}

//判断用户是否已经喜欢过某张图片
func (p *PicService) ExistsLike(userId, imgId int) bool {
	like := models.NewUserLike().GetLikeByUid(userId, imgId)
	if like.Id > 0 {
		return true
	}
	return false
}
