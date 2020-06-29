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

type UserMap map[int]models.User

func NewUserService() *UserService {
	once.Do(func() {
		user = &UserService{}
	})
	return user
}

//根据UId获取用户信息
func (u *UserService) GetUserDataByUidList(uidList []int) (userData UserMap) {
	userList := models.NewUser().GetUserByUidList(uidList)
	if len(userList) == 0 {
		userData = make(UserMap)
		return
	}
	userData = make(UserMap, len(userList))
	for _, user := range userList {
		userData[user.Id] = user
	}
	return
}
