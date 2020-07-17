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

//个人资料设置
type SettingUser struct {
	UserName  string `form:"username" binding:"required"`
	NickName  string `form:"nick_name" binding:"required"`
	FirstName string `form:"first_name"`
	Email     string `form:"email" binding:"required"`
	Intro     string `form:"intro"`
	Facebook  string `form:"facebook" `
	Twitter   string `form:"twitter" `
	Website   string `form:"website" `
	Token     string `form:"token" binding:"required" json:"token"`
}

//修改密码
type UserChangePassword struct {
	OldPassword  string `form:"old_password" binding:"required"`
	NewPassword  string `form:"new_password" binding:"required"`
	ConfPassword string `form:"new_password2" binding:"required"`
}

//关注列表
type FollowList struct {
	list []*Follow
}

type Follow struct {
	UserName      string
	HeadPhotoUrl  string
	HeadPhotoFile string
	PicNum        int
	ViewNum       int
	DownloadsNum  int
	LikeNum       int
	CommentNum    int
	FollowerNum   int
}
