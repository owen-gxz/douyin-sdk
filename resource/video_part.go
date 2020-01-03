package resource

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/owen-gxz/douyin-sdk/util"
	"github.com/rs/xid"
	"io/ioutil"
	"math"
	"mime/multipart"
	"net/http"
	"net/url"
)

//视频分片上传
var (
	videoPartInitUrl     = fmt.Sprintf("%s/video/part/init/", util.ServiceUrl)
	videoPartUploadUrl   = fmt.Sprintf("%s/video/part/upload/", util.ServiceUrl)
	videoPartCompleteUrl = fmt.Sprintf("%s/video/part/complete/", util.ServiceUrl)
	// 建议最小视频为5M
	minPartSize = 5 * 1024 * 1024
)

type VideoPartInitResponse struct {
	Data struct {
		UploadID    string `json:"upload_id"`
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	} `json:"data"`
}

// 初始化
func VideoPartInit(accountToken, openID string) (*VideoPartInitResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(videoPartInitUrl)
	v := url.Values{
		"access_token": {accountToken},
		"open_id":      {openID},
	}
	buf.WriteByte('?')
	buf.WriteString(v.Encode())
	resp := &VideoPartInitResponse{}
	err := util.Post2Response(buf.String(), nil, resp)
	if err != nil {
		return nil, err
	}
	fmt.Println("upid:", resp.Data.UploadID)
	return resp, nil
}

type VideoPartUploadResponse struct {
	Data struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	} `json:"data"`
}

// 分片上传
func VideoPartUpload(accountToken, openID string, uploadID string, fileData []byte) (error) {
	var buf bytes.Buffer
	buf.WriteString(videoUploadUrl)
	v := url.Values{
		"access_token": {accountToken},
		"open_id":      {openID},
		"upload_id":    {uploadID},
	}
	buf.WriteByte('?')
	var partNum = 0
	if len(fileData) < minPartSize {
		partNum = 1
	} else {
		partNum = int(math.Ceil(float64(len(fileData)) / float64(minPartSize)))
	}
	fmt.Println("分块数量：", partNum)
	index := 0
	errorChain := make(chan error)
	done := make(chan int)
	for {
		var fd []byte
		if index+1 == partNum {
			fmt.Println(index)
			fd = fileData[index*minPartSize:]
		} else {
			fmt.Println(index)
			fd = fileData[index*minPartSize : (index+1)*minPartSize]
		}
		index++
		go func(num int, fdata []byte) {
			var b bytes.Buffer
			w := multipart.NewWriter(&b)
			fw, _ := w.CreateFormFile("video", xid.New().String())
			_, err := fw.Write(fdata)
			if err != nil {
				errorChain <- err
				return
			}
			w.Close()
			v.Set("part_number", fmt.Sprintf("%d", num))
			req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", buf.String(), v.Encode()), &b)
			if err != nil {
				errorChain <- err
				return
			}
			req.Header.Add("Content-Type", w.FormDataContentType())
			response, err := http.DefaultClient.Do(req)
			if err != nil {
				errorChain <- err
				return
			}
			data, err := ioutil.ReadAll(response.Body)
			if err != nil {
				errorChain <- err
				return
			}
			vr := &VideoPartUploadResponse{}
			err = json.Unmarshal(data, vr)
			if err != nil {
				errorChain <- err
				return
			}
			if vr.Data.ErrorCode != 0 {
				errorChain <- errors.New(vr.Data.Description)
				return
			}
			done <- 1
		}(index, fd)
		if index == partNum {
			break
		}
	}
	var err error
	var su int
	for {
		select {
		case err = <-errorChain:
			fmt.Println(err.Error())
			return err
		case <-done:
			fmt.Println("完成一块")
			su++
			if su == partNum {
				return nil
			}
		}
	}
}

type VideoPartCompleteResponse struct {
	VideoUploadResponse
}

// 上传完成
func VideoPartComplete(accountToken, openID string, uploadID string) (*VideoPartCompleteResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(videoPartCompleteUrl)
	v := url.Values{
		"access_token": {accountToken},
		"open_id":      {openID},
		"upload_id":    {uploadID},
	}
	buf.WriteByte('?')
	buf.WriteString(v.Encode())
	resp := &VideoPartCompleteResponse{}
	//err := util.Post2Response(fmt.Sprintf("%s&upload_id=%s", buf.String(), uploadID), nil, resp)
	err := util.Post2Response(buf.String(), nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func VideoPart(accountToken, openID string, fileData []byte) (*VideoPartCompleteResponse, error) {
	vp, err := VideoPartInit(accountToken, openID)
	if err != nil {
		return nil, err
	}
	if vp.Data.ErrorCode != 0 {
		return nil, errors.New(vp.Data.Description)
	}
	upID := vp.Data.UploadID
	fmt.Println(upID)
	err = VideoPartUpload(accountToken, openID, upID, fileData)
	if err != nil {
		return nil, err
	}
	vpc, err := VideoPartComplete(accountToken, openID, upID)
	if err != nil {
		return nil, err
	}
	return vpc, nil
}
