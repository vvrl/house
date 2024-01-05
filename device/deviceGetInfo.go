package device

import "fmt"

func (d Device) getGeneralDeviceOptions() {

	fmt.Printf("\vНазвание: \t%s\nБренд: \t%s\nМодель: \t%s\nМощность: \t%d Вт\nДлина: \t%.2f см\nШирина: \t%.2f см\nВысота: \t%.2f см\nЦвет: \t%s\n",
		d.Name, d.Brand, d.Model, d.Power, d.Length, d.Width, d.Height, d.Colour)

}

//                             Hall devices
//------------------------------------------------------------------------------

func (w wifiRouter) getInfo() {
	w.Device.getGeneralDeviceOptions()
	fmt.Printf("Максимальная скорость интернета: \t%d Мбит/с\nКоличество портов: \t%d\n", w.MaxInternetSpeed, w.NumberOfPorts)
}

func (r robotVacuumCleaner) getInfo() {
	r.Device.getGeneralDeviceOptions()

	wetClean := "Нет"
	if r.IfWetCleaning {
		wetClean = "Да"
	}

	fmt.Printf("Емкость аккумулятора: \t%d мА*ч\nВлажная уборка: \t%s\n", r.BatteryCapacity, wetClean)

}

func (i intercomSet) getInfo() {
	i.Device.getGeneralDeviceOptions()

	display := "Нет"
	if i.IfDisplay {
		display = "Да"
	}

	fmt.Printf("Наличие дисплея: \t%s", display)
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

	fmt.Printf("Морозилка: \t%s\nОбъем: \t%.2f л\nКоличество камер: \t%d\n", freezer, f.Volume, f.NumOfCameras)
}

func (d dishwasher) getInfo() {
	d.Device.getGeneralDeviceOptions()

	fmt.Printf("Количество программ: \t%d\nРасход воды: \t%.2f л\n", d.NumOfPrograms, d.WaterConsumption)
}

func (k electricKettle) getInfo() {
	k.Device.getGeneralDeviceOptions()

	light := "Нет"
	if k.Backlight {
		light = "Да"
	}
	fmt.Printf("Объем: \t%.2f л\nПодсветка: \t%s\n", k.Volume, light)
}

func (s stove) getInfo() {
	s.Device.getGeneralDeviceOptions()
	fmt.Printf("Тип: \t%s\nКоличество конфорок: \t%d\nОбъем духовки: \t%.2f л\n", s.Type, s.NumOfBurners, s.OvenCapacity)
}

func (m microwaveOven) getInfo() {
	m.Device.getGeneralDeviceOptions()

	grill := "Нет"
	if m.IfGrillMode {
		grill = "Да"
	}
	fmt.Printf("Тип управления: \t%s\nОбъем: \t%.2f л\nРежим гриля: \t%s\n", m.ControlType, m.Volume, grill)
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
	fmt.Printf("Тип: \t%s\nДиагональ: \t%.2f дюймов\nРазрешение экрана: \t%s\nГолосовое управление: \t%s\n", tv.Type, tv.Diagonal, tv.Resolution, voice)
}

func (l laptop) getInfo() {
	l.Device.getGeneralDeviceOptions()

	fmt.Printf("Процеcсор: \t%s\nВидеокарта: \t%s\nДиагональ: \t%.2f дюймов\nОбъем оперативной памяти: \t%d Гб\n", l.CPU, l.GPU, l.Diagonal, l.CapacityOfRAM)
}

func (g gameConsole) getInfo() {
	g.Device.getGeneralDeviceOptions()

	drive := "Нет"
	if g.DriveInDesign {
		drive = "Да"
	}

	fmt.Printf("Процеcсор: \t%s\nВидеокарта: \t%s\nДисковод: \t%s\nОбъем памяти: \t%.2f\n", g.CPU, g.GPU, drive, g.MemoryCapacity)
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
	fmt.Printf("Количество режимов: \t%d\n", h.NumberOfMods)
}

func (w washingMachine) getInfo() {
	w.Device.getGeneralDeviceOptions()
	fmt.Printf("Объем барабана: \t%d литров\nКоличество программ: \t%d\n", w.DrumVolume, w.NumOfPrograms)
}

func (d dryer) getInfo() {
	d.Device.getGeneralDeviceOptions()
	fmt.Printf("Объем барабана: \t%d литров\nМаксимальный вес загрузки: \t%d кг\n", d.DrumVolume, d.MaxLoad)
}

func (b BathroomDevices) BathroomDevicesInfo() {
	b.HairDryer.getInfo()
	b.WashingMachine.getInfo()
	b.Dryer.getInfo()
}
