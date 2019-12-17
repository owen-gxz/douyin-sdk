package resource

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/owen-gxz/douyin-sdk/util"
	"github.com/rs/xid"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)


var (
	imageUploadUrl = fmt.Sprintf("%s/image/upload/", util.ServiceUrl)
	imageCreateUrl = fmt.Sprintf("%s/image/create/", util.ServiceUrl)
)

type imageUploadResponse struct {
	Data struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
		Image       struct {
			ImageID string `json:"image_id"`
			Width   int    `json:"width"`
			Height  int    `json:"height"`
		} `json:"image"`
	} `json:"data"`
}

func ImageUpload(accountToken, openID string, fileData []byte) (*imageUploadResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(imageUploadUrl)
	v := url.Values{
		"access_token": {accountToken},
		"open_id":      {openID},
	}
	if strings.Contains(imageUploadUrl, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	fw, _ := w.CreateFormFile("image", xid.New().String())
	_, err := fw.Write(fileData)
	if err != nil {
		return nil, err
	}
	w.Close()
	req, err := http.NewRequest(http.MethodPost, buf.String(), &b)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", w.FormDataContentType())
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	ir := &imageUploadResponse{}
	err = json.Unmarshal(data, ir)
	if err != nil {
		return nil, err
	}
	return ir, nil
}

type ImageCreateReq struct {
	//video_id, 通过/video/upload/接口得到。注意每次调用/video/create/都要调用/video/upload/生成新的video_id。
	ImageID string `json:"image_id,omitempty"`
	//视频标题， 可以带话题,@用户。 如title1#话题1 #话题2 @openid1
	Text string `json:"text,omitempty"`
	//地理位置id (未开放)
	PoiID string `json:"poi_id,omitempty"`
	//地理位置名称 (未开放)
	PoiName string `json:"poi_name,omitempty"`
	//小程序id
	MicroAppID string `json:"micro_app_id,omitempty"`
	//小程序标题
	MicroAppTitle string `json:"micro_app_title,omitempty"`
	//吊起小程序时的参数
	MicroAppURL string `json:"micro_app_url,omitempty"`
	//将传入的指定时间点对应帧设置为视频封面（单位：秒）
	CoverTsp float64 `json:"cover_tsp,omitempty"`
	//如果需要at其他用户。将text中@nickname对应的open_id放到这里。
	AtUsers []string `json:"at_users,omitempty"`
}

type ResourceCreateResponse struct {
	Data struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
		ItemID      string `json:"item_id"`
	} `json:"data"`
}

func ImageCreate(accountToken, openID string, req ImageCreateReq) (*ResourceCreateResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(imageCreateUrl)
	v := url.Values{
		"access_token": {accountToken},
		"open_id":      {openID},
	}
	if strings.Contains(imageCreateUrl, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	resp := &ResourceCreateResponse{}
	data, err := json.Marshal(&req)
	if err != nil {
		return nil, err
	}
	err = util.Post2Response(buf.String(), bytes.NewReader(data), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// 该方法直接创建视频并发布到抖音
func Image2DouYin(accountToken, openID string, fileData []byte, title string, ats []string) (*ResourceCreateResponse, error) {
	vu, err := ImageUpload(accountToken, openID, fileData)
	if err != nil {
		return nil, err
	}
	if vu.Data.ErrorCode != 0 {
		return nil, errors.New(vu.Data.Description)
	}
	icReq := ImageCreateReq{
		ImageID: vu.Data.Image.ImageID,
		Text:    title,
	}
	if ats != nil {
		icReq.AtUsers = ats
	}
	vcResp, err := ImageCreate(accountToken, openID, icReq)
	if err != nil {
		return nil, err
	}
	return vcResp, nil
}
