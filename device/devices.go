package device

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
