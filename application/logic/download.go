/*
@Time : 2020/7/7 11:20
@Author : zxr
@File : download
@Software: GoLand
*/
package logic

import (
	"io/ioutil"
	"net"
	"net/http"
	"pix/configs"
	"time"
)

type HttpRequest struct {
}

func DownloadPic(srcUrl string, req *HttpRequest) (body []byte, err error) {
	var (
		request  *http.Request
		response *http.Response
	)
	if request, err = http.NewRequest(http.MethodGet, srcUrl, nil); err != nil {
		return nil, err
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")
	request.Header.Add("Referer", configs.WEBSITE)

	if response, err = NewHttpClient().Do(request); err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	return
}

func NewHttpClient() (client *http.Client) {
	client = &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
	return
}
