package main

import (
	"fmt"
	"goProjects/house/furniture"
	"goProjects/house/rooms"
)

type House struct {
	Kitchen rooms.Kitchen
	Bedroom rooms.Bedroom
}

func main() {
	h := House{Kitchen: rooms.Kitchen{Furniture: furniture.CreateKitchenSet(), Room: rooms.CreateRoom("Кухня")},
		Bedroom: rooms.Bedroom{Room: rooms.CreateRoom("Спальня"), Furniture: furniture.CreateBedroomSet()},
	}

	//h.kitchen.Furniture = h.CreateKitchenSet()
	fmt.Println("\n\n\n\n", h)
}
