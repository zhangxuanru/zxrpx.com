/*
@Time : 2020/7/2 16:45
@Author : zxr
@File : comments
@Software: GoLand
*/
package services

import (
	"pix/application/models"
	"time"
)

type CommentsService struct {
}

type CommentItem struct {
	Uid          int
	PicId        int
	HeadPortrait string
	HeadFileName string
	QiNiu        int
	NickName     string
	Content      string
	AddTime      time.Time
}

type CommentRes []CommentItem

func NewComments() *CommentsService {
	return &CommentsService{}
}

//根据图片ID获取评论列表
func (c *CommentsService) GetCommentsByPicId(picId int, page, limit int) (result CommentRes, count int) {
	offset := (page - 1) * limit
	list, count := models.NewComments().GetListByPicId(picId, offset, limit)
	uidList := c.getUidList(list)
	users := models.NewUser().GetUserByUidList(uidList)
	userMap := make(map[int]models.User)
	for _, user := range users {
		userMap[user.Id] = user
	}
	result = make([]CommentItem, len(list))
	for k, comment := range list {
		uid := comment.Uid
		item := CommentItem{
			Uid:          uid,
			PicId:        comment.PicId,
			HeadPortrait: userMap[uid].HeadPortrait,
			HeadFileName: userMap[uid].FileName,
			QiNiu:        userMap[uid].IsQiniu,
			NickName:     userMap[uid].NickName,
			Content:      comment.Content,
			AddTime:      comment.AddTime,
		}
		result[k] = item
	}
	return result, count
}

//从评论里获取uid
func (c *CommentsService) getUidList(comments []models.Comments) []int {
	uMap := make(map[int]int)
	for index, item := range comments {
		uMap[index] = item.Uid
	}
	list := make([]int, len(uMap))
	for index, uid := range uMap {
		list[index] = uid
		delete(uMap, index)
	}
	return list
}
