/*
@Time : 2020/6/29 18:23
@Author : zxr
@File : user
@Software: GoLand
*/
package services

import (
	"errors"
	"fmt"
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

//根据uid获取近期的num张图像,排队pxId的图片
func (u *UserService) GetRecentImagesByUid(uid int, num int, pxId int) (result []*PhotoResult) {
	var picList []models.Picture
	params := &models.QueryParams{
		Offset:  0,
		Limit:   num,
		IsCount: false,
		Fields:  Fields,
		Order:   "add_time DESC",
		Where:   fmt.Sprintf("px_uid=%d AND state=1 AND px_img_id!=%d", uid, pxId),
	}
	if picList, _ = models.NewPicture().GetPicListByParams(params); len(picList) == 0 {
		return
	}
	return NewPicService().combinePicAttrTagData(picList)
}

//更新用户统计下载量
func (u *UserService) IncrByUserDownNum(pxUid, num int) (affected int64, err error) {
	if pxUid == 0 {
		return 0, errors.New("uid is zero")
	}
	affected, err = models.NewUserStat().IncrByUserDownNum(pxUid, num)
	return
}

//根据uid查询用户的扩展信息
func (u *UserService) GetUserExtend(uid int) (extend models.UserExtend) {
	if uid == 0 {
		extend = models.UserExtend{}
	} else {
		extend = models.NewUserExtend().GetUserExtendByUid(uid)
	}
	return
}

//根据UID查询用户信息
func (u *UserService) GetUserInfoByUid(uid int) (user models.User) {
	list := models.NewUser().GetUserByUidList([]int{uid})
	if len(list) == 0 {
		return
	}
	return list[0]
}

//修改用户资料
func (u *UserService) UpdateUserInfo(uid int, buildMap map[string]interface{}) (err error) {
	_, err = models.NewUser().UpdateUserInfo(uid, buildMap)
	return
}

//修改用户扩展资料
func (u *UserService) UpdateUserExtend(uid int, buildMap map[string]interface{}) (err error) {
	_, err = models.NewUserExtend().UpdateExtend(uid, buildMap)
	return
}
