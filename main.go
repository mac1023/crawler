package main

import (
	"gocode/crawler/engine"
	"gocode/crawler/zhenai/parser"
)

func main() {

	url := "http://www.zhenai.com/zhenghun"
	engine.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
	})
}
