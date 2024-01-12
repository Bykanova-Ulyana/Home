package main

import (
	"golang/Home/House/rooms"
)

func main() {
	var home rooms.Home
	home.CreatHome()
	rooms.PrintHome(home)
}
