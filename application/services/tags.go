/*
@Time : 2020/6/24 14:13
@Author : zxr
@File : tag
@Software: GoLand
*/
package services

import (
	"bytes"
	"math/rand"
	"pix/application/models"
	"strconv"
	"strings"
	"sync"
	"time"
)

type TagService struct {
}

type tagMapRes map[int][]models.Tag

var tag *TagService
var once sync.Once

func NewTagService() *TagService {
	once.Do(func() {
		tag = &TagService{}
	})
	return tag
}

//根据图片ID获取标签信息
func (t *TagService) GetTagListByPicIds(picIds []int) (result tagMapRes) {
	var (
		picTagList []models.PictureTag
		tagData    map[int]string
	)
	result = make(tagMapRes)
	if picTagList = models.NewPictureTag().GetTagByPicIds(picIds); len(picTagList) == 0 {
		return
	}
	tagIdList := t.GetTagIdByPicTag(picTagList)
	if tagData = models.NewTag().GetTagListByIds(tagIdList); len(tagData) == 0 {
		return
	}
	for _, pic := range picTagList {
		idList := strings.Split(pic.TagId, ",")
		for _, idStr := range idList {
			id, _ := strconv.Atoi(idStr)
			if _, ok := tagData[id]; !ok {
				continue
			}
			result[pic.PicId] = append(result[pic.PicId], models.Tag{Id: id, TagName: tagData[id]})
		}
	}
	return
}

//获取TAG列表
func (t *TagService) GetRandOffsetTagList(limit int) (list []models.Tag) {
	rand.Seed(time.Now().UnixNano())
	offset := rand.Intn(900)
	data := models.NewTag().GetTagList(offset, limit)
	for k, tag := range data {
		tag.TagName = strings.TrimSpace(tag.TagName)
		data[k] = tag
	}
	return data
}

//根据picTag 获取tagId列表
func (t *TagService) GetTagIdByPicTag(picTag []models.PictureTag) []int {
	tagIds := make([]int, 0)
	for _, tag := range picTag {
		if tag.TagId == "" {
			continue
		}
		tagArr := strings.Split(tag.TagId, ",")
		for _, idStr := range tagArr {
			id, _ := strconv.Atoi(idStr)
			tagIds = append(tagIds, id)
		}
	}
	return tagIds
}

//根据tag数组获取tag的字符串
func (t *TagService) GetTagStrByTagModels(models []models.Tag) string {
	if len(models) == 0 {
		return ""
	}
	var buf bytes.Buffer
	for _, tag := range models {
		tagName := strings.TrimSpace(tag.TagName)
		buf.WriteString(tagName + ",")
	}
	return buf.String()
}
