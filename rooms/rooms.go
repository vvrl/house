package rooms

import (
	"fmt"
	"goProjects/house/device"
	"goProjects/house/furniture"
)

type kitchen struct {
	Room      Room
	Furniture furniture.KitchenSet
	Devices   device.KitchenDevices
}

type bedroom struct {
	Room      Room
	Furniture furniture.BedroomSet
	Devices   device.BedroomDevices
}

type hall struct {
	Room      Room
	Furniture furniture.HallSet
	Devices   device.HallDevices
}

type bath struct {
	Room      Room
	Furniture furniture.BathroomSet
	Devices   device.BathroomDevices
}

type Rooms struct {
	Kitchen  kitchen
	Bedroom  bedroom
	Hall     hall
	Bathroom bath
}

func CreateRooms() Rooms {

	return Rooms{Kitchen: kitchen{Room: CreateRoom("КУХНЯ"), Furniture: furniture.CreateKitchenSet(), Devices: device.CreateKitchenDevices()},
		Bedroom:  bedroom{Room: CreateRoom("СПАЛЬНЯ"), Furniture: furniture.CreateBedroomSet(), Devices: device.CreateBedroomDevices()},
		Hall:     hall{Room: CreateRoom("ПРИХОЖАЯ"), Furniture: furniture.CreateHallSet(), Devices: device.CreateHallDevices()},
		Bathroom: bath{Room: CreateRoom("ВАННАЯ КОМНАТА"), Furniture: furniture.CreateBathroomSet(), Devices: device.CreateBathroomDevices()}}

}

func PrintLine() {
	fmt.Println("------------------------------------------------")
}
