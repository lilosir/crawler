package main

import (
	"firstCrawler/engine"
	"firstCrawler/persist"
	"firstCrawler/scheduler"
	"firstCrawler/zhenai/parser"
	"fmt"
)

func main() {
	fmt.Println("start...")

	//simple engine
	// currentEngine := engine.SimpleEngine{}

	// concurrent engines, simple or queue
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	currentEngine := engine.ConcurrentEngine{
		// Scheduler:     &scheduler.SimpleScheduler{},
		Scheduler:     &scheduler.QueueScheduler{},
		WorkerCounter: 50,
		ItemChan:      itemChan,
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
