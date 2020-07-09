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

//图片标签表
type PictureTag struct {
	Id         int       `gorm:"primary_key; AUTO_INCREMENT; comment:'自增ID'" json:"id"`
	PicId      int       `gorm:"index:pic_id; not null; comment:'图片ID'" json:"pic_id"`
	TagId      string    `gorm:"type:varchar(50); not null; comment:'标签ID 以,分割'" json:"tag_id"`
	State      int       `gorm:"type:TINYINT(1); NOT NULL;default:1; comment:'状态 1:状态正常 0:删除'" json:"state"`
	AddTime    time.Time `gorm:"comment:'添加时间'" json:"add_time"`
	UpdateTime time.Time `gorm:"comment:'修改时间'" json:"update_time"`
}

func NewPictureTag() *PictureTag {
	return &PictureTag{}
}

//根据图片ID查询标签ID
func (p *PictureTag) GetTagByPicIds(picIds []int) (result []PictureTag) {
	fields := "pic_id,tag_id"
	GetDB().Where("pic_id IN (?) AND state=?", picIds, 1).Select(fields).Find(&result)
	return
}
