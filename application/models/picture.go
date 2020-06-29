/*
@Time : 2020/6/12 17:32
@Author : zxr
@File : picture
@Software: GoLand
*/
package models

import (
	"time"
)

//图片主表
type Picture struct {
	Id             int       `gorm:"primary_key; AUTO_INCREMENT; comment:'自增ID'" json:"id"`
	Uid            int       `gorm:"not null; comment:'用户ID'" json:"uid"`
	PxUid          int64     `gorm:"not null; comment:'px 用户ID'" json:"px_uid"`
	PxImgId        int       `gorm:"unique; not null; comment:'px站的图片ID'" json:"px_img_id"`
	PageURL        string    `gorm:"type:varchar(100); NOT NULL; comment:'网页页面地址'" json:"page_url"`
	CategoryId     int       `gorm:"index:category_id; not null; comment:'分类ID'" json:"category_id"`
	ImageFormat    int       `gorm:"type:TINYINT(1); NOT NULL;default:0; comment:'图片格式 1:jpg 2:png 0:其它'" json:"image_format"`
	ThemeColor     string    `gorm:"type:varchar(20); comment:'主体颜色'" json:"theme_color"`
	ImageDirection int       `gorm:"type:TINYINT(1); NOT NULL;default:1; comment:'图像方向 1:未区分 2:水平 3:垂直'" json:"image_direction"`
	ImageType      int       `gorm:"type:TINYINT(1); NOT NULL;default:0; comment:'图片类型 0:未区分 1:照片 2:插图 3:矢量'" json:"image_type"`
	ViewNum        uint      `gorm:"NOT NULL; default:0; comment:'总浏览次数'" json:"view_num"`
	DownloadsNum   uint      `gorm:"NOT NULL; default:0; comment:'下载总数'" json:"downloads_num"`
	LikeNum        uint      `gorm:"NOT NULL; default:0; comment:'喜欢总数'" json:"like_num"`
	FavoritesNum   uint      `gorm:"NOT NULL; default:0; comment:'收藏总数'" json:"favorites_num"`
	CommentsNum    uint      `gorm:"NOT NULL; default:0; comment:'评论总数'" json:"comments_num"`
	Sort           uint      `gorm:"NOT NULL; default:0; comment:'排序字段,从小到大'" json:"sort"`
	IsHandpick     int       `gorm:"type:TINYINT(1); NOT NULL;default:0; comment:'是否精选 1:精选 0:不是精选'" json:"is_hand_pick"`
	State          int       `gorm:"type:TINYINT(1); NOT NULL;default:1; comment:'状态 1:状态正常 0:删除'" json:"state"`
	AddTime        time.Time `gorm:"comment:'添加时间'" json:"add_time"`
	UpdateTime     time.Time `gorm:"comment:'修改时间'" json:"update_time"`
}

func NewPicture() *Picture {
	return &Picture{}
}

func (p *Picture) GetPicListByParams(params *QueryParams) (pic []Picture, count int) {
	params = setDefaultParams(params)
	if params.IsCount {
		GetDB().Model(p).Where(params.Where).Count(&count)
	}
	db := GetDB().Select(params.Fields)
	if params.Where != "" {
		db = db.Where(params.Where)
	}
	if params.Order != "" {
		db = db.Order(params.Order)
	}
	db.Offset(params.Offset).Limit(params.Limit).Find(&pic)
	return
}
