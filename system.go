package main

import (
	"encoding/json"
	"log"
	"math"
	"os"
	"regexp"
	"strings"
)

type System struct {
	Name       string
	X          float64
	Y          float64
	Z          float64
	Economy    string
	Allegiance string
	Stations   map[string]*Station
}

type Station struct {
	Name     string
	Market   Market
	Services map[string]bool
}

var systems = make(map[string]*System)

func GetSystem(name string) *System {
	if system, ok := systems[name]; ok {
		return system
	}

	system := &System{
		Name:       name,
		Stations:   make(map[string]*Station),
		Economy:    "None",
		Allegiance: "None",
	}
	systems[name] = system

	return system
}

func SystemExists(name string) bool {
	_, ok := systems[name]
	return ok
}

func (system *System) GetStation(name string) *Station {
	if station, ok := system.Stations[name]; ok {
		return station
	}

	system.Stations[name] = &Station{
		Name:     name,
		Services: make(map[string]bool),
		Market:   NewMarket(),
	}

	return system.Stations[name]
}

func (system *System) Distance(other *System) float64 {
	log.Println(system, other)
	return math.Floor(math.Sqrt(math.Pow(system.X-other.X, 2)+math.Pow(system.Y-other.Y, 2)+math.Pow(system.Z-other.Z, 2))*10.0) / 10.0
}

func (system *System) FindNearby(max float64) []*System {
	var res []*System
	for _, other := range systems {
		/*if other == system {
			continue
		}*/

		if system.Distance(other) <= max {
			res = append(res, other)
		}
	}
	return res
}

var normRegexp = regexp.MustCompile("\\W+")

func (system *System) getDBFile() string {
	return "market/" + normRegexp.ReplaceAllString(strings.ToLower(system.Name), "_") + ".json"
}

func Save() error {
	for _, system := range systems {
		file, err := os.OpenFile(system.getDBFile(), os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return err
		}
		encoder := json.NewEncoder(file)
		encoder.Encode(system)
		file.Close()
	}

	return nil
}

func Load() error {
	dir, err := os.Open("market")
	if err != nil {
		return err
	}

	files, err := dir.Readdirnames(0)
	if err != nil {
		return err
	}

	for _, filename := range files {
		if !strings.HasSuffix(filename, ".json") {
			continue
		}

		file, err := os.Open("market/" + filename)
		if err != nil {
			return err
		}

		var system System
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&system)
		if err != nil {
			return err
		}

		for _, station := range system.Stations {
			if station.Services == nil {
				station.Services = make(map[string]bool)
			}
			station.Market.Check()
		}

		systems[system.Name] = &system
		file.Close()
	}

	return nil
}
