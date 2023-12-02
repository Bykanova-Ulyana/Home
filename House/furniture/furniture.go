package furniture

type Table struct {
	Height             int
	Width              int
	Color              string
	Material           string
	Purpose            string
	NumberShelves      int
	NumberDrawers      int
	MoistureResistance bool
}

type Bed struct {
	Number   int
	Height   int
	Width    int
	Color    string
	Material string
}

type Chair struct {
	Number   int
	Height   int
	Width    int
	Color    string
	Material string
}

type Armchair struct {
	Number   int
	Height   int
	Width    int
	Color    string
	Material string
}

type Cot struct { // Детская кроватка
	Number   int
	Height   int
	Width    int
	Color    string
	Material string
}

type CoffeeTable struct {
	Number   int
	Height   int
	Width    int
	Color    string
	Material string
}

type Sideboard struct { // сервант
	Height   int
	Width    int
	Color    string
	Material string
	Number   int
}

type Sofa struct { // диван
	Height   int
	Width    int
	Color    string
	Material string
	Number   int
}

type Wardrobe struct { // Шкаф для одежды
	Height   int
	Width    int
	Color    string
	Material string
	Number   int
}

type BedsideTable struct { // Прикроватная тумбочка
	Height   int
	Width    int
	Color    string
	Material string
	Number   int
}
