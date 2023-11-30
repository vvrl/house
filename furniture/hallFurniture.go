package furniture

type pouf struct {
	Furniture Furniture
	ifStorage bool
}

type mirror struct {
	Furniture Furniture
	ifMirror  bool
}

type shoeCabinet struct {
	Furniture Furniture
	BoxCount  int
}

type hallTree struct {
	Furniture  Furniture
	HooksCount int
}

type HallSet struct {
	Poufs       []pouf
	Mirror      mirror
	ShoeCabinet shoeCabinet
	HallTree    hallTree
}

func CreateHallSet() HallSet {

	var poufs []pouf

	for i := 0; i < 2; i++ {
		poufs = append(poufs, pouf{Furniture: Furniture{Name: "Пуфик", Length: 50, Width: 50, Height: 40, Material: "Ротанг", Colour: "Желтый"}, ifStorage: true})
	}

	Mirror := mirror{Furniture: Furniture{Name: "Зеркало", Length: 75, Height: 168}, ifMirror: true}
	ShoeCabinet := shoeCabinet{Furniture: Furniture{Name: "Шкаф для обуви", Length: 89, Width: 30, Height: 127, Material: "Дерево", Colour: "Белый"}, BoxCount: 2}
	HallTree := hallTree{Furniture: Furniture{Name: "Вешалка", Length: 100, Width: 42, Height: 182, Material: "Металл", Colour: "Черный"}, HooksCount: 12}

	return HallSet{Poufs: poufs, Mirror: Mirror, ShoeCabinet: ShoeCabinet, HallTree: HallTree}
}
