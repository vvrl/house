package home

import (
	"fmt"
	"house/animal"
	"house/family"
	"house/rooms"
	"reflect"
)

type House struct {
	Address   string
	SumSquare float32
	Rooms     rooms.Rooms
	Family    family.Family
	Animals   animal.Animals
}

func (h House) GetInfo() {
	NumRooms := reflect.ValueOf(h.Rooms).NumField()
	fmt.Println()
	rooms.PrintLine()

	fmt.Printf("\t\tОБЩАЯ ИНФОРМАЦИЯ\n")
	rooms.PrintLine()

	fmt.Printf("Адрес: \t%s\nСуммарная площадь: \t%.2f м^2\nКоличество комнат: \t%d\nКоличество жильцов: \t%d\nКоличество животных\t%d\n",
		h.Address, h.SumSquare, NumRooms, len(h.Family.Family), len(h.Animals.Animals))

	rooms.PrintLine()

	h.Family.FamilyGetInfo()
	h.Animals.AnimalsGetInfo()
	h.Rooms.RoomsGetInfo()
}

func (h *House) calcSquare() {
	h.SumSquare = h.Rooms.Kitchen.Room.Square + h.Rooms.Bedroom.Room.Square +
		h.Rooms.Hall.Room.Square + h.Rooms.Bathroom.Room.Square
}

func CreateHouse(address string) House {

	h := House{Rooms: rooms.CreateRooms(),
		Family:  family.CreateFamily(),
		Animals: animal.CreateAnimal(),
		Address: address,
	}
	h.calcSquare()
	return h
}
