package family

import "fmt"

type familyMember struct {
	Name         string
	Age          int
	Gender       bool
	FamilyStatus string
}

type Family struct {
	Family []familyMember
}

func CreateFamily() Family {

	fam := []familyMember{{Name: "Артем Жиренко", Age: 20, Gender: true, FamilyStatus: "Холост"},
		{Name: "Райан Гослинг", Age: 43, Gender: true, FamilyStatus: "Женат"},
		{Name: "Меган Фокс", Age: 37, Gender: false, FamilyStatus: "Замужем"}}

	return Family{Family: fam}
}

func (f Family) FamilyGetInfo() {
	fmt.Print("\t\tЖИЛЬЦЫ И ЖИВОТНЫЕ")

	for i := 0; i < len(f.Family); i++ {
		gender := "Женский"
		if f.Family[i].Gender {
			gender = "Мужской"
		}

		fmt.Println("\nЖитель ", i+1)
		fmt.Printf("Имя: \t%s\nВозраст: \t%d\nПол: \t%s\nСеймейное положение: \t%s\n", f.Family[i].Name, f.Family[i].Age, gender, f.Family[i].FamilyStatus)
	}
}
