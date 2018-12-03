package util

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
)

func GetImgFromURL(url string) string {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	imgBytes, _ := ioutil.ReadAll(resp.Body)
	//base64压缩
	return base64.StdEncoding.EncodeToString(imgBytes)
}
