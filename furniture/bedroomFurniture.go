package furniture

// все размеры в метрах
type bed struct {
	Material string
	Type     string
	Length   float32
	Width    float32
	Colour   string
}

type bedsideTable struct {
	Material string
	Length   float32
	Width    float32
	Height   float32
	Colour   string
}

type dresser struct {
	Material string
	Length   float32
	Width    float32
	Height   float32
	BoxCount int
	Colour   string
}
type wardrobe struct {
	Material string
	Length   float32
	Width    float32
	Height   float32
	Mirror   bool
	Colour   string
}

type BedroomSet struct {
	Bed      bed
	Bedsides []bedsideTable
	Dresser  dresser
	Wardrobe wardrobe
}

func CreateBedroomSet() BedroomSet {

	var bedsideTables []bedsideTable

	for i := 0; i < 2; i++ {
		bedsideTables = append(bedsideTables, bedsideTable{Material: "Дерево",
			Width:  40,
			Length: 43,
			Height: 50,
			Colour: "Темно-серый",
		})
	}

	return BedroomSet{Bed: bed{Material: "Металл", Type: "двуспальная кровать", Length: 200, Width: 160, Colour: "Черный"},
		Bedsides: bedsideTables,
		Dresser:  dresser{Material: "Дерево", Length: 40, Width: 130, Height: 70, BoxCount: 6, Colour: "Белый"},
		Wardrobe: wardrobe{Material: "Дерево", Length: 180, Width: 60, Height: 220, Mirror: true, Colour: "Темно-серый"},
	}

}
