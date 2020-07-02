/*
@Time : 2020/6/12 17:32
@Author : zxr
@File : picture
@Software: GoLand
*/
package models

import (
	"time"
)

//图片评论表
type Comments struct {
	Id      int       `gorm:"primary_key; AUTO_INCREMENT; comment:'自增ID'" json:"id"`
	PicId   int       `gorm:"index:pic_id; not null; comment:'图片ID'" json:"pic_id"`
	Content string    `gorm:"type:text; NOT NULL; comment:'评论内容'" json:"content"`
	Uid     int       `gorm:"index:uid;   comment:'用户ID'" json:"uid"`
	State   int       `gorm:"type:TINYINT(1); NOT NULL;default:1; comment:'状态 1:状态正常 0:删除'" json:"state"`
	Sort    int       `gorm:"type:TINYINT(1); NOT NULL;default:1; comment:'显示排序，从小到大'" json:"sort"`
	AddTime time.Time `gorm:"comment:'添加时间'" json:"add_time"`
}

func NewComments() *Comments {
	return &Comments{}
}

//根据PICID获取评论列表
func (c *Comments) GetListByPicId(picId int, offset, limit int) (list []Comments, count int) {
	field := "pic_id,content,uid,add_time"
	GetDB().Select(field).Where("pic_id =? AND state =?", picId, 1).Order("sort ASC,add_time DESC").Offset(offset).Limit(limit).Find(&list)
	GetDB().Model(c).Where("pic_id =? AND state =?", picId, 1).Count(&count)
	return
}

func (c *Comments) Insert() (id int, err error) {
	create := GetDB().Create(c)
	return c.Id, create.Error
}
