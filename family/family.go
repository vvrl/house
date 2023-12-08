package family

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

	//fam = append(fam, familyMember{Name: "Артем Жиренко", Age: 20, Gender: true, FamilyStatus: "Холост"})
	//fam = append(fam, familyMember{Name: "Райан Гослинг", Age: 43, Gender: true, FamilyStatus: "Женат"})
	//fam = append(fam, familyMember{Name: "Роман Самохин", Age: 20, Gender: true, FamilyStatus: "Холост"})
	//fam = append(fam, familyMember{Name: "Меган Фокс", Age: 37, Gender: false, FamilyStatus: "Замужем"})

	return Family{Family: fam}
}
