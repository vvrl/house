package device

type hairDryer struct {
	Device       Device
	NumberOfMods int
}

type washingMachine struct {
	Device        Device
	DrumVolume    int // объем барабана в литрах
	NumOfPrograms int
}

type dryer struct {
	Device     Device
	DrumVolume int
	MaxLoad    int //максимальный вес загрузки кг
}

type BathroomDevices struct {
	WashingMachine washingMachine
	Dryer          dryer
	HairDryer      hairDryer
}

func CreateBathroomDevices() BathroomDevices {

	return BathroomDevices{WashingMachine: washingMachine{Device: Device{Name: "Стиральная машина", Brand: "Bosch", Model: "WAV28MX0ME", Power: 500, Length: 60, Width: 59, Height: 85, Colour: "Серебряный"}, DrumVolume: 45, NumOfPrograms: 14},
		Dryer:     dryer{Device: Device{Name: "Сушильная машина", Brand: "Asko", Model: "T208H.W", Power: 700, Length: 59.5, Width: 65.4, Height: 85, Colour: "Серебряный"}, DrumVolume: 40, MaxLoad: 8},
		HairDryer: hairDryer{Device: Device{Name: "Фен", Brand: "Dyson", Model: "Supersonic HD08", Power: 1600, Length: 10, Width: 7.8, Height: 24.5, Colour: "Розовый - серебристый"}, NumberOfMods: 4}}

}
