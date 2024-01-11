package furniture

import "fmt"

func (f Furniture) FurniturePrint() {
	fmt.Println("\n\tтип мебели: ", f.Type, "\n\tцвет: ", f.Color, "\n\tматериал: ", f.Material, "\n\tколичество: ", f.Number, "\n\tширина: ", f.Width, "\n\tвысота: ", f.Height, "\n")
}

func getInputForFurniture() (Furniture, error) {
	var furniture Furniture
	var err [7]error
	fmt.Println("Тип мебели:")
	_, err[0] = fmt.Scanln(&furniture.Type)

	fmt.Println("Количество предметов:")
	_, err[1] = fmt.Scanln(&furniture.Number)

	fmt.Println("Цвет мебели:")
	_, err[2] = fmt.Scanln(&furniture.Color)

	fmt.Println("Высота:")
	_, err[3] = fmt.Scanln(&furniture.Height)

	fmt.Println("Длина:")
	_, err[4] = fmt.Scanln(&furniture.Length)

	fmt.Println("Ширина:")
	_, err[5] = fmt.Scanln(&furniture.Width)

	fmt.Println("Матеиал изготовления:")
	_, err[6] = fmt.Scanln(&furniture.Material)

	for _, value := range err {
		if value != nil {
			return furniture, value
		}
	}
	return furniture, nil
}
