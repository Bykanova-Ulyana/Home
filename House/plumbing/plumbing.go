package plumbing

// сантехника

type Plumbing struct {
	Type      string // Тип устройства (например, раковина, унитаз, душ и т.д.)
	Brand     string // Бренд устройства
	Model     string // Модель устройства
	Material  string // Материал, из которого сделано устройство
	Installed bool   // Флаг, указывающий, установлено ли устройство
}
