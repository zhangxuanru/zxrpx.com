/*
@Time : 2020/6/23 18:00
@Author : zxr
@File : picture
@Software: GoLand
*/
package services

import (
	"fmt"
	"pix/application/models"
)

type PicService struct {
	Page  int
	Limit int
}

func NewPicService(page int) *PicService {
	if page < 1 {
		page = 1
	}
	return &PicService{
		Page: page,
	}
}

//根据添加时间查询图片列表
func (p *PicService) GetPicListByAddTimeOrder() {
	offset := (p.Page - 1) * p.Limit
	params := &models.QueryParams{
		Offset: offset,
		Limit:  p.Limit,
		Fields: "id,px_img_id,like_num,favorites_num,comments_num,image_type",
		Order:  "add_time DESC,like_num DESC",
		Where:  "state=1",
	}
	picList := models.NewPicture().GetPicListByParams(params)
	if len(picList) == 0 {
		return
	}
	idList := make([]int, len(picList))
	for key, pic := range picList {
		idList[key] = pic.PxImgId
	}
	fmt.Println()
	fmt.Println("idList:", idList)

	list := models.NewPictureAttr().GetListByPicIds(idList)

	fmt.Println()
	fmt.Println("list:", list)

	fmt.Println()
	fmt.Printf("%+v:", picList)
}
