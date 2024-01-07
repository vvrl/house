package rooms

func (k kitchen) KitchenGetInfo() {
	k.Room.getRoomOptions()
	k.Furniture.KitchenFurnitureInfo()
	k.Devices.KitchenDevicesInfo()
}

func (h hall) HallGetInfo() {
	h.Room.getRoomOptions()
	h.Furniture.HallFurnitureInfo()
	h.Devices.HallDevicesInfo()
}

func (b bedroom) BedroomGetInfo() {
	b.Room.getRoomOptions()
	b.Furniture.BedroomFurnitureInfo()
	b.Devices.BedroomDevicesInfo()
}

func (b bath) BathroomGetInfo() {
	b.Room.getRoomOptions()
	b.Furniture.BathroomFurnitureInfo()
	b.Devices.BathroomDevicesInfo()
}

func (r Rooms) RoomsGetInfo() {
	r.Hall.HallGetInfo()
	r.Kitchen.KitchenGetInfo()
	r.Bedroom.BedroomGetInfo()
	r.Bathroom.BathroomGetInfo()
}
