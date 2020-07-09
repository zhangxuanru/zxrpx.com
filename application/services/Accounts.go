/*
@Time : 2020/7/9 17:55
@Author : zxr
@File : accounts
@Software: GoLand
*/
package services

type account struct {
}

func NewAccount() *account {
	return &account{}
}

//关注
func (a *account) Follow(userId, authorId int) {

}
