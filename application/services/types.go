/*
@Time : 2020/7/2 15:31
@Author : zxr
@File : types
@Software: GoLand
*/
package services

import (
	"pix/application/models"
	"time"
)

type PicService struct {
	Page  int
	Limit int
}

type attrMapRes map[int][]models.PictureAttr

type PhotoResult struct {
	Id             int       `json:"id"`
	Uid            int       `json:"uid"`
	PxImgId        int       `json:"px_img_id"`
	CategoryId     int       `json:"category_id"`
	ImageType      int       `json:"image_type"`
	ImageFormat    int       `json:"image_format"`
	ImageFormatStr string    `json:"image_format_str"`
	LikeNum        uint      `json:"like_num"`
	FavoritesNum   uint      `json:"favorites_num"`
	CommentsNum    uint      `json:"comments_num"`
	ViewNum        uint      `json:"view_num"`
	DownloadsNum   uint      `json:"downloads_num"`
	AddTime        time.Time `json:"add_time"`
	TagStr         string
	Attr           []models.PictureAttr
	Tags           []models.Tag
	User           *UserStat
}
