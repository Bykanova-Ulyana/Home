package main

import (
	"fmt"
	"golang/House"
)

func main() {

	microwave := House.Microwave{
		Brand: "Газель",
	}
	fmt.Print(microwave)
	fmt.Printf("Введите бренд газели")
	fmt.Scanf("%s\n", &microwave.Brand)

	fmt.Print(microwave)
}

func show() {

}
