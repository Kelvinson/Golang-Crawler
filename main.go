// main
package main

import (
	"github.com/kelvinson/go/crawler/engine"
	"github.com/kelvinson/go/crawler/scheduler"
	"github.com/kelvinson/go/crawler/indeed/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		//	Scheduler:   &scheduler.SimpleScheduler{},
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url:        "https://www.indeed.com/companies",
		ParserFunc: parser.ParseCityList,
	})

}
