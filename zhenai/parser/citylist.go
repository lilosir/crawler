package parser

import (
	"firstCrawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

//ParseCityList returns all the citys in a byte slice
func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	limitation := 5
	for _, match := range matches {
		result.Items = append(result.Items, "City "+string(match[2]))
		result.Requests = append(result.Requests, engine.Request{
			URL:       string(match[1]),
			ParseFunc: ParseCity,
		})
		limitation--
		if limitation == 0 {
			break
		}
	}
	return result
}
