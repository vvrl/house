package rooms

import (
	"fmt"
	"goProjects/house/furniture"
	"math/rand"
)

type Room struct {
	Name   string
	Length float32
	Width  float32
	Height float32
	Square float32
}

func (r *Room) calculateSquare() {
	r.Square = r.Length * r.Width
}

func CreateRoom(name string) Room {

	rm := Room{Name: name,
		Length: rand.Float32()*(10-2) + 2,
		Width:  rand.Float32()*(7-2) + 2,
		Height: 2.8,
	}
	rm.calculateSquare()

	fmt.Println(rm)
	return rm
}

type Kitchen struct {
	Room      Room
	Furniture furniture.KitchenSet
}

type Bedroom struct {
	Room      Room
	Furniture furniture.BedroomSet
}

type Hall struct {
	Room      Room
	Furniture furniture.HallSet
}

type Bath struct {
	Room      Room
	Furniture furniture.BathroomSet
}
