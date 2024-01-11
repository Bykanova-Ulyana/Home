package family

import (
	"fmt"
	"strings"
)

// Функция для ввода данных о питомце
func getInputForPets() Pets {
	var pet Pets
	all := "кошка собака кролик хомяк попугай лошадь овца"

	fmt.Println("Введите вид животного: ")
	fmt.Scanln(&pet.TypeAnimal)

	pet.AllergicFactor = "нет"
	if strings.Contains(all, pet.TypeAnimal) {
		pet.AllergicFactor = "шерсть"
	}

	fmt.Println("Введите кличку животного: ")
	fmt.Scanln(&pet.Name)
	fmt.Println("Введите породу животного: ")
	fmt.Scanln(&pet.Breed)
	fmt.Println("Введите окрас животного: ")
	fmt.Scanln(&pet.Colour)
	fmt.Println("Пол животного(мальчик/девочка): ")
	fmt.Scanln(&pet.SexAnimal)
	fmt.Println("Вес животного: ")
	fmt.Scanln(&pet.Weight)

	return pet
}

// Функция для вывода информации о питомце

func PrintPet(pet Pets) {
	fmt.Printf("%s по кличке %s породы %s имеет %s окрас, пол: %s, вес: %d",
		pet.TypeAnimal, pet.Name, pet.Breed, pet.Colour, pet.SexAnimal, pet.Weight)
}
