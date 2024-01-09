package crockery

// Посуда

type Crockery struct {
	Number    int    // Количество
	Material  string // Материал изготовления
	Engraving bool   // Наличие гравировки
}

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

type TableSetting struct { // Сервировка стола
	Number int
	Cup    Cups
	Plate  Plates
	Spoon  Spoons
	Fork   Forks
}

type Cookware struct { // Кухонная утварь
	Number int
	Pot    Pots
	Knife  Knives
	Lid    Lids
	Pan    Pans
}

type StorageUtensils struct { // Посуда для хранения
	Number    int
	lid       Lids
	Container Containers
}
