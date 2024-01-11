package rooms

import "golang/Home/House/furniture"

// функция проверяет, можно ли установить мебель в комнате

func canFitAllFurniture(furniture []furniture.Furniture, room Room, newFurniture furniture.Furniture) bool {
	totalVolume := 0
	for _, f := range furniture {
		totalVolume += f.Length * f.Width * f.Height
	}
	totalVolume += newFurniture.Length * newFurniture.Width * newFurniture.Height

	roomVolume := room.Length * room.Width * room.Height

	return totalVolume <= roomVolume
}
