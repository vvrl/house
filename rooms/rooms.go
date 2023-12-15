package rooms

import (
	"goProjects/house/device"
	"goProjects/house/furniture"
)

type kitchen struct {
	Room      Room
	Furniture furniture.KitchenSet
	Devices   device.KitchenDevices
}

func (k kitchen) KitchenGetInfo() {
	k.Devices.KitchenDevicesInfo()
}

type bedroom struct {
	Room      Room
	Furniture furniture.BedroomSet
	Device    device.BedroomDevices
}

type hall struct {
	Room      Room
	Furniture furniture.HallSet
	Devices   device.HallDevices
}

func (h hall) HallGetInfo() {
	h.Devices.HallDevicesInfo()
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

func (r Rooms) RoomsInfo() {
	r.Hall.HallGetInfo()
}

func CreateRooms() Rooms {

	return Rooms{Kitchen: kitchen{Room: CreateRoom("Кухня"), Furniture: furniture.CreateKitchenSet(), Devices: device.CreateKitchenDevices()},
		Bedroom:  bedroom{Room: CreateRoom("Спальня"), Furniture: furniture.CreateBedroomSet(), Device: device.CreateBedroomDevices()},
		Hall:     hall{Room: CreateRoom("Прихожая"), Furniture: furniture.CreateHallSet(), Devices: device.CreateHallDevices()},
		Bathroom: bath{Room: CreateRoom("Ванная комната"), Furniture: furniture.CreateBathroomSet(), Devices: device.CreateBathroomDevices()}}

}
