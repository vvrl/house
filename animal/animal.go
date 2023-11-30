package animal

type animal struct {
	Type   string
	Name   string
	Age    int
	Colour string
	Gender bool
}

type Animals struct {
	Animals []animal
}

func CreateAnimal() Animals {
	var animals []animal

	animals = append(animals, animal{Type: "Кот", Name: "Лох", Age: 5, Colour: "Рыжый", Gender: false})

	return Animals{Animals: animals}
}
