package rooms

import (
	"golang/Home/House/appliances"
	"golang/Home/House/crockery"
	"golang/Home/House/family"
	"golang/Home/House/furniture"
	"golang/Home/House/plumbing"
	"golang/Home/House/windows"
)

type Room struct {
	Type       string
	Width      float32 // Ширина
	Length     float32 // Длина
	Height     float32 // Высота
	Area       float32 // Площадь
	Appliances []appliances.Appliance
	Family     []family.FamilyMember
	Crockery   []crockery.Crockery
	Plumbings  []plumbing.Plumbing
	Furniture  []furniture.Furniture
	Pets       []family.Pets
	Window     windows.Window
}

type Home struct {
	Rooms []Room
}
