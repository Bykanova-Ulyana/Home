package appliances

// бытовые приборы

type Microwave struct { // СВЧ
	Quantity        int    // количество
	CountryOfOrigin string // страна-производитель
	Colour          string // цвет
	Turntable       bool   // Есть ли поворотный стол
	Volume          string // объём
	Brand           string // бренд
}

type Refrigerator struct { // холодильник
	Quantity           int    // количество
	EnergyClass        string // Класс энергопотребления
	Color              string // цвет
	MinimumTemperature int    // Минимальная температура
	MaximumTemperature int    // Максимальная температура
	PhoneControl       bool   // Управление с телефона
	Width              int    // ширина
	Height             int    // высота
	Freezer            bool   // Наличие морозилки
}

type TV struct { // телевизор
	DiagonalInch      int    // Диагональ, дюймы
	ScreenRefreshRate int    // Частота обновления экрана
	MaximumPower      int    // Максимальная потребляемая мощность
	Width             int    // Ширина
	Height            int    // Высота
	Model             string // Модель
	SmartTV           bool   // Поддержка Smart TV
}

type HomePhone struct { // Домашний телефон
	Wireless          bool   // Проводной или беспроводной
	Brand             string // Бренд
	NumberOfRingtones int    // Количество рингтонов
	AlarmClock        bool   // Наличие будильника
	NumberOfTubes     int    // Количество трубок
}

type Multicooker struct { // мультиварка
	Color         string // Цвет
	Power         int    // Мощность
	Height        int    // Высота
	Width         int    // Ширина
	TypeOfControl string // Тип контроля
	Display       bool   // Наличие дисплея
}

type CoffeeMachine struct { // Кофемашина
	VolumeWater       int  // Объём воды
	CupHeating        bool // Нагрев чашки
	AutomaticShutdown bool // Автоматическое выключение
	NumberCups        int  // Количество чашек
	VoiceAssistant    bool // Наличие голосового помощника
}

type Сooktop struct { // Варочная панель
	Type            string // Тип
	Installation    string // Тип установки
	NumberOfBurners int    // Количество конфорок
	PanelControl    string // Вид панели управления
	NumPowerLev     int    // Количество уровней мощности
}

type Dishwasher struct { // Посудомоечная машина
	NumberSets       int    // Количество загружаемых наборов
	NumberBaskets    int    // Количество корзин
	WaterConsumption int    // Расход воды
	Color            string // Цвет
	ProtectionLeaks  bool   // Наличие защиты от утечек
}

type Fan struct { // Вентилятор
	Color         string // Цвет
	Installation  string // Тип установки
	Power         int    // Мощность
	NumberSpeeds  int    // Количество скоростей
	NumberBlades  int    // Количество лопастей
	Height        int    // Высота
	RemoteControl bool   // Наличие ДУ
}
