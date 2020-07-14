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

const UserPx = 2
const UserStation = 1

//用户表
type User struct {
	Id           int       `gorm:"primary_key; AUTO_INCREMENT; comment:'自增ID'" json:"id"`                               //指定主键并自增
	PxUid        int       `gorm:"unique; not null; comment:'px 用户ID'" json:"px_uid"`                                   //用户ID
	NickName     string    `gorm:"type:varchar(100); NOT NULL; comment:'用户昵称'" json:"nick_name"`                        //昵称
	UserName     string    `gorm:"index:user_pass; type:varchar(50);  comment:'登录用户名'" json:"user_name"`                //登录用户名
	Passwd       string    `gorm:"index:user_pass; type:char(32);     comment:'登录密码'" json:"passwd"`                    //用户密码
	Email        string    `gorm:"type:varchar(100);    comment:'邮箱地址'" json:"email"`                                   //邮箱地址
	UserType     int       `gorm:"type:TINYINT(1); NOT NULL;default:1; comment:'用户类型 1:本站注册 2:px站抓取'" json:"user_type"` //用户类型 1,本站注册， 2：px站抓取
	HeadPortrait string    `gorm:"type:varchar(100); NOT NULL;comment:'头像地址'" json:"head_portrait"`                     //头像地址
	FileName     string    `gorm:"type:varchar(255); NOT NULL; comment:'图片名称'" json:"file_name"`
	IsQiniu      int       `gorm:"type:TINYINT(1); NOT NULL;default:0; comment:'是否上传七牛 1:已上传 0:未上传'" json:"is_qiniu"`
	AddTime      time.Time `gorm:"comment:'添加时间'" json:"add_time"`    //创建时间
	UpdateTime   time.Time `gorm:"comment:'修改时间'" json:"update_time"` //修改后自动更新时间
}

func NewUser() *User {
	return &User{}
}

//根据用户ID列表查询用户信息
func (u *User) GetUserByUidList(uidList []int) (list []User) {
	if len(uidList) == 0 {
		return
	}
	fields := "id,px_uid,nick_name,head_portrait,file_name,is_qiniu"
	GetDB().Where("id IN (?) ", uidList).Select(fields).Find(&list)
	return
}

//根据用户名和密码查询用户信息
func (u *User) GetUserInfoByNameAndPass(userName string, password string) (user User) {
	fields := "id,px_uid,nick_name,head_portrait,user_name,passwd"
	GetDB().Where("user_name=? AND passwd=?", userName, password).Select(fields).Find(&user)
	return
}

//根据用户名查询用户信息
func (u *User) GetUserInfoByName(userName string) (user User) {
	fields := "id"
	GetDB().Where("user_name=?", userName).Select(fields).Find(&user)
	return
}

//根据邮箱查询用户信息
func (u *User) GetUserInfoByEmail(email string) (user User) {
	fields := "id"
	GetDB().Where("email=?", email).Select(fields).Find(&user)
	return
}

//插入用户
func (u *User) Insert(user *User) (id int, err error) {
	db := GetDB().Create(user)
	return user.Id, db.Error
}

//根据id更新pxUid
func (u *User) UpdatePxUidById(id int, pxUid int) (affected int64, err error) {
	buildMap := map[string]interface{}{
		"px_uid": pxUid,
	}
	updates := GetDB().Model(u).Where("id = ?", id).Updates(buildMap).Omit("add_time")
	return updates.RowsAffected, updates.Error
}
