package furniture

type bed struct {
	Material string
	SeatsNum int
	Length   float32
	Width    float32
}

type bedsideTables struct {
	Material string
	Length   float32
	Width    float32
	Height   float32
}

type dresser struct {
	Material string
	Length   float32
	Width    float32
	Height   float32
	BoxCount int
}
type wardrobe struct {
	Material string
	Length   float32
	Width    float32
	Height   float32
	Volume   float32
}
