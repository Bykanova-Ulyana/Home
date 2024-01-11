package crockery

// Посуда

type Crockery struct {
	Type      string // Вид прибора
	Purpose   string // Назначение
	Number    int    // Количество
	Material  string // Материал изготовления
	Engraving bool   // Наличие гравировки
	Capacity  int    // Вместимость
}

/*
type Spoons struct { // Ложки
	Crockery Crockery
}

type Forks struct { // Вилки
	Crockery Crockery
}

type Knives struct { // Ножи
	Crockery Crockery
}

type Plates struct { // Тарелки
	Crockery Crockery
}

type Cups struct { // Чашки
	Crockery Crockery
}

type Pots struct { // Кастрюли
	Crockery Crockery
}

type Pans struct { // Сковородки
	Crockery Crockery
}

type Lids struct { // Крышки
	Crockery Crockery
}
type Containers struct { // Контейнеры
	Crockery Crockery
}

type tableSetting struct { // Сервировка стола
	Cup   Cups
	Plate Plates
	Spoon Spoons
	Fork  Forks
}

type cookware struct { // Кухонная утварь
	Pot   Pots
	Knife Knives
	Lid   Lids
	Pan   Pans
}

type storageUtensils struct { // Посуда для хранения
	Lid       Lids
	Container Containers
}

type TableSetting struct {
	TableSetting []tableSetting
}

type Cookware struct {
	Cookware []cookware
}

type StorageUtensils struct {
	StorageUtensils []storageUtensils
}
*/
