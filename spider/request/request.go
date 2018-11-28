package request

type CityListRequest struct {
	URL      string
	CityName string
}

type CityListParseResult struct {
	CityListParseResults []CityListRequest
}

//解析每个城市下的人的列表
type CityPersonListRquest struct {
	URL      string
	UserName string
	Gender   string
	Age      string
	Height   string
	Income   string
	Marriage string
}
type CityPersonsParseResult struct {
	CityPersonsParseResult []CityPersonListRquest
}
