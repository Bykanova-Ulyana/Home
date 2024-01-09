package main

import (
	"fmt"
	"golang/Home/House/appliances"
	"golang/Home/House/crockery"
	"golang/Home/House/family"
	"golang/Home/House/plumbing"
	"golang/Home/House/rooms"
	"golang/Home/House/windows"
)

func main() {
	fmt.Printf("Всё работает!")
	println(appliances.CoffeeMachine{NumberCups: 5})
	println(crockery.Containers{})
	println(family.Mother{Works: true})
	println(plumbing.Bath{})
	println(rooms.Kitchen{})
	println(windows.Window{})
	mom := &Mother{Person{Name: "Anna", Surname: "Bykanova", Patronymic: "Nikolayevna", Age: 45, Birthday: "21.03.1978", Gender: true},
		Works:  true,
		Salary: 26000}
}
