/*
@Time : 2020/7/9 17:55
@Author : zxr
@File : accounts
@Software: GoLand
*/
package services

import (
	"errors"
	"pix/application/models"
	"strings"
	"time"
)

type Account struct {
	UserName string
	PassWord string
	Email    string
	Token    string
	Account  *models.User
}

func NewAccount() *Account {
	return &Account{}
}

//登录操作
func (a *Account) Login() (account *Account, err error) {
	if a.UserName == "" || a.PassWord == "" {
		return a, errors.New("the username or password is empty")
	}
	user := models.NewUser().GetUserInfoByNameAndPass(a.UserName, a.PassWord)
	if user.Id == 0 {
		return a, errors.New("the user does not exist")
	}
	a.Account = &user
	a.Token, err = NewJWT().GenUserToken(user)
	return a, nil
}

//注册操作
func (a *Account) Register() (account *Account, err error) {
	var userModel *models.User
	if a.UserName == "" || a.PassWord == "" || a.Email == "" {
		return a, errors.New("the username or password or email is empty")
	}
	userModel = models.NewUser()
	if user := userModel.GetUserInfoByName(a.UserName); user.Id > 0 {
		return a, errors.New("用户名已存在")
	}
	if user := models.NewUserExtend().GetUserExtendByEmail(a.Email); user.Id > 0 {
		return a, errors.New("邮箱已存在")
	}
	user := &models.User{
		UserName: a.UserName,
		Passwd:   a.PassWord,
		UserType: 1,
		AddTime:  time.Now(),
	}
	if id, err := userModel.Insert(user); err != nil {
		return a, err
	} else {
		_, _ = user.UpdatePxUidById(id, id)
		//插入扩展表
		extend := &models.UserExtend{
			Uid:     id,
			Email:   a.Email,
			AddTime: time.Now(),
		}
		_, _ = models.NewUserExtend().Insert(extend)
	}
	a.Token, err = NewJWT().GenUserToken(*user)
	a.Account = user
	return a, err
}

//修改个人资料
func (a *Account) SettingUser(setting *SettingUser, uid int) (err error) {
	if user := models.NewUserExtend().GetUserExtendByEmail(setting.Email); user.Uid > 0 && user.Uid != uid {
		return errors.New("邮箱已存在")
	}
	var userService *UserService
	userService = NewUserService()
	userInfo := userService.GetUserInfoByUid(uid)
	build := make(map[string]interface{})

	setting.UserName = strings.TrimSpace(setting.UserName)
	setting.NickName = strings.TrimSpace(setting.NickName)

	if userInfo.UserName != setting.UserName {
		build["user_name"] = setting.UserName
	}
	if userInfo.NickName != setting.NickName {
		build["nick_name"] = setting.NickName
	}
	if len(build) > 0 {
		if err = userService.UpdateUserInfo(uid, build); err != nil {
			return
		}
	}
	build = map[string]interface{}{
		"name":             setting.FirstName,
		"intro":            setting.Intro,
		"facebook":         setting.Facebook,
		"twitter":          setting.Twitter,
		"website":          setting.Website,
		"email":            setting.Email,
		"last_update_time": time.Now(),
	}
	if err = userService.UpdateUserExtend(uid, build); err != nil {
		return
	}
	return nil
}

//关注
func (a *Account) Follow(userId, authorId int) {

}
