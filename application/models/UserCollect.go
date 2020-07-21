/*
@Time : 2020/6/11 17:57
@Author : zxr
@File : user
@Software: GoLand
*/
package models

import (
	"errors"
	"time"
)

//用户收藏表
type UserCollect struct {
	Id         int       `gorm:"primary_key; AUTO_INCREMENT; comment:'自增ID'" json:"id"`
	Uid        int       `gorm:"comment:'用户ID'" json:"uid"`
	PxImgId    int       `gorm:"comment:'图片ID'" json:"px_img_id"`
	Status     int       `gorm:"comment:'状态,1:正常，0：删除'" json:"status"`
	AddTime    time.Time `gorm:"comment:'添加时间'" json:"add_time"`
	UpdateTime time.Time `gorm:"comment:'修改时间'" json:"update_time"`
}

func NewUserCollect() *UserCollect {
	return &UserCollect{}
}

//根据用户ID 查询收藏列表
func (u *UserCollect) GetCollectListByUid(uid int) (list []UserCollect) {
	GetDB().Where("uid = ? AND status =? ", uid, 1).Select("px_img_id").Find(&list)
	return
}

//添加关注
func (u *UserCollect) AddCollect(userId, PxImgId int) (id int, err error) {
	var user UserCollect
	GetDB().Where("uid = ? AND px_img_id =? ", userId, PxImgId).Select("id").Find(&user)
	if user.Id > 0 {
		return 0, errors.New("已收藏")
	}
	user = UserCollect{
		Uid:        userId,
		PxImgId:    PxImgId,
		Status:     1,
		AddTime:    time.Now(),
		UpdateTime: time.Now(),
	}
	db := GetDB().Create(&user)
	return user.Id, db.Error
}
