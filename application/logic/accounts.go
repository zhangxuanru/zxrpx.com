/*
@Time : 2020/7/13 14:28
@Author : zxr
@File : accounts
@Software: GoLand
*/
package logic

//生成登录表单token
func GenLoginFormToken() string {
	token := Md5(GetRandomString(6))
	return token
}
