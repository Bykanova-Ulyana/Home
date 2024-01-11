package rooms

import (
	"golang/Home/House/furniture"
)

// Функция, которая проверяет, можно ли установить предмет на полу и влезет ли он по высоте

func CanFitOnFloor(furniture []furniture.Furniture, room Room, newFurniture furniture.Furniture) bool {
	// Проверяем, достаточно ли места на полу для новой мебели
	if newFurniture.Length <= room.Length && newFurniture.Width <= room.Width {
		// Проверяем, влезет ли новая мебель по высоте
		heightLimit := room.Height - newFurniture.Height
		for _, item := range furniture {
			if item.Height > heightLimit {
				return false
			}
		}
		return true
	}
	return false
}

// считаем площадь комнаты

func (r Room) AreaCalculate() float32 {
	return r.Width * r.Length
}
