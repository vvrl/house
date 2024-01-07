package animal

import (
	"fmt"
)

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

	animals := []animal{{Type: "Кот", Name: "Лох", Age: 5, Colour: "Рыжый", Gender: false},
		{Type: "Хомяк", Name: "Хрустик", Age: 1, Colour: "Черный", Gender: true}}

	return Animals{Animals: animals}
}

func (a Animals) AnimalsGetInfo() {

	for i := 0; i < len(a.Animals); i++ {
		gender := "Женский"
		if a.Animals[i].Gender {
			gender = "Мужской"
		}
		fmt.Println("\nЖивотное ", i+1)
		fmt.Printf("Вид: \t%s\nИмя: \t%s\nВозраст:\t%d\nЦвет: \t%s\nПол: \t%s\n", a.Animals[i].Type, a.Animals[i].Name, a.Animals[i].Age, a.Animals[i].Colour, gender)

	}
}
