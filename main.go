package main

import (
	"fmt"
	"goProjects/house/furniture"
	"goProjects/house/rooms"
)

type House struct {
	SumSquare float32
	Kitchen   rooms.Kitchen
	Bedroom   rooms.Bedroom
	Hall      rooms.Hall
	Bathroom  rooms.Bath
}

func main() {
	h := House{Kitchen: rooms.Kitchen{Furniture: furniture.CreateKitchenSet(), Room: rooms.CreateRoom("Кухня")},
		Bedroom:  rooms.Bedroom{Room: rooms.CreateRoom("Спальня"), Furniture: furniture.CreateBedroomSet()},
		Hall:     rooms.Hall{Room: rooms.CreateRoom("Прихожая"), Furniture: furniture.CreateHallSet()},
		Bathroom: rooms.Bath{Room: rooms.CreateRoom("Ванная"), Furniture: furniture.CreateBathroomSet()},
	}

	fmt.Println("\n\n\n\n", h)
}
