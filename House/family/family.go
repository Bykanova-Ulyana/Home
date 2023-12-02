package family

type Person struct {
	Name       string // Имя
	Surname    string // Фамилия
	Patronymic string // Отчество
	Age        int    // Возраст
	Birthday   string // Дата рождения
	Gender     bool   // Пол
}

type Mother struct { // Мама
	Mam    Person // Основные параметры: ФИО, возраст, дата рождения, пол
	Works  bool   // Наличие работы
	Salary int    // Зарплата, руб.
}

type Father struct { // Папа
	Dad    Person // Основные параметры: ФИО, возраст, дата рождения, пол
	Works  bool   // Наличие работы
	Salary int    // Зарплата, руб.
}

type Sons struct {
	Son     Person // Основные параметры: ФИО, возраст, дата рождения, пол
	Student bool   // Учится или нет
	Works   bool   // Работает или нет
}

type Daughters struct {
	Daughter Person // Основные параметры: ФИО, возраст, дата рождения, пол
	Student  bool   // Учится или нет
	Works    bool   // Работает или нет
}

type Pets struct {
	TypeAnimal string // Вид животного
	Breed      string // Порода животного
	Name       string // Имя животного
	Age        int    // Возраст питомца
	Colour     string // Окрас питомца
	Weight     int    // Масса питомца
	SexAnimal  bool   // Пол животного: 1 - женский, 0 - мужской
}
