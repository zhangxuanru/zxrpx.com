/*
@Time : 2020/6/11 17:57
@Author : zxr
@File : user
@Software: GoLand
*/
package models

import (
	"strings"
	"time"
)

//标签表
type Tag struct {
	Id         int       `gorm:"primary_key; AUTO_INCREMENT; comment:'自增ID'" json:"id"`
	TagName    string    `gorm:"type:varchar(50); unique; NOT NULL; comment:'标签名'" json:"tag_name"`
	State      int       `gorm:"type:TINYINT(1); NOT NULL;default:1; comment:'状态 1:状态正常 0:删除'" json:"state"`
	AddTime    time.Time `gorm:"comment:'添加时间'" json:"add_time"`
	UpdateTime time.Time `gorm:"comment:'修改时间'" json:"update_time"`
}

func NewTag() *Tag {
	return &Tag{}
}

//根据标签ID查询标签信息
func (t *Tag) GetTagListByIds(tagIds []int) (result map[int]string) {
	var list []Tag
	fields := "id,tag_name"
	GetDB().Where("id IN (?) AND state=?", tagIds, 1).Select(fields).Find(&list)
	result = make(map[int]string, len(list))
	for _, tag := range list {
		result[tag.Id] = strings.TrimSpace(tag.TagName)
	}
	return result
}

//获取tag列表
func (t *Tag) GetTagList(offset, limit int) (list []Tag) {
	fields := "id,tag_name"
	GetDB().Where("state=?", 1).Offset(offset).Limit(limit).Select(fields).Find(&list)
	return
}
