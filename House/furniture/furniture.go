package furniture

import (
	"fmt"
)

type Furniture struct {
	Type     string
	Number   int
	Height   int
	Width    int
	Color    string
	Material string
}

type Table struct {
	table              Furniture
	Purpose            string
	NumberShelves      int
	NumberDrawers      int
	MoistureResistance bool
}

type Bed struct {
	bed Furniture
}

type Chair struct {
	chair Furniture
}

type Armchair struct {
	armchair Furniture
}

type Cot struct { // Детская кроватка
	cot Furniture
}

type CoffeeTable struct {
	coffeeTable Furniture
}

type Sideboard struct { // сервант
	sideboard Furniture
}

type Sofa struct { // диван
	sofa Furniture
}

type Wardrobe struct { // Шкаф для одежды
	wardrobe Furniture
}

type BedsideTable struct { // Прикроватная тумбочка
	bedsideTable Furniture
}

func (f Furniture) FurniturePrint() {
	fmt.Println("\n\tтип мебели: ", f.Type, "\n\tцвет: ", f.Color, "\n\tматериал: ", f.Material, "\n\tколичество: ", f.Number, "\n\tширина: ", f.Width, "\n\tвысота: ", f.Height, "\n")
}
