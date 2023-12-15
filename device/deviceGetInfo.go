package device

import "fmt"

//                             Hall devices
//------------------------------------------------------------------------------

func (w wifiRouter) getInfo() {
	w.Device.getGeneralDeviceOptions()
	fmt.Printf("Максимальная скорость интернета: %20d Мбит/с\nКоличество портов: %20d\n", w.MaxInternetSpeed, w.NumberOfPorts)
}

func (r robotVacuumCleaner) getInfo() {
	r.Device.getGeneralDeviceOptions()

	wetClean := "Нет"
	if r.IfWetCleaning {
		wetClean = "Да"
	}

	fmt.Printf("Емкость аккумулятора: %20d мА*ч\nВлажная уборка: %20s\n", r.BatteryCapacity, wetClean)

}

func (i intercomSet) getInfo() {
	i.Device.getGeneralDeviceOptions()

	display := "Нет"
	if i.IfDisplay {
		display = "Да"
	}

	fmt.Println("Наличие дисплея: " + display)
}

func (h HallDevices) HallDevicesInfo() {
	h.WifiRouter.getInfo()
	h.RobotVacuumCleaner.getInfo()
	h.IntercomSet.getInfo()
}

//---------------------------------------------------------------------------
//
//
//
//                              Kitchen devices
//---------------------------------------------------------------------------

func (f fridge) getInfo() {
	f.Device.getGeneralDeviceOptions()

	freezer := "Нет"
	if f.IfFreezer {
		freezer = "Да"
	}

	fmt.Printf("Морозилка: %20s\nОбъем: %20f л\nКоличество камер: %20d\n", freezer, f.Volume, f.NumOfCameras)
}

func (d dishwasher) getInfo() {
	d.Device.getGeneralDeviceOptions()

	fmt.Printf("Количество программ: %20d\nРасход воды: %20f л\n", d.NumOfPrograms, d.WaterConsumption)
}

func (k electricKettle) getInfo() {
	k.Device.getGeneralDeviceOptions()

	light := "Нет"
	if k.Backlight {
		light = "Да"
	}
	fmt.Printf("Объем: %20f л\nПодсветка: %20s\n", k.Volume, light)
}

func (s stove) getInfo() {
	s.Device.getGeneralDeviceOptions()
	fmt.Printf("Тип: %20s\nКоличество конфорок: %20d\nОбъем духовки: %20f л\n", s.Type, s.NumOfBurners, s.OvenCapacity)
}

func (m microwaveOven) getInfo() {
	m.Device.getGeneralDeviceOptions()

	grill := "Нет"
	if m.IfGrillMode {
		grill = "Да"
	}
	fmt.Printf("Тип управления: %20s\nОбъем: %20f л\nРежим гриля: %20s\n", m.ControlType, m.Volume, grill)
}

func (k KitchenDevices) KitchenDevicesInfo() {
	k.Fridge.getInfo()
	k.Dishwasher.getInfo()
	k.Kettle.getInfo()
	k.Stove.getInfo()
	k.Microwave.getInfo()
}

//-------------------------------------------------------------------------------
//
//
//
//                             Bedroom devices
//-------------------------------------------------------------------------------

func (tv tv) getInfo() {
	tv.Device.getGeneralDeviceOptions()

	voice := "Нет"
	if tv.VoiceAssistant {
		voice = "Да"
	}
	fmt.Printf("Тип: %20s\nДиагональ: %20f дюймов\nРазрешение экрана: %20s\nГолосовое управление: %20s\n", tv.Type, tv.Diagonal, tv.Resolution, voice)
}

func (l laptop) getInfo() {
	l.Device.getGeneralDeviceOptions()

	fmt.Printf("Процеcсор: %20s\nВидеокарта: %20s\nДиагональ: %20f дюймов\nОбъем оперативной памяти: %20d Гб\n", l.CPU, l.GPU, l.Diagonal, l.CapacityOfRAM)
}

func (g gameConsole) getInfo() {
	g.Device.getGeneralDeviceOptions()

	drive := "Нет"
	if g.DriveInDesign {
		drive = "Да"
	}

	fmt.Printf("Процеcсор: %20s\nВидеокарта: %20s\nДисковод: %20s\nОбъем памяти: %20f\n", g.CPU, g.GPU, drive, g.MemoryCapacity)
}

func (b BedroomDevices) BedroomDevicesInfo() {
	b.TV.getInfo()
	b.Laptop.getInfo()
	b.GameConsole.getInfo()
}

//-------------------------------------------------------------------------------------
//
//
//
//                                  Bathroom devices
//-------------------------------------------------------------------------------------

func (h hairDryer) getInfo() {
	h.Device.getGeneralDeviceOptions()
	fmt.Printf("Количество режимов: %20d\n", h.NumberOfMods)
}

func (w washingMachine) getInfo() {
	w.Device.getGeneralDeviceOptions()
	fmt.Printf("Объем барабана: %20d\nКоличество программ: %20d\n", w.DrumVolume, w.NumOfPrograms)
}

func (d dryer) getInfo() {
	d.Device.getGeneralDeviceOptions()
	fmt.Printf("Объем барабана: %20d\nМаксимальный вес загрузки: %20d\n", d.DrumVolume, d.MaxLoad)
}

func (b BathroomDevices) BathroomDevicesInfo() {
	b.HairDryer.getInfo()
	b.WashingMachine.getInfo()
	b.Dryer.getInfo()
}
