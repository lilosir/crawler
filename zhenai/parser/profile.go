package parser

import (
	"bytes"
	"firstCrawler/engine"
	"firstCrawler/model"
	"log"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var ageRe = regexp.MustCompile(`([0-9]+)岁`)
var zodiacRe = regexp.MustCompile(`(.+)座`)
var heightRe = regexp.MustCompile(`([0-9]+)cm`)
var areaRe = regexp.MustCompile(`工作地:(.+)`)
var incomeRe = regexp.MustCompile(`月收入:(.+)`)
var genderRe = regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[^>]+">[\D]+([\D]{1})士征婚</a>`)

//ParseProfile reture all the useful info need to be known
func ParseProfile(contents []byte, name string) engine.ParseResult {

	profile := model.Profile{}

	r := bytes.NewReader(contents)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatal(err)
	}

	matchGender := genderRe.FindSubmatch(contents)
	if len(matchGender) == 2 {
		profile.Gender = string(matchGender[1])
	}

	doc.Find(".m-content-box .purple-btns .m-btn.purple[data-v-8b1eac0c]").Each(func(i int, s *goquery.Selection) {
		value := s.Text()
		profile.Name = name
		if i == 0 {
			profile.Marriage = value
		} else {
			if age, ok := extractInt(value, ageRe); ok {
				profile.Age = age
			}
			if height, ok := extractInt(value, heightRe); ok {
				profile.Height = height
			}
			if zodiac, ok := extractString(value, zodiacRe); ok {
				profile.Zodiac = zodiac
			}
			if area, ok := extractString(value, areaRe); ok {
				profile.Area = area
			}
			if income, ok := extractString(value, incomeRe); ok {
				profile.Income = income
			}
		}
	})

	doc.Find(".m-content-box .pink-btns .m-btn.pink[data-v-8b1eac0c]").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		value := s.Text()
		if income, ok := extractString(value, incomeRe); ok {
			profile.Income = income
		}
	})

	return engine.ParseResult{
		Items: []interface{}{profile},
	}
}

func extractString(value string, re *regexp.Regexp) (string, bool) {
	match := re.FindSubmatch([]byte(value))
	if len(match) >= 2 {
		return string(match[1]), true
	}
	return "", false
}

func extractInt(value string, re *regexp.Regexp) (int, bool) {
	match := re.FindSubmatch([]byte(value))
	if len(match) >= 2 {
		age, err := strconv.Atoi(string(match[1]))
		if err == nil {
			return age, true
		}
	}
	return 0, false
}
