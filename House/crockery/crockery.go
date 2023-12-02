package crockery

// Посуда

type Spoons struct { // Ложки
	Number    int    // Количество
	Material  string // Материал изготовления
	Engraving bool   // Наличие гравировки
}

type Forks struct { // Вилки
	Number    int    // Количество
	Material  string // Материал изготовления
	Engraving bool   // Наличие гравировки
}

type Knives struct { // Ножи
	Number    int    // Количество
	Material  string // Материал изготовления
	Engraving bool   // Наличие гравировки
}

type Plates struct { // Тарелки
	Number    int    // Количество
	Material  string // Материал изготовления
	Engraving bool   // Наличие гравировки
}

type Cups struct { // Чашки
	Number    int    // Количество
	Material  string // Материал изготовления
	Engraving bool   // Наличие гравировки
}

type Pots struct { // Кастрюли
	Number    int    // Количество
	Material  string // Материал изготовления
	Engraving bool   // Наличие гравировки
}

type Pans struct { // Сковородки
	Number    int    // Количество
	Material  string // Материал изготовления
	Engraving bool   // Наличие гравировки
}

type Lids struct { // Крышки
	Number    int    // Количество
	Material  string // Материал изготовления
	Engraving bool   // Наличие гравировки
}
type Containers struct { // Контейнеры
	Number    int    // Количество
	Material  string // Материал изготовления
	Engraving bool   // Наличие гравировки
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
