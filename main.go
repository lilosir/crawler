package main

import (
	"firstCrawler/engine"
	"firstCrawler/scheduler"
	"firstCrawler/zhenai/parser"
	"fmt"
)

func main() {
	fmt.Println("start...")

	//which engine should run
	// currentEngine := engine.SimpleEngine{}

	currentEngine := engine.ConcurrentEngine{
		Scheduler:     &scheduler.SimpleScheduler{},
		WorkerCounter: 10,
	}

	currentEngine.Run(engine.Request{
		URL:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
