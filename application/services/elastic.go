/*
@Time : 2020/7/24 11:32
@Author : zxr
@File : elastic
@Software: GoLand
*/
package services

import (
	"context"
	"fmt"
	"pix/configs"
	"reflect"
	"strings"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

var client *elastic.Client
var host = configs.ES_HOST

//搜索
type Search struct {
	client *elastic.Client
}

func NewElastic() *Search {
	log := logrus.New()
	var err error
	client, err = elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetErrorLog(log),
		elastic.SetURL(host),
		elastic.SetBasicAuth(configs.ESUSER, configs.ESPASSWORD),
	)
	if err != nil {
		logrus.Error("elastic.NewClient err:", err)
		panic(err)
	}
	info, code, err := client.Ping(host).Do(context.Background())
	if err != nil {
		logrus.Error("client.Ping err:", err)
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	if _, err := client.ElasticsearchVersion(host); err != nil {
		panic(err)
	}
	return &Search{
		client: client,
	}
}

//按tag名搜索
func (s *Search) SearchTag(tag string, offset, limit int) (photoResult []*PhotoResult, tagList []string, total int, err error) {
	boolQuery := elastic.NewBoolQuery()
	query := elastic.NewMatchPhraseQuery("tags", tag).Analyzer("ik_smart")
	boolQuery.Must(query)

	//只查指定字段
	fsc := elastic.NewFetchSourceContext(true).Include("pic_id", "tags")

	//test
	//src, _ := boolQuery.Source()
	//data, _ := json.MarshalIndent(src, "", "  ")
	//logrus.Println("source:", string(data))
	//test end

	result, err := s.client.Search().Index(configs.ES_INDEX).Query(boolQuery).
		From(offset).Size(limit).FetchSourceContext(fsc).Pretty(true).Do(context.Background())
	if err != nil {
		return nil, nil, 0, err
	}
	total = int(result.Hits.TotalHits.Value)
	if total > 0 {
		photoResult, tagList = s.getPhotoResult(result, limit)
	}
	return
}

//根据搜索结果获取图片信息
func (s *Search) getPhotoResult(result *elastic.SearchResult, limit int) (photoResult []*PhotoResult, tagList []string) {
	var photo PhotoIndexData
	pxIdList := make([]int, limit)
	tagList = make([]string, 0)
	for k, item := range result.Each(reflect.TypeOf(photo)) {
		if photoData, ok := item.(PhotoIndexData); ok {
			pxIdList[k] = photoData.PicId
			if len(tagList) < 15 {
				tagStr := photoData.Tags
				tagStrList := strings.Split(tagStr, ",")
				tagList = append(tagList, tagStrList...)
			}
		}
	}
	photoResult = NewPicService().GetPicListByIds(pxIdList)
	return
}
