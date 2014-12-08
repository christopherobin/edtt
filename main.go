package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"log"
)

func main() {
	err := Load()
	if err != nil {
		log.Fatalln(err)
	}

	defer (func() {
		err := Save()
		if err != nil {
			log.Fatalln(err)
		}
	})()

	m := martini.Classic()
	m.Use(martini.Static("dist"))
	m.Use(render.Renderer(render.Options{Delims: render.Delims{"{[{", "}]}"}}))

	m.Get("/api/systems", ListSystemHandler)
	m.Post("/api/systems", binding.Bind(SystemForm{}), CreateSystemHandler)
	m.Get("/api/systems/:system", GetSystemHandler)
	m.Post("/api/systems/:system", binding.Bind(StationForm{}), CreateStationHandler)
	m.Get("/api/systems/:system/:station", GetStationHandler)
	m.Post("/api/systems/:system/:station/market/:entry", binding.Bind(MarketEntryForm{}), SetMarketEntry)
	m.Post("/api/systems/:system/:station/services", SetService)
	m.Get("/api/market", GetGoods)
	m.Post("/api/market/:good", binding.Bind(GoodFinderForm{}), FindGoods)
	m.Post("/api/market/:system/:other", binding.Bind(TradeFinderForm{}), TradeFinder)

	m.Run()
}
