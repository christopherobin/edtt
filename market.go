package main

import (
//"encoding/csv"
//"log"
//"os"
//"regexp"
//"strconv"
//"strings"
)

type Good struct {
	Name        string
	Category    string
	GalacticAvg int
}

type MarketEntry struct {
	Name   string
	Sell   int
	Buy    int
	Demand int
	Supply int
}

type Market map[string]*MarketEntry

var Goods = map[string]Good{
	// CHEMICALS
	"Explosives":    Good{"Explosives", "Chemicals", 372},
	"Hydrogen Fuel": Good{"Hydrogen Fuel", "Chemicals", 179},
	"Mineral Oil":   Good{"Mineral Oil", "Chemicals", 252},
	"Pesticides":    Good{"Pesticides", "Chemicals", 286},
	// CONSUMER ITEMS
	"Clothing":            Good{"Clothing", "Consumer Items", 389},
	"Consumer Technology": Good{"Consumer Technology", "Consumer Items", 7025},
	"Domestic Appliances": Good{"Domestic Appliances", "Consumer Items", 625},
	// DRUGS
	"Narcotics": Good{"Narcotics", "Drugs", 233},
	// FOODS
	"Algae":                Good{"Algae", "Foods", 193},
	"Animal Meat":          Good{"Animal Meat", "Foods", 1454},
	"Coffee":               Good{"Coffee", "Foods", 1454},
	"Fish":                 Good{"Fish", "Foods", 783},
	"Food Cartridges":      Good{"Food Cartridges", "Foods", 198},
	"Fruit and Vegetables": Good{"Fruit and Vegetables", "Foods", 389},
	"Grain":                Good{"Grain", "Foods", 268},
	"Synthetic Meat":       Good{"Synthetic Meat", "Foods", 318},
	"Tea":                  Good{"Tea", "Foods", 1640},
	// INDUSTRIAL MATERIALS
	"Polymers":        Good{"Polymers", "Industrial Materials", 209},
	"Semiconductors":  Good{"Semiconductors", "Industrial Materials", 1023},
	"Superconductors": Good{"Superconductors", "Industrial Materials", 7025},
	// LEGAL DRUGS
	"Beer":    Good{"Beer", "Legal Drugs", 233},
	"Liquor":  Good{"Liquor", "Legal Drugs", 732},
	"Tabacco": Good{"Tabacco", "Legal Drugs", 5084},
	"Wine":    Good{"Wine", "Legal Drugs", 318},
	// MACHINERY
	"Atmospheric Processors": Good{"Atmospheric Processors", "Machinery", 487},
	"Crop Harvesters":        Good{"Crop Harvesters", "Machinery", 2372},
	"Heliostatic Furnaces":   Good{"Heliostatic Furnaces", "Machinery", 260},
	"Marine Equipment":       Good{"Marine Equipment", "Machinery", 4469},
	"Mineral Extractors":     Good{"Mineral Extractors", "Machinery", 694},
	"Microbial Furnaces":     Good{"Microbial Furnaces", "Machinery", 260},
	"Power Generators":       Good{"Power Generators", "Machinery", 625},
	"Water Purifiers":        Good{"Water Purifiers", "Machinery", 372},
	// MEDECINES
	"Agri-Medecines":        Good{"Agri-Medecines", "Medecines", 1148},
	"Basic Medecines":       Good{"Basic Medecines", "Medecines", 389},
	"Combat Stabilisers":    Good{"Combat Stabilisers", "Medecines", 3049},
	"Performance Enhancers": Good{"Performance Enhancers", "Medecines", 7025},
	"Progenitor Cells":      Good{"Progenitor Cells", "Medecines", 7025},
	// METALS
	"Aluminium": Good{"Aluminium", "Metals", 406},
	"Beryllium": Good{"Beryllium", "Metals", 8543},
	"Cobalt":    Good{"Cobalt", "Metals", 817},
	"Copper":    Good{"Copper", "Metals", 564},
	"Gallium":   Good{"Gallium", "Metals", 5421},
	"Gold":      Good{"Gold", "Metals", 9737},
	"Indium":    Good{"Indium", "Metals", 6170},
	"Lithium":   Good{"Lithium", "Metals", 1744},
	"Palladium": Good{"Palladium", "Metals", 13522},
	"Silver":    Good{"Silver", "Metals", 5082},
	"Tantalum":  Good{"Tantalum", "Metals", 4191},
	"Titanium":  Good{"Titanium", "Metals", 1148},
	"Uranium":   Good{"Uranium", "Metals", 2863},
	// MINERALS
	"Bauxite":     Good{"Bauxite", "Minerals", 209},
	"Bertrandite": Good{"Bertrandite", "Minerals", 2689},
	"Coltan":      Good{"Coltan", "Minerals", 1544},
	"Gallite":     Good{"Gallite", "Minerals", 2096},
	"Indite":      Good{"Indite", "Minerals", 2372},
	"Lepidolite":  Good{"Lepidolite", "Minerals", 694},
	"Rutile":      Good{"Rutile", "Minerals", 406},
	"Uraninite":   Good{"Uraninite", "Minerals", 1023},
	// TECHNOLOGY
	"Advanced Catalysers":     Good{"Advanced Catalysers", "Technology", 3049},
	"Animal Monitors":         Good{"Animal Monitors", "Technology", 372},
	"Aquaponic Systems":       Good{"Aquaponic Systems", "Technology", 343},
	"Auto-Fabricators":        Good{"Auto-Fabricators", "Technology", 3931},
	"Bioreducing Lichen":      Good{"Bioreducing Lichen", "Technology", 1083},
	"Computer Components":     Good{"Computer Components", "Technology", 625},
	"H.E. Suits":              Good{"H.E. Suits", "Technology", 343},
	"Land Enrichment Systems": Good{"Land Enrichment Systems", "Technology", 5082},
	"Resonating Separators":   Good{"Resonating Separators", "Technology", 6170},
	"Robotics":                Good{"Robotics", "Technology", 1970},
	// TEXTILES
	"Leather":           Good{"Leather", "Textiles", 233},
	"Natural Fabrics":   Good{"Natural Fabrics", "Textiles", 487},
	"Synthetic Fabrics": Good{"Synthetic Fabrics", "Textiles", 245},
	// WASTE
	"Biowaste":       Good{"Biowaste", "Waste", 74},
	"Chemical Waste": Good{"Chemical Waste", "Waste", 126},
	"Scrap":          Good{"Scrap", "Waste", 96},
	// WEAPONS
	"Non-Lethal Weapons": Good{"Non-Lethal Weapons", "Weapons", 1970},
	"Personal Weapons":   Good{"Personal Weapons", "Weapons", 4469},
	"Reactive Armour":    Good{"Reactive Armour", "Weapons", 2229},
}

func NewMarket() Market {
	market := make(Market)

	// register every product in there
	market.Check()

	return market
}

func (market Market) Check() {
	for name, good := range Goods {
		_, ok := market[name]
		if !ok {
			market[name] = &MarketEntry{
				Name: good.Name,
			}
		}
	}
	for name, _ := range market {
		_, ok := Goods[name]
		if !ok {
			delete(market, name)
		}
	}
}

func (market Market) ByCategory() map[string]map[string]*MarketEntry {
	res := make(map[string]map[string]*MarketEntry)

	for name, entry := range market {
		category := entry.Good().Category

		_, ok := res[category]
		if !ok {
			res[category] = make(map[string]*MarketEntry)
		}

		res[category][name] = entry
	}

	return res
}

// Return a list of profitable routes from market to other
type Trade struct {
	Name    string
	Buy     int
	Sell    int
	Revenue int
}

func (market Market) Compare(other Market) []Trade {
	var res []Trade
	for name, entry := range market {
		otherEntry := other[name]

		// reverse find
		/*if otherEntry.Buy > 0 && entry.Sell > otherEntry.Buy {
			res[name] = Trade{
				Buy:     otherEntry.Buy,
				Sell:    entry.Sell,
				Revenue: otherEntry.Buy - entry.Sell,
			}
		}*/

		if entry.Buy > 0 && otherEntry.Sell > entry.Buy {
			res = append(res, Trade{
				Name:    name,
				Buy:     entry.Buy,
				Sell:    otherEntry.Sell,
				Revenue: otherEntry.Sell - entry.Buy,
			})
		}
	}

	return res
}

func (entry *MarketEntry) Good() Good {
	return Goods[entry.Name]
}
