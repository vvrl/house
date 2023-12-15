package furniture

// все размеры в сантиметрах
type table struct {
	Furniture Furniture
	Shape     string
}

type chair struct {
	Furniture Furniture
}

type wallCabinet struct {
	Furniture Furniture
	DoorCount int
}

type floorCabinet struct {
	Furniture        Furniture
	TableTopMaterial string
}

type KitchenSet struct {
	WallCabinets  []wallCabinet
	FloorCabinets []floorCabinet
	Table         table
	Chairs        []chair
}

func CreateKitchenSet() KitchenSet {

	var chairsSlice []chair
	var wallCabinetSlice []wallCabinet
	var floorCabinetSlice []floorCabinet

	for i := 0; i < 4; i++ {
		chairsSlice = append(chairsSlice, chair{Furniture: Furniture{Name: "Стул", Material: "Дерево", Length: 49, Width: 45, Height: 50, Colour: "Белый"}})
		wallCabinetSlice = append(wallCabinetSlice, wallCabinet{Furniture: Furniture{Name: "Подвесной шкаф", Material: "Дерево", Length: 50, Width: 25, Height: 100, Colour: "Серый"}, DoorCount: 2})
		floorCabinetSlice = append(floorCabinetSlice, floorCabinet{Furniture: Furniture{Name: "Кухонный стол", Material: "Дерево", Length: 50, Width: 25, Height: 100, Colour: "Серый"}, TableTopMaterial: "Мрамор"})
	}

	Table := table{Furniture: Furniture{Height: 100, Length: 210, Width: 75, Material: "Стекло", Colour: "Белый"}, Shape: "Прямоугольник"}
	kSet := KitchenSet{Chairs: chairsSlice, FloorCabinets: floorCabinetSlice, Table: Table, WallCabinets: wallCabinetSlice}

	return kSet
}
