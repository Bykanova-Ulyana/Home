package windows

import "fmt"

func createWindow() Window {
	window := Window{Number: 1, Length: 1.2, Width: 1, OpenWindow: false}
	return window
}

func StatusOpenCloseWindow(w Window, command bool) {
	switch command {
	case command == true || w.OpenWindow == true:
		println("Окно открыто")
	case false:
		println("Окно закрыто")
	default:
		println("ОШИБКА! ОКНО СЛОМАЛОСЬ!")
	}
}

func printWindow(w Window) {
	fmt.Printf("Длина окна %f, ширина окна %f\n", w.Length, w.Width)
}
