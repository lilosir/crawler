package parser

import (
	"firstCrawler/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank"><img src="https://photo.zastatic.com/images/photo/[0-9]+/[0-9]+/[0-9]+[^<>]+alt="([^>]+)"></a>`)
	//In city details page also contains others cities or next page
	cityURLRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

//ParseCity returns all the users in a byte slice
func ParseCity(contents []byte) engine.ParseResult {
	profileMatches := profileRe.FindAllSubmatch(contents, -1)
	citiesMatches := cityURLRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, profile := range profileMatches {
		name := string(profile[2])
		// result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			URL: string(profile[1]),
			ParseFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			},
		})
	}

	for _, city := range citiesMatches {
		result.Requests = append(result.Requests, engine.Request{
			URL:       string(city[1]),
			ParseFunc: ParseCity,
		})
	}

	return result
}
