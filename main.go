package main

import (
	"fmt"
	"golang/home/House/rooms"
)

func main() {
	var home rooms.Home
	home = home.CreatHome()
	fmt.Print("\nВЫВЕДЕМ ДОМ!\n\n")
	rooms.PrintHome(home)
	/*for _, room := range Home.Rooms {
		for _, family := range room.Family {
			room = append(family.Family, family.GetInputForFamilyMember())
		}

	}*/
	println("Готово!")
}
