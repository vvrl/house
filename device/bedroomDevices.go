package device

type tv struct {
	Device         Device
	Type           string
	Diagonal       float32 // в дюймах
	Resolution     string
	VoiceAssistant bool
}

type laptop struct {
	Device        Device
	CPU           string
	GPU           string
	Diagonal      float32
	CapacityOfRAM int
}

type gameConsole struct {
	Device         Device
	CPU            string
	GPU            string
	DriveInDesign  bool
	MemoryCapacity float32
}

type BedroomDevices struct {
	TV          tv
	Laptop      laptop
	GameConsole gameConsole
}

func CreateBedroomDevices() BedroomDevices {

	return BedroomDevices{TV: tv{Device: Device{Name: "Телевизор", Brand: "Samsung", Model: "QE98QN100BUXRU", Power: 535, Length: 217.6, Width: 124.4, Height: 2, Colour: "Серый"}, Type: "LED", Diagonal: 98, Resolution: "4K 3840x2160", VoiceAssistant: true},
		Laptop:      laptop{Device: Device{Name: "Ноутбук", Brand: "MSI", Model: "Titan GT77 HX 13VI-096RU", Power: 330, Length: 33, Width: 39.7, Height: 2.3, Colour: "Черный"}, CPU: "Intel Core i9-13980HX", GPU: "GeForce RTX 4090 for laptops", Diagonal: 17.3, CapacityOfRAM: 32},
		GameConsole: gameConsole{Device: Device{Name: "Игровая консоль", Brand: "SONY", Model: "PlayStation 5", Power: 100, Length: 20.4, Width: 10.4, Height: 39, Colour: "Белый"}, CPU: "AMD Ryzen Zen 2", GPU: "AMD Radeon RDNA 2", DriveInDesign: true, MemoryCapacity: 825}}

}
