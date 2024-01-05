package rooms

import (
	"fmt"
	"math/rand"
)

type Room struct {
	Name   string
	Length float32 //метры
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

	return rm
}

func (r Room) getRoomOptions() {
	fmt.Println()
	PrintLine()
	fmt.Printf("\t\t%s\n", r.Name)
	PrintLine()
	fmt.Printf("\vДлина:  \t%.2f м\nШирина: \t%.2f м\nВысота: \t%.2f м\nПлощадь: \t%.2f м^2\n", r.Length, r.Width, r.Height, r.Square)
}
