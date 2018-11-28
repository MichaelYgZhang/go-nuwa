package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/tidwall/gjson"
)

func main() {
	fmt.Print("hello,faceplusplus")
	form := url.Values{}
	form.Add("api_key", "dzJzPS7APmA356SyxJAJWcYIgwtZTBzT")
	form.Add("api_secret", "9RP6WK5E2pLNVJEaeLvLR4eTLz0EC4ic")
	form.Add("return_landmark", "1")
	form.Add("return_attributes", "age,beauty")

	form.Add("image_base64", ReadImg("img/img.jpg"))

	resp, err := http.PostForm("https://api-cn.faceplusplus.com/facepp/v3/detect", form)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", resp.Body)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	faces := gjson.Get(string(body), "faces")
	for _, face := range faces.Array() {
		attributes := face.Get("attributes")
		beauty := attributes.Get("beauty")
		female_score := beauty.Get("female_score")
		male_score := beauty.Get("male_score")
		fmt.Printf("%s, %s", female_score, male_score)
	}
}

func ReadImg(imgUrl string) string {
	img, err := os.Open(imgUrl)
	if err != nil {
		panic(err)
	}
	sourcebuffer := make([]byte, 500000)
	n, _ := img.Read(sourcebuffer)
	//base64压缩
	return base64.StdEncoding.EncodeToString(sourcebuffer[:n])
}
