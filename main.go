package main

import (
	"firstCrawler/engine"
	"firstCrawler/scheduler"
	"firstCrawler/zhenai/parser"
	"fmt"
)

func main() {
	fmt.Println("start...")

	//simple engine
	// currentEngine := engine.SimpleEngine{}

	// concurrent engines, simple or queue
	currentEngine := engine.ConcurrentEngine{
		// Scheduler:     &scheduler.SimpleScheduler{},
		Scheduler:     &scheduler.QueueScheduler{},
		WorkerCounter: 50,
	}

	currentEngine.Run(engine.Request{
		URL:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
	// currentEngine.Run(engine.Request{
	// 	URL:       "http://www.zhenai.com/zhenghun/shanghai",
	// 	ParseFunc: parser.ParseCity,
	// })
}
