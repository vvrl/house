package furniture

import "fmt"

// все размеры в сантиметрах
type table struct {
	Height   float32
	Length   float32
	Width    float32
	Material string
	Shape    string
}

type chair struct {
	Material string
	Height   float32
}

type wallCabinet struct {
	Material string
	Length   float32
	Width    float32
	Height   float32
}

type floorCabinet struct {
	Material         string
	Length           float32
	Width            float32
	Height           float32
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
		chairsSlice = append(chairsSlice, chair{Material: "дерево", Height: 50})
		wallCabinetSlice = append(wallCabinetSlice, wallCabinet{Material: "дерево", Length: 50, Width: 25, Height: 100})
		floorCabinetSlice = append(floorCabinetSlice, floorCabinet{Material: "дерево", Length: 50, Width: 25, Height: 100})
	}

	for i := 0; i < 4; i++ {
		fmt.Println(chairsSlice[i])
	}
	Table := table{Height: 100, Length: 210, Width: 75, Material: "дерево", Shape: "прямоугольник"}
	kSet := KitchenSet{Chairs: chairsSlice, FloorCabinets: floorCabinetSlice, Table: Table, WallCabinets: wallCabinetSlice}

	return kSet
}
