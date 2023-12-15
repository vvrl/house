package main

import (
	"fmt"
	"goProjects/house/animal"
	"goProjects/house/family"
	"goProjects/house/rooms"
)

type House struct {
	SumSquare float32
	Rooms     rooms.Rooms
	Family    family.Family
	Animals   animal.Animals
}

func main() {
	h := House{Rooms: rooms.CreateRooms(),
		Family:  family.CreateFamily(),
		Animals: animal.CreateAnimal(),
	}

	h.getInfo()
	fmt.Println("\n\n\n\n", h)
}

func (h House) getInfo() {
	h.Rooms.RoomsInfo()
}
