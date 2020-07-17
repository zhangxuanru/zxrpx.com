/*
@Time : 2020/6/11 17:57
@Author : zxr
@File : user
@Software: GoLand
*/
package models

import (
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
	fields := "author_id"
	GetDB().Where("uid = ? AND status =? ", uid, 1).Select(fields).Find(&list)
	return
}
