package device

import "fmt"

type Device struct {
	Name   string
	Brand  string
	Model  string
	Power  int // мощность в ваттах
	Length float32
	Width  float32
	Height float32 // размеры в сантиметрах
	Colour string
}

func (d Device) getGeneralDeviceOptions() {

	fmt.Printf("Название: %20s\nБренд: %20s\nМодель: %20s\nМощность: %20d Вт\nДлина: %20f см\nШирина: %20f см\nВысота: %20f см\nЦвет: %20s\n",
		d.Name, d.Brand, d.Model, d.Power, d.Length, d.Width, d.Height, d.Colour)

}
