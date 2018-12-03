package util

import (
	"fmt"
	"spider/faceplusplus"
	"testing"
)

func TestGetImgFromURL(t *testing.T) {
	faceInfo := faceplusplus.FaceInfo{
		ImgURL:      "",
		ImgBase64:   GetImgFromURL("https://photo.zastatic.com/images/photo/460892/1843566077/8308752193150958.png"),
		FemaleScore: 0,
		MaleScore:   0,
	}
	faceplusplus.GetFaceScore(&faceInfo)
	fmt.Printf("%v", faceInfo)
}
