package main

import (
	"firstCrawler/engine"
	"firstCrawler/zhenai/parser"
	"fmt"
)

func main() {
	fmt.Println("start...")
	engine.Run(engine.Request{
		URL:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
