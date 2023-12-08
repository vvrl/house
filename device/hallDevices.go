package device

type wifiRouter struct {
	Device           Device
	MaxInternetSpeed int // mbit/sec
	NumberOfPorts    int
}

type robotVacuumCleaner struct {
	Device          Device
	BatteryCapacity int
	IfWetCleaning   bool
}

type intercomSet struct {
	Device    Device
	IfDisplay bool
}

type HallDevices struct {
	WifiRouter         wifiRouter
	RobotVacuumCleaner robotVacuumCleaner
	IntercomSet        intercomSet
}

func CreateHallDevices() HallDevices {

	return HallDevices{WifiRouter: wifiRouter{Device: Device{Name: "Wi-fi роутер", Brand: "ASUS ROG", Model: "Rapture GT6", Power: 19, Length: 42, Width: 14, Height: 28, Colour: "Черный"}, MaxInternetSpeed: 4804, NumberOfPorts: 7},
		RobotVacuumCleaner: robotVacuumCleaner{Device: Device{Name: "Робот-пылесос", Brand: "DreameBot", Model: "L30 Ultra", Power: 75, Length: 35, Width: 35, Height: 103.8, Colour: "Черный"}, BatteryCapacity: 6400, IfWetCleaning: true},
		IntercomSet:        intercomSet{Device: Device{Name: "Домофон", Brand: "Intercom kit Hikvision", Model: "DS-KIS603-P(C)", Power: 20, Length: 18, Width: 6, Height: 12, Colour: "Серебряный"}, IfDisplay: true}}
}
