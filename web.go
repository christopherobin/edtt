package main

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"math"
	"net/http"
)

func ListSystemHandler(r render.Render) {
	var systemList = []map[string]interface{}{}

	for name, system := range systems {
		systemList = append(systemList, map[string]interface{}{
			"name":       name,
			"x":          system.X,
			"y":          system.Y,
			"z":          system.Z,
			"economy":    system.Economy,
			"allegiance": system.Allegiance,
		})
	}

	r.JSON(200, systemList)
}

func GetSystemHandler(params martini.Params, r render.Render) {
	if !SystemExists(params["system"]) {
		r.Status(404)
		return
	}

	system := GetSystem(params["system"])
	var stations []string

	for name := range system.Stations {
		stations = append(stations, name)
	}

	r.JSON(200, map[string]interface{}{
		"name":       system.Name,
		"x":          system.X,
		"y":          system.Y,
		"z":          system.Z,
		"economy":    system.Economy,
		"allegiance": system.Allegiance,
		"stations":   stations,
	})
}

type SystemForm struct {
	Name       string  `form:"name" binding:"required"`
	X          float64 `form:"x"`
	Y          float64 `form:"y"`
	Z          float64 `form:"z"`
	Economy    string  `form:"economy" binding:"required"`
	Allegiance string  `form:"allegiance" binding:"required"`
}

func CreateSystemHandler(systemForm SystemForm, r render.Render) {
	system := GetSystem(systemForm.Name)

	system.X = systemForm.X
	system.Y = systemForm.Y
	system.Z = systemForm.Z
	system.Economy = systemForm.Economy
	system.Allegiance = systemForm.Allegiance

	Save()

	r.JSON(200, map[string]interface{}{
		"name":       system.Name,
		"x":          system.X,
		"y":          system.Y,
		"z":          system.Z,
		"economy":    system.Economy,
		"allegiance": system.Allegiance,
	})
}

type StationForm struct {
	Name string `form:"name" binding:"required"`
}

func CreateStationHandler(stationForm StationForm, params martini.Params, r render.Render) {
	if !SystemExists(params["system"]) {
		r.Status(404)
		return
	}

	Save()

	r.JSON(200, GetSystem(params["system"]).GetStation(stationForm.Name))
}

func GetStationHandler(params martini.Params, r render.Render) {
	if !SystemExists(params["system"]) {
		r.Status(404)
		return
	}

	station := GetSystem(params["system"]).GetStation(params["station"])

	market := station.Market.ByCategory()

	res := map[string]interface{}{
		"name":     station.Name,
		"market":   market,
		"services": station.Services,
	}

	r.JSON(200, res)
}

type MarketEntryForm struct {
	Sell   int `form:"sell"`
	Buy    int `form:"buy"`
	Demand int `form:"demand"`
	Supply int `form:"supply"`
}

func SetMarketEntry(marketEntry MarketEntryForm, params martini.Params, r render.Render) {
	if !SystemExists(params["system"]) {
		r.Status(404)
		return
	}

	entry, ok := GetSystem(params["system"]).GetStation(params["station"]).Market[params["entry"]]
	if !ok {
		r.Status(404)
		return
	}

	if marketEntry.Sell != 0 {
		entry.Sell = int(math.Max(0, float64(marketEntry.Sell)))
	}

	if marketEntry.Buy != 0 {
		entry.Buy = int(math.Max(0, float64(marketEntry.Buy)))
	}

	if marketEntry.Demand != 0 {
		entry.Demand = int(math.Max(0, float64(marketEntry.Demand)))
	}

	if marketEntry.Supply != 0 {
		entry.Supply = int(math.Max(0, float64(marketEntry.Supply)))
	}

	Save()

	r.JSON(200, entry)
}

func SetService(req *http.Request, params martini.Params, r render.Render) {
	if !SystemExists(params["system"]) {
		r.Status(404)
		return
	}

	station := GetSystem(params["system"]).GetStation(params["station"])

	// decode JSON
	var payload = make(map[string]bool)
	d := json.NewDecoder(req.Body)
	err := d.Decode(&payload)

	if err != nil {
		r.Error(500)
		return
	}

	for service, availability := range payload {
		station.Services[service] = availability
	}

	Save()

	r.JSON(200, station)
}

type TradeFinderForm struct {
	Range float64 `form:"range"`
}

func TradeFinder(tradeFinder TradeFinderForm, params martini.Params, r render.Render) {
	if !SystemExists(params["system"]) {
		r.Status(404)
		return
	}

	if params["other"] != "ANY" && !SystemExists(params["other"]) {
		r.Status(404)
		return
	}

	if params["other"] == "ANY" && tradeFinder.Range == 0.0 {
		r.Status(400)
		return
	}

	system := GetSystem(params["system"])

	var dests []*System
	if params["other"] != "ANY" {
		dests = []*System{GetSystem(params["other"])}
	} else {
		dests = system.FindNearby(tradeFinder.Range)
	}

	res := []map[string]interface{}{}
	for from, station := range system.Stations {
		for _, other := range dests {
			for to, otherStation := range other.Stations {
				distance := system.Distance(other)
				trades := station.Market.Compare(otherStation.Market)

				if len(trades) == 0 {
					continue
				}

				routes := map[string]interface{}{
					"from": map[string]interface{}{
						"system":  system.Name,
						"station": from,
					},
					"to": map[string]interface{}{
						"system":  other.Name,
						"station": to,
					},
					"distance": distance,
					"trades":   trades,
				}
				res = append(res, routes)
			}
		}
	}

	r.JSON(200, res)
}

func GetGoods(r render.Render) {
	res := make(map[string]map[string]interface{})
	for _, good := range Goods {
		_, ok := res[good.Category]
		if !ok {
			res[good.Category] = make(map[string]interface{})
		}
		res[good.Category][good.Name] = map[string]interface{}{
			"name":         good.Name,
			"galactic_avg": good.GalacticAvg,
		}
	}
	r.JSON(200, res)
}

type GoodFinderForm struct {
	From  string  `form:"form"`
	Range float64 `form:"range"`
}

type FoundSystem struct {
	Distance float64        `json:"distance"`
	Stations []FoundStation `json:"stations"`
}

type FoundStation struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func FindGoods(goodFinder GoodFinderForm, params martini.Params, r render.Render) {
	good, ok := Goods[params["good"]]

	if !ok {
		r.Status(404)
		return
	}

	if !SystemExists(goodFinder.From) {
		r.Status(404)
		return
	}

	from := GetSystem(goodFinder.From)
	targets := from.FindNearby(goodFinder.Range)

	res := make(map[string]FoundSystem)
	for _, system := range targets {
		for _, station := range system.Stations {
			if entry, ok := station.Market[good.Name]; ok {
				if entry.Buy > 0 {
					if _, ok := res[system.Name]; !ok {
						res[system.Name] = FoundSystem{
							Distance: from.Distance(system),
							Stations: make([]FoundStation, 0),
						}
					}
					tmp := res[system.Name]
					tmp.Stations = append(tmp.Stations, FoundStation{
						Name:  station.Name,
						Price: entry.Buy,
					})
					res[system.Name] = tmp
				}
			}
		}
	}

	r.JSON(200, res)
}
