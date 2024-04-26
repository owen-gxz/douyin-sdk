package goods

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
	"testing"
)

func TestCreate(t *testing.T) {
	req := AttrKeyValueMap{
		AutoRenew:   "true",
		MarketPrice: "1000",
		Appointment: &AppointmentStruct{
			NeedAppointment: false,
			ExternalLink:    "12323",
		},
		DescriptionRichText: []*NoteStruct{
			{
				Content: "111",
			},
		},
	}
	vm := make(map[string]string)
	ttt(req, "12323", vm)
	fmt.Println(vm)
}

func ttt(req AttrKeyValueMap, accountID string, vm map[string]string) error {
	t1 := reflect.TypeOf(req)
	v2 := reflect.ValueOf(req)
	var err error
	for i := 0; i < t1.NumField(); i++ {
		k := t1.Field(i).Tag.Get("json")
		switch v2.Field(i).Kind() {
		case reflect.Ptr:
			if !v2.Field(i).IsZero() {
				d, _ := json.Marshal(v2.Field(i).Interface())
				vm[k] = string(d)
			}
		case reflect.Slice:
			if !v2.Field(i).IsZero() {
				switch k {
				case CreateImageIL:
					for _, item := range req.ImageList {
						item.Url, err = WriteFile(accountID, item.Url)
						if err != nil {
							return err
						}
					}
				case CreateImageEIL:
					for _, item := range req.EnvironmentImageList {
						item.Url, err = WriteFile(accountID, item.Url)
						if err != nil {
							return err
						}
					}
				case CreateImageDIL:
					for _, item := range req.DishesImageList {
						item.Url, err = WriteFile(accountID, item.Url)
						if err != nil {
							return err
						}
					}
				case CreateImageDEIL:
					for _, item := range req.DetailImageList {
						item.Url, err = WriteFile(accountID, item.Url)
						if err != nil {
							return err
						}
					}
				}
				d, _ := json.Marshal(v2.Field(i).Interface())
				vm[k] = string(d)
			}
		case reflect.String:
			if !v2.Field(i).IsZero() {
				vm[k] = fmt.Sprintf("%v", v2.Field(i).Interface())
			} else {
				v2.Field(i).SetString("123")
				//v2.Field(i).Elem().SetString("123")
			}
		}
	}
	return nil
}

const (
	path = "./public/"
	uri  = "https://www.baidu.com/static/"
)

// 写入文件,保存
func WriteFile(accountID, bic string) (string, error) {
	b, _ := regexp.MatchString(`^data:\s*image\/(\w+);base64,`, bic)
	if !b {
		return "", errors.New("异常")
	}
	re, err := regexp.Compile(`^data:\s*image\/(\w+);base64,`)
	if err != nil {
		return "", err
	}
	allData := re.FindAllSubmatch([]byte(bic), 2)
	fileType := string(allData[0][1]) //png ，jpeg 后缀获取
	base64Str := re.ReplaceAllString(bic, "")

	var filePath = fmt.Sprintf("%s/%s.%s", accountID, "123", fileType)
	ib, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(fmt.Sprintf("%s%s", path, filePath), ib, 0666)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", uri, filePath), nil
}
