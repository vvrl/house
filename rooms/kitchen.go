package rooms

import "goProjects/house/furniture"

type Kitchen struct {
	Room      Room
	Furniture furniture.KitchenSet
}

func Set(f Kitchen) {
	f.Furniture = furniture.CreateKitchenSet()
}
