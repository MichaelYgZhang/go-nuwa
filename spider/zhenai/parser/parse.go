package parser

import (
	"encoding/json"
	"regexp"
	"spider/request"

	"github.com/tidwall/gjson"
)

var (
	cityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	userRe     = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	genderRex  = regexp.MustCompile(`性别：</span>([^<]+)</td>`)
	ageRe      = regexp.MustCompile(`年龄：</span>([^<]+)</td>`)
	heightRe   = regexp.MustCompile(`身   高：</span>([^<]+)</td>`)
	incomeRe   = regexp.MustCompile(`月   薪：</span>([^<]+)</td>`)
	marriageRe = regexp.MustCompile(`婚况：</span>([^<]+)</td>`)
	cityUrlRe  = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
	pageRe     = regexp.MustCompile(`INITIAL_STATE__=([^;(]+)`)
)

//解析所有的城市列表
func ParseCityList(contents []byte) request.CityListParseResult {
	matches := cityListRe.FindAllSubmatch(contents, -1)
	result := request.CityListParseResult{}
	for _, match := range matches {
		result.CityListParseResults = append(
			result.CityListParseResults,
			request.CityListRequest{
				URL:      string(match[1]),
				CityName: string(match[2]),
			})
	}
	return result
}

type PageInfo struct {
	MemberListData    string
	RecommendListData string
	FooterData        string
	NavigationData    string
}
type Student struct {
	Name   string
	Sex    int
	Height int
}

type Person struct {
	Age              int
	AvatarURL        string
	Education        string
	Height           int
	IntroduceContent string
	Marriage         string
	MemberID         int64
	NickName         string
	Occupation       string
	Salary           string
	Sex              int
	WorkCity         string
	FemaleScore      float64
	MaleScore        float64
}

//解析每个城市下面的人的列表
func ParseCitysPerson(contents []byte) []Person {
	pages := pageRe.FindAllSubmatch(contents, -1)
	var persons []Person
	for _, m := range pages {
		page := string(m[1])
		memberListData := gjson.Get(page, "memberListData")
		memberList := memberListData.Get("memberList")
		json.Unmarshal([]byte(memberList.String()), &persons)
		return persons
	}
	return nil
}
