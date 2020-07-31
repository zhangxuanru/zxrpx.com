/*
@Time : 2020/7/31 15:57
@Author : zxr
@File : search
@Software: GoLand
*/
package logic

import (
	"net/url"
	"pix/application/services"
)

//构建URL查询参数
func BuildSearchUrl(search *services.SearchPhoto) string {
	values := url.Values{}
	if len(search.KeyWord) > 0 {
		values.Set("key", search.KeyWord)
	}
	if len(search.Cat) > 0 {
		values.Set("cat", search.Cat)
	}
	if len(search.PhotoType) > 0 {
		values.Set("type", search.PhotoType)
	}
	if len(search.Orientation) > 0 {
		values.Set("orientation", search.Orientation)
	}
	if len(search.Colors) > 0 {
		values.Set("colors", search.Colors)
	}
	if len(search.Order) > 0 {
		values.Set("order", search.Order)
	}
	genUrl := "/photo/search/?" + values.Encode()
	return genUrl
}
