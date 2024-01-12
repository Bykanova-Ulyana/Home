package main

import (
	"golang/Home/House/rooms"
)

func main() {
	var home rooms.Home
	home = home.CreatHome()
	rooms.PrintHome(home)
	println("Готово!")
}
