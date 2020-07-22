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

//用户关注表
type UserFollow struct {
	Id         int       `gorm:"primary_key; AUTO_INCREMENT; comment:'自增ID'" json:"id"`
	Uid        int       `gorm:"comment:'用户ID'" json:"uid"`
	AuthorId   int       `gorm:"comment:'对方UID'" json:"author_id"`
	Status     int       `gorm:"comment:'状态,1:正常，0：删除'" json:"status"`
	AddTime    time.Time `gorm:"comment:'添加时间'" json:"add_time"`
	UpdateTime time.Time `gorm:"comment:'修改时间'" json:"update_time"`
}

func NewUserFollow() *UserFollow {
	return &UserFollow{}
}

//根据用户ID 查询关注列表
func (u *UserFollow) GetFollowListByUid(uid int) (list []UserFollow) {
	GetDB().Where("uid = ? AND status =? ", uid, 1).Select("author_id").Find(&list)
	return
}

//添加关注
func (u *UserFollow) AddFollow(userId, authorId int) (id int, err error) {
	var user UserFollow
	GetDB().Where("uid = ? AND author_id =? ", userId, authorId).Select("id").Find(&user)
	if user.Id > 0 {
		return 0, errors.New("已关注")
	}
	user = UserFollow{
		Uid:        userId,
		AuthorId:   authorId,
		Status:     1,
		AddTime:    time.Now(),
		UpdateTime: time.Now(),
	}
	db := GetDB().Create(&user)
	return user.Id, db.Error
}

//查询用户ID关注作者ID的信息
func (u *UserFollow) GetFollowAuthIdByUid(userId, authorId int) (follow UserFollow) {
	GetDB().Where("uid = ? AND status =? AND author_id =?", userId, 1, authorId).Select("id").Find(&follow)
	return
}
