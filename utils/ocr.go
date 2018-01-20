package utils

import (
	"net/url"
	"fmt"
	"net/http"
	"io/ioutil"
	"zcm_tools/file"
	"encoding/json"
)

const (
	ocrUrl = "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic?access_token=24.1effd226601df614a0df44cd82ac74e4.2592000.1518351251.282335-10682710"
)

type OCRResponse struct {
	Words_result []BaseStruct
}
type BaseStruct struct {
	Words string
}

func OCR(path string) string {
	resss := &OCRResponse{}
	result := ""
	ss, err := file.GetFileToBase64(path)
	data := make(url.Values)
	data["image"] = []string{ss}
	//把post表单发送给目标服务器
	res, err := http.PostForm(ocrUrl, data)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &resss)
	for _, k := range resss.Words_result {
		result += k.Words
	}
	defer res.Body.Close()
	return result
}
