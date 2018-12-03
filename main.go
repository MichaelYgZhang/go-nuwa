package main

import (
	"fmt"
	"spider/faceplusplus"
	"spider/process"
	"spider/store"
	"spider/util"
	"spider/zhenai/parser"

	"gopkg.in/olivere/elastic"
)

const URL = "http://www.zhenai.com/zhenghun"

func main() {
	client, err := elastic.NewClient(
		// Must turn off sniff in docker
		elastic.SetSniff(true))

	if err != nil {
		panic(err)
	}

	body, err := process.ProcessPageBody(URL)
	if err != nil {
		panic(err)
	}
	parseCityListResult := parser.ParseCityList(body)

	for _, pr := range parseCityListResult.CityListParseResults {
		fmt.Printf("%s, %s", pr.URL, pr.CityName)
		fmt.Println()
	}

	//解析每一个城市的人的列表
	// TODO 分页
	for _, cityList := range parseCityListResult.CityListParseResults {
		bodyCity, err := process.ProcessPageBody(cityList.URL)
		if err != nil {
			continue
		}
		persons := parser.ParseCitysPerson(bodyCity)
		for _, person := range persons {
			faceInfo := faceplusplus.FaceInfo{
				ImgBase64:   util.GetImgFromURL(person.AvatarURL),
				FemaleScore: 0,
				MaleScore:   0,
			}
			faceplusplus.GetFaceScore(&faceInfo)
			person.FemaleScore = faceInfo.FemaleScore
			person.MaleScore = faceInfo.MaleScore
			if person.FemaleScore > 80 || person.FemaleScore > 80 {
				fmt.Printf("beauty person %v", person)
			}
			err := store.Save(client, "spider", "zhenai", person)
			if err != nil {
				fmt.Printf("%v", err)
				continue
			}
		}
	}
	//解析每个人的详情 分页查询
}
