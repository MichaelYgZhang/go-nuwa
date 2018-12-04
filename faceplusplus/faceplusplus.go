package faceplusplus

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/tidwall/gjson"
)

//ReadImg return img base64
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

type FaceInfo struct {
	ImgURL      string
	ImgBase64   string
	FemaleScore float64
	MaleScore   float64
}

func BuildRequestFaceParam() url.Values {
	form := url.Values{}
	form.Add("api_key", "dzJzPS7APmA356SyxJAJWcYIgwtZTBzT")
	form.Add("api_secret", "9RP6WK5E2pLNVJEaeLvLR4eTLz0EC4ic")
	form.Add("return_landmark", "1")
	form.Add("return_attributes", "age,beauty")
	return form
}

const FacePlusPlusURL = "https://api-cn.faceplusplus.com/facepp/v3/detect"

func GetFaceScore(faceInfo *FaceInfo) {
	form := BuildRequestFaceParam()
	// form.Add("image_base64", ReadImg("img/img.jpg"))
	// form.Add("image_base64", ReadImg(faceInfo.ImgURL))
	form.Add("image_base64", faceInfo.ImgBase64)
	resp, err := http.PostForm(FacePlusPlusURL, form)
	if err != nil {
		panic(err)
	}
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
		femaleScore := beauty.Get("female_score")
		maleScore := beauty.Get("male_score")
		faceInfo.FemaleScore = femaleScore.Float()
		faceInfo.MaleScore = maleScore.Float()
	}
}
