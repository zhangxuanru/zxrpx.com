/*
@Time : 2020/7/13 12:05
@Author : zxr
@File : base
@Software: GoLand
*/
package controllers

type response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	SuccessCode = 200
	ErrCode     = -1
)

func Response(code int, data interface{}, msg ...string) *response {
	r := &response{
		Code: code,
		Data: data,
	}
	if len(msg) > 0 {
		r.Msg = msg[0]
	}
	return r
}

func ResponseSucc(data interface{}, msg string) *response {
	return Response(SuccessCode, data, msg)
}

func ResponseErr(data interface{}, msg string) *response {
	return Response(ErrCode, data, msg)
}
