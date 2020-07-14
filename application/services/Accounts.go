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
	if user := userModel.GetUserInfoByEmail(a.Email); user.Id > 0 {
		return a, errors.New("邮箱已存在")
	}
	user := &models.User{
		UserName: a.UserName,
		Passwd:   a.PassWord,
		Email:    a.Email,
		UserType: 1,
		AddTime:  time.Now(),
	}
	if id, err := userModel.Insert(user); err != nil {
		return a, err
	} else {
		_, _ = user.UpdatePxUidById(id, id)
	}
	a.Token, err = NewJWT().GenUserToken(*user)
	a.Account = user
	return
}

//关注
func (a *Account) Follow(userId, authorId int) {

}
