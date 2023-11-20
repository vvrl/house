package rooms

import "goProjects/house/furniture"

type Kitchen struct {
	Room      Room
	Furniture furniture.KitchenSet
}

//func (f Kitchen) Init() {
//	f.Furniture = furniture.CreateKitchenSet()
//}
