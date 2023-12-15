package rooms

import (
	"fmt"
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

	//fmt.Println(rm)
	return rm
}

func (r Room) getRoomOptions() {
	fmt.Printf("%")
}
