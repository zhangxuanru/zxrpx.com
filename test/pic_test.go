/*
@Time : 2020/6/24 11:57
@Author : zxr
@File : map_test
@Software: GoLand
*/
package test

import (
	"fmt"
	"pix/application/services"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func TestPicService(t *testing.T) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	result, count := services.NewPicService(1, 10).GetPicListByAddTimeOrder()
	fmt.Println(count)
	bytes, _ := json.Marshal(result)
	fmt.Println(string(bytes))
}
