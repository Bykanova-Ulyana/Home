package rooms

import (
	"golang/Home/House/appliances"
	"golang/Home/House/crockery"
	"golang/Home/House/furniture"
	"golang/Home/House/plumbing"
	"golang/Home/House/windows"
)

// комнаты (попытка 1)

type Bedroom struct {
	TV           appliances.TV
	Table        furniture.Table
	Bed          furniture.Bed
	BedsideTable furniture.BedsideTable
	Wardrobe     furniture.Wardrobe
	Cot          furniture.Cot
	Window       windows.Window
}

type LivingRoom struct { // гостиная
	Table     furniture.Table
	TV        appliances.TV
	Armchair  furniture.Armchair
	Sofa      furniture.Sofa
	Fan       appliances.Fan
	Sideboard furniture.Sideboard
	Window    windows.Window
}

type Kitchen struct {
	Sink            plumbing.Sink
	Table           furniture.Table
	Chair           furniture.Chair
	CoffeeMachine   appliances.CoffeeMachine
	TableSetting    crockery.TableSetting
	Cookware        crockery.Cookware
	StorageUtensils crockery.StorageUtensils
	Dishwasher      appliances.Dishwasher
	Multicooker     appliances.Multicooker
	Refrigerator    appliances.Refrigerator
	Window          windows.Window
}

type Hallway struct { // коридор
	Window      windows.Window
	HomePhone   appliances.HomePhone
	CoffeeTable furniture.CoffeeTable
	Wardrobe    furniture.Wardrobe
}

type Bathroom struct {
	Bath   plumbing.Bath
	Sink   plumbing.Sink
	Faucet plumbing.Faucet
	Toilet plumbing.Toilet
}
