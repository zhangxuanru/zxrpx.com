/*
@Time : 2020/6/29 18:23
@Author : zxr
@File : user
@Software: GoLand
*/
package services

import (
	"pix/application/models"
)

type UserService struct {
}

var user *UserService

type UserCenter map[int]*UserStat

type UserStat struct {
	Uid          int    `json:"uid"`
	PxUid        int    `json:"px_uid"`        //用户ID
	NickName     string `json:"nick_name"`     //昵称
	HeadPortrait string `json:"head_portrait"` //头像地址
	FileName     string `json:"file_name"`
	IsQiniu      int    `json:"is_qiniu"`
	PicNum       int    ` json:"pic_num"`
	ViewNum      int    `json:"view_num"`
	DownloadsNum int    `json:"downloads_num"`
	LikeNum      int    `json:"like_num"`
	CommentNum   int    `json:"comment_num"`
	FollowerNum  int    `json:"follower_num"`
}

func NewUserService() *UserService {
	once.Do(func() {
		user = &UserService{}
	})
	return user
}

//根据UId获取用户信息
func (u *UserService) GetUserDataByUidList(uidList []int) (userMap UserCenter) {
	userInfoList := models.NewUser().GetUserByUidList(uidList)
	if len(userInfoList) == 0 {
		userMap = make(UserCenter)
		return
	}
	userStatMap := models.NewUserStat().GetUserStatByUidList(uidList)
	userMap = make(UserCenter, len(userInfoList))
	for _, user := range userInfoList {
		stat := userStatMap[user.Id]
		userInfo := &UserStat{
			Uid:          user.Id,
			PxUid:        user.PxUid,
			NickName:     user.NickName,
			HeadPortrait: user.HeadPortrait,
			FileName:     user.FileName,
			IsQiniu:      user.IsQiniu,
			PicNum:       stat.PicNum,
			ViewNum:      stat.ViewNum,
			DownloadsNum: stat.DownloadsNum,
			LikeNum:      stat.LikeNum,
			CommentNum:   stat.CommentNum,
			FollowerNum:  stat.FollowerNum,
		}
		userMap[user.Id] = userInfo
	}
	return
}
