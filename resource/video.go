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
	videoUploadUrl = fmt.Sprintf("%s/video/upload/", util.ServiceUrl)
	videoCreateUrl = fmt.Sprintf("%s/video/create/", util.ServiceUrl)
	videoListUrl   = fmt.Sprintf("%s/video/list/", util.ServiceUrl)
	videoInfoUrl   = fmt.Sprintf("%s/video/data/", util.ServiceUrl)
	videoDeleteUrl = fmt.Sprintf("%s/video/delete/", util.ServiceUrl)
)

type VideoUploadResponse struct {
	Data struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
		Video       struct {
			VideoID string `json:"video_id"`
			Width   int    `json:"width"`
			Height  int    `json:"height"`
		} `json:"video"`
	} `json:"data"`
}

func VideoUpload(accountToken, openID string, fileData []byte) (*VideoUploadResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(videoUploadUrl)
	v := url.Values{
		"access_token": {accountToken},
		"open_id":      {openID},
	}
	if strings.Contains(videoUploadUrl, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	err := w.WriteField("Content-Type", http.DetectContentType(fileData[:512]))
	if err != nil {
		return nil, err
	}
	fw, _ := w.CreateFormFile("video", xid.New().String())
	_, err = fw.Write(fileData)
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
	vr := &VideoUploadResponse{}
	err = json.Unmarshal(data, vr)
	if err != nil {
		return nil, err
	}
	return vr, nil
}

type VideoCreateReq struct {
	//video_id, 通过/video/upload/接口得到。注意每次调用/video/create/都要调用/video/upload/生成新的video_id。
	VideoID string `json:"video_id,omitempty"`
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

func VideoCreate(accountToken, openID string, req VideoCreateReq) (*ResourceCreateResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(videoCreateUrl)
	v := url.Values{
		"access_token": {accountToken},
		"open_id":      {openID},
	}
	if strings.Contains(videoCreateUrl, "?") {
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
func Video2DouYin(accountToken, openID string, fileData []byte, title string, ats []string) (*ResourceCreateResponse, error) {
	vu, err := VideoUpload(accountToken, openID, fileData)
	if err != nil {
		return nil, err
	}
	if vu.Data.ErrorCode != 0 {
		return nil, errors.New(vu.Data.Description)
	}
	vcReq := VideoCreateReq{
		VideoID: vu.Data.Video.VideoID,
		Text:    title,
	}
	if ats != nil {
		vcReq.AtUsers = ats
	}
	vcResp, err := VideoCreate(accountToken, openID, vcReq)
	if err != nil {
		return nil, err
	}
	return vcResp, nil
}

type VideoListResponse struct {
	Data struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
		Cursor      int    `json:"cursor"`
		HasMore     bool   `json:"has_more"`
		List        []struct {
			ItemID     string `json:"item_id"`
			Title      string `json:"title"`
			Cover      string `json:"cover"`
			IsTop      bool   `json:"is_top"`
			CreateTime int    `json:"create_time"`
			IsReviewed bool   `json:"is_reviewed"`
			ShareURL   string `json:"share_url"`
			Statistics struct {
				CommentCount  int `json:"comment_count"`
				DiggCount     int `json:"digg_count"`
				DownloadCount int `json:"download_count"`
				PlayCount     int `json:"play_count"`
				ShareCount    int `json:"share_count"`
				ForwardCount  int `json:"forward_count"`
			} `json:"statistics"`
		} `json:"list"`
	} `json:"data"`
}

// cursor 分页游标, 第一页请求cursor是0, response中会返回下一页请求用到的cursor, 同时response还会返回has_more来表明是否有更多的数据。
// count 每页数量
func GetVideoList(accountToken, openID string, cursor, count int) (*VideoListResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(videoListUrl)
	v := url.Values{
		"access_token": {accountToken},
		"open_id":      {openID},
		"cursor":       {fmt.Sprintf("%d", cursor)},
		"count":        {fmt.Sprintf("%d", count)},
	}
	if strings.Contains(videoListUrl, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	resp := &VideoListResponse{}
	err := util.Get2Response(buf.String(), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type VideoInfoResponse struct {
	VideoListResponse
}
type VideoItemReq struct {
	ItemID []string `json:"item_ids"`
}

//
func GetVideosInfo(accountToken, openID string, itemIDs []string) (*VideoInfoResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(videoInfoUrl)
	v := url.Values{
		"access_token": {accountToken},
		"open_id":      {openID},
	}
	buf.WriteByte('?')
	buf.WriteString(v.Encode())
	req := VideoItemReq{itemIDs}
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp := &VideoInfoResponse{}
	err = util.Post2Response(buf.String(), bytes.NewReader(data), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type RemoveVideoResponse struct {
	Data struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	} `json:"data"`
}

//删除多视频
func RemoveVideo(accountToken, openID string, itemID string) (*RemoveVideoResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(videoDeleteUrl)
	v := url.Values{
		"access_token": {accountToken},
		"open_id":      {openID},
	}
	buf.WriteByte('?')
	buf.WriteString(v.Encode())
	req := map[string]interface{}{"item_id": itemID}
	data, err := json.Marshal(&req)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(data))
	resp := &RemoveVideoResponse{}
	err = util.Post2Response(buf.String(), bytes.NewReader(data), resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
