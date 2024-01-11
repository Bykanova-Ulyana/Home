package crockery

import "fmt"

/*func CreateCrockery() TableSetting {

	tableSet := []tableSetting{{Cup: Cups{Crockery: Crockery{Number: 6, Material: "Серебро", Engraving: true}}},
		{Plate: Plates{Crockery: Crockery{Number: 6, Material: "Фарфор", Engraving: true}}},
		{Spoon: Spoons{Crockery: Crockery{Number: 6, Material: "Серебро", Engraving: false}}},
		{Fork: Forks{Crockery: Crockery{Number: 6, Material: "Серебро", Engraving: false}}}}

	return TableSetting{tableSet}
}*/

// Функция для ввода данных о посуде

func GetInputForCrockery() Crockery {
	var crockery Crockery
	var purpose int
	fmt.Println("Введите название посуды:")
	fmt.Scanln(&crockery.Type)
	fmt.Println("Введите цифру - назначение посуды:\n 1. Сервировка стола\n2. Кухонная утварь\n3. Посуда для хранения")
	fmt.Scanln(purpose)
	crockery.Purpose = purposeTranscript(purpose)
	fmt.Println("Введите материал посуды:")
	fmt.Scanln(&crockery.Material)
	fmt.Println("Есть ли гравировка на посуде (true/false):")
	fmt.Scanln(&crockery.Engraving)
	fmt.Println("Введите вместимость посуды:")
	fmt.Scanln(&crockery.Capacity)
	return crockery
}

func purposeTranscript(pr int) string {
	switch pr {
	case 1:
		return "Сервировка стола"
	case 2:
		return "Кухонная утварь"
	case 3:
		return "Посуда для хранения"
	}
	return "Назначение не выбрано!"
}

// Функция для вывода информации о посуде
func printCrockery(crockery Crockery) {
	engraving := "без гравировки"
	if crockery.Engraving {
		engraving = "есть гранировка"
	}
	fmt.Printf("%s\nПредназначение: %s\n Материал: %s\nКоличество: %d\nНаличие гравировки: %d\nОбъём: %d", crockery.Type,
		crockery.Purpose, crockery.Material, crockery.Number, engraving, crockery.Capacity)
}
