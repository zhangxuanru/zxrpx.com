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

//用户扩展表
type UserExtend struct {
	Id             int       `gorm:"primary_key; AUTO_INCREMENT; comment:'自增ID'" json:"id"`
	Uid            int       `gorm:"unique; not null; comment:'用户ID'" json:"uid"`
	Gender         int       `gorm:"comment:'1:男2:女'" json:"gender"`
	Name           string    `gorm:"comment:'姓名'" json:"name"`
	Email          string    `gorm:"comment:'邮箱地址'" json:"email"`
	BirthDay       string    `gorm:"comment:'出生日期'" json:"birth_day"`
	Intro          string    `gorm:"comment:'个人简介'" json:"intro"`
	Facebook       string    `gorm:"comment:'facebook地址'" json:"facebook"`
	Twitter        string    `gorm:"comment:'twitter'" json:"twitter"`
	Website        string    `gorm:"comment:'website'" json:"website"`
	AddTime        time.Time `gorm:"comment:'添加时间'" json:"add_time"`
	LastUpdateTime time.Time `gorm:"comment:'修改时间'" json:"last_update_time"`
}

func NewUserExtend() *UserExtend {
	return &UserExtend{}
}

//插入扩展信息
func (u *UserExtend) Insert(extend *UserExtend) (id int, err error) {
	db := GetDB().Create(extend)
	return u.Id, db.Error
}

//根据uid查询用户扩展信息
func (u *UserExtend) GetUserExtendByUid(uid int) (user UserExtend) {
	fields := "id,uid,name,intro,facebook,twitter,website,email"
	GetDB().Where("uid=?", uid).Select(fields).Find(&user)
	return
}

//根据邮箱查询用户扩展信息
func (u *UserExtend) GetUserExtendByEmail(email string) (user UserExtend) {
	GetDB().Where("email=?", email).Select("id").Find(&user)
	return
}

//更新扩展信息
func (u *UserExtend) UpdateExtend(uid int, buildMap map[string]interface{}) (affected int64, err error) {
	updates := GetDB().Model(u).Where("uid = ?", uid).Updates(buildMap).Omit("add_time")
	return updates.RowsAffected, updates.Error
}
