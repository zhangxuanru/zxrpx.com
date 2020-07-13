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
)

type Account struct {
	UserName string
	PassWord string
	Token    string
	Account  *models.User
}

func NewAccount() *Account {
	return &Account{}
}

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

//关注
func (a *Account) Follow(userId, authorId int) {

}
