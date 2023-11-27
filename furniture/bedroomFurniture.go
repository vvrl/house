package furniture

// все размеры в сантиметрах
type bed struct {
	Furniture Furniture
	Type      string
}

type bedsideTable struct {
	Furniture Furniture
}

type dresser struct {
	Furniture Furniture
	BoxCount  int
}
type wardrobe struct {
	Furniture Furniture
	Mirror    bool
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
		bedsideTables = append(bedsideTables, bedsideTable{Furniture: Furniture{Name: "Тумбочка", Material: "Дерево",
			Width:  40,
			Length: 43,
			Height: 50,
			Colour: "Темно-серый"},
		})
	}

	return BedroomSet{Bed: bed{Furniture: Furniture{Name: "Кровать", Material: "Металл", Height: 60, Length: 200, Width: 160, Colour: "Черный"}, Type: "Двуспальная"},
		Bedsides: bedsideTables,
		Dresser:  dresser{Furniture: Furniture{Name: "Комод", Material: "Дерево", Length: 40, Width: 130, Height: 70, Colour: "Белый"}, BoxCount: 6},
		Wardrobe: wardrobe{Furniture: Furniture{Name: "Гардероб", Material: "Дерево", Length: 180, Width: 60, Height: 220, Colour: "Темно-серый"}, Mirror: true},
	}

}
