package plumbing

import "fmt"

// Метод для установки устройства

func (device *Plumbing) Install() {
	device.Installed = true
	fmt.Println("Устройство успешно установлено!")
}

// Метод для удаления устройства

func (device *Plumbing) Uninstall() {
	device.Installed = false
	fmt.Println("Устройство успешно удалено!")
}
