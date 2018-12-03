package faceplusplus

import (
	"fmt"
	"testing"
)

func TestGetFaceScore(*testing.T) {
	faceInfo := FaceInfo{
		ImgURL:      "imges/img.jpg",
		FemaleScore: 0,
		MaleScore:   0,
	}
	GetFaceScore(&faceInfo)
	fmt.Printf("%v", faceInfo)
}
