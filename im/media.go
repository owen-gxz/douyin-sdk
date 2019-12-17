package im

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/owen-gxz/douyin-sdk/util"
	"github.com/rs/xid"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
)

var (
	mediaUploadUrl = fmt.Sprintf("%s/media/upload/", util.ServiceUrl)
)

type MediaUploadResponse struct {
	Data struct {
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
		Media       struct {
			MediaID string `json:"media_id"`
		} `json:"media"`
	} `json:"data"`
}

func MediaUpload(accountToken, openID string, fileData []byte) (*MediaUploadResponse, error) {
	var buf bytes.Buffer
	buf.WriteString(mediaUploadUrl)
	v := url.Values{
		"access_token": {accountToken},
		"open_id":      {openID},
	}
	buf.WriteByte('?')
	buf.WriteString(v.Encode())
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	err := w.WriteField("Content-Type", http.DetectContentType(fileData[:512]))
	if err != nil {
		return nil, err
	}
	fw, _ := w.CreateFormFile("media", xid.New().String())
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
	vr := &MediaUploadResponse{}
	err = json.Unmarshal(data, vr)
	if err != nil {
		return nil, err
	}
	return vr, nil
}
