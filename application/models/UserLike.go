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

//用户喜欢表
type UserLike struct {
	Id         int       `gorm:"primary_key; AUTO_INCREMENT; comment:'自增ID'" json:"id"`
	Uid        int       `gorm:"comment:'用户ID'" json:"uid"`
	ImgId      int       `gorm:"comment:'图片ID'" json:"img_id"`
	Status     int       `gorm:"comment:'状态,1:正常，0：删除'" json:"status"`
	AddTime    time.Time `gorm:"comment:'添加时间'" json:"add_time"`
	UpdateTime time.Time `gorm:"comment:'修改时间'" json:"update_time"`
}

func NewUserLike() *UserLike {
	return &UserLike{}
}

//根据用户ID 查询喜欢数据
func (u *UserLike) GetLikeByUid(uid, imgId int) (user UserLike) {
	GetDB().Where("uid = ? AND status =? AND img_id =?", uid, 1, imgId).Select("id").Find(&user)
	return
}

//添加喜欢
func (u *UserLike) AddLike(userId, ImgId int) (id int, err error) {
	var user UserLike
	GetDB().Where("uid = ? AND img_id =?", userId, ImgId).Select("id").Find(&user)
	if user.Id > 0 {
		return 0, errors.New("已喜欢成功")
	}
	user = UserLike{
		Uid:        userId,
		ImgId:      ImgId,
		Status:     1,
		AddTime:    time.Now(),
		UpdateTime: time.Now(),
	}
	db := GetDB().Create(&user)
	return user.Id, db.Error
}
