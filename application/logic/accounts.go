/*
@Time : 2020/7/13 14:28
@Author : zxr
@File : accounts
@Software: GoLand
*/
package logic

func GenLoginToken() string {
	token := Md5(GetRandomString(6))
	return token
}
