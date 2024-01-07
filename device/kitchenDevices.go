package device

type fridge struct {
	Device       Device
	IfFreezer    bool
	Volume       float32
	NumOfCameras int
}

type dishwasher struct {
	Device           Device
	NumOfPrograms    int
	WaterConsumption float32
}

type electricKettle struct {
	Device    Device
	Volume    float32
	Backlight bool
}

type stove struct {
	Device       Device
	Type         string
	NumOfBurners int
	OvenCapacity float32
}

type microwaveOven struct {
	Device      Device
	ControlType string
	Volume      float32
	IfGrillMode bool
}

type KitchenDevices struct {
	Fridge     fridge
	Dishwasher dishwasher
	Kettle     electricKettle
	Stove      stove
	Microwave  microwaveOven
}

func CreateKitchenDevices() KitchenDevices {

	return KitchenDevices{Fridge: fridge{Device: Device{Name: "Холодильник", Brand: "Hitachi", Model: "R-WX 630 KU XW", Power: 400, Length: 75, Width: 73.8, Height: 183.3, Colour: "Белый"}, IfFreezer: true, Volume: 473, NumOfCameras: 5},
		Dishwasher: dishwasher{Device: Device{Name: "Посудомоечная машина", Brand: "Bomann", Model: "GSP 7410", Power: 2100, Length: 59.8, Width: 60, Height: 84.5, Colour: "Белый"}, NumOfPrograms: 6, WaterConsumption: 11},
		Kettle:     electricKettle{Device: Device{Name: "Чайник", Brand: "KitchenAid", Model: "5KEK1322ESS", Power: 2000, Length: 23.9, Width: 21.1, Height: 27, Colour: "Серебристый"}, Volume: 1.5, Backlight: true},
		Stove:      stove{Device: Device{Name: "Плита", Brand: "Gorenje", Model: "GECS5B70CLI", Power: 9200, Length: 50, Width: 59.4, Height: 85, Colour: "Бежевый"}, Type: "Электрическая плита", NumOfBurners: 4, OvenCapacity: 70},
		Microwave:  microwaveOven{Device: Device{Name: "Микроволновая печь", Brand: "Samsung", Model: "MC35R8088LC/BW", Power: 2950, Length: 48.1, Width: 52.8, Height: 40.4, Colour: "Серый"}, ControlType: "Сенсор", Volume: 35, IfGrillMode: true}}

}
