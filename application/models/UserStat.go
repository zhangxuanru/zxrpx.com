/*
@Time : 2020/6/12 14:56
@Author : zxr
@File : userstat
@Software: GoLand
*/
package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//用户数据统计表
type UserStat struct {
	Id           int       `gorm:"primary_key; AUTO_INCREMENT; comment:'自增ID'" json:"id"`
	Uid          int64     `gorm:"unique; not null; comment:'用户ID'" json:"uid"`
	PxUid        int64     `gorm:"unique; not null; comment:'px 用户ID'" json:"px_uid"`
	PicNum       int       `gorm:"type:INT(11); NOT NULL;  default:0;  comment:'图片总数'"  json:"pic_num"`
	ViewNum      int       `gorm:"NOT NULL; default:0; comment:'总浏览次数'" json:"view_num"`
	DownloadsNum int       `gorm:"type:INT(11); NOT NULL;  default:0;  comment:'下载量'"    json:"downloads_num"`
	LikeNum      int       `gorm:"type:INT(11); NOT NULL;  default:0;  comment:'喜欢总量'"  json:"like_num"`
	CommentNum   int       `gorm:"type:INT(11); NOT NULL;  default:0;  comment:'评论总量'"  json:"comment_num"`
	FollowerNum  int       `gorm:"type:INT(11); NOT NULL;  default:0;  comment:'收藏总量'"  json:"follower_num"`
	AddTime      time.Time `gorm:"comment:'添加时间'" json:"add_time"`
	UpdateTime   time.Time `gorm:"comment:'修改时间'" json:"update_time"`
}

func NewUserStat() *UserStat {
	return &UserStat{}
}

//修改用户统计表
func (u *UserStat) UpdateStat() (affected int64, err error) {
	buildMap := map[string]interface{}{
		"pic_num":       gorm.Expr("pic_num + ?", u.PicNum),
		"view_num":      gorm.Expr("view_num + ?", u.ViewNum),
		"downloads_num": gorm.Expr("downloads_num + ?", u.DownloadsNum),
		"like_num":      gorm.Expr("like_num + ?", u.LikeNum),
		"comment_num":   gorm.Expr("comment_num + ?", u.CommentNum),
		"follower_num":  gorm.Expr("follower_num + ?", u.FollowerNum),
	}
	updates := GetDB().Model(u).Where("px_uid = ?", u.PxUid).Updates(buildMap).Omit("add_time")
	return updates.RowsAffected, updates.Error
}

//根据用户ID数组批量获取用户统计信息
func (u *UserStat) GetUserStatByUidList(uidList []int) (userMap map[int]UserStat) {
	if len(uidList) == 0 {
		userMap = make(map[int]UserStat)
		return
	}
	var list []UserStat
	fields := "id,uid,px_uid,pic_num,view_num,downloads_num,like_num,comment_num,follower_num"
	GetDB().Where("uid IN (?) ", uidList).Select(fields).Find(&list)
	userMap = make(map[int]UserStat, len(list))
	for _, user := range list {
		uid := int(user.Uid)
		userMap[uid] = user
	}
	return
}

func (u *UserStat) GetUserStatByPxUidList(uidList []int) (userMap map[int]UserStat) {
	if len(uidList) == 0 {
		userMap = make(map[int]UserStat)
		return
	}
	var list []UserStat
	fields := "id,uid,px_uid,pic_num,view_num,downloads_num,like_num,comment_num,follower_num"
	GetDB().Where("px_uid IN (?) ", uidList).Select(fields).Find(&list)
	userMap = make(map[int]UserStat, len(list))
	for _, user := range list {
		uid := int(user.PxUid)
		userMap[uid] = user
	}
	return
}

//更新用户统计下载量
func (u *UserStat) IncrByUserDownNum(pxUid, num int) (affected int64, err error) {
	buildMap := map[string]interface{}{
		"downloads_num": gorm.Expr("downloads_num + ?", num),
	}
	updates := GetDB().Model(u).Where("px_uid = ?", pxUid).Updates(buildMap).Omit("add_time")
	return updates.RowsAffected, updates.Error
}

//更新用户统计浏览量
func (u *UserStat) IncrByUserViewNum(pxUid, num int) (affected int64, err error) {
	buildMap := map[string]interface{}{
		"view_num": gorm.Expr("view_num + ?", num),
	}
	updates := GetDB().Model(u).Where("px_uid = ?", pxUid).Updates(buildMap).Omit("add_time")
	return updates.RowsAffected, updates.Error
}
