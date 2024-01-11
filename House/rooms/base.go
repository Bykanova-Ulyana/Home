package rooms

import (
	"golang/Home/House/appliances"
	"golang/Home/House/furniture"
)

// Описываем эту квартиру: https://www.cian.ru/rent/flat/296918422/

func CreatHome() []Home {
	home := []Home{{Rooms: []Room{{
		Type:   "Коридор",
		Width:  3.18,
		Length: 5.94,
		Height: 2,
		Area:   0,
		Appliances: []appliances.Appliance{{
			Name:   "Пылесос",
			Brand:  "Hyundai",
			Power:  220,
			Width:  0.3,
			Length: 0.29,
			Height: 0.49,
			IsOn:   false,
		}},
		Family:    nil,
		Crockery:  nil,
		Plumbings: nil,
		Furniture: []furniture.Furniture{{
			Type: "Кровать",
		}},
		Pets: nil,
	}}}}
	return home
}
