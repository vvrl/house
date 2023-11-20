package main

import (
	"fmt"
	"goProjects/house/rooms"
)

type House struct {
	kitchen rooms.Kitchen
}

func main() {
	h := House{}
	rooms.Set(h.kitchen)

	//h.kitchen.Furniture = h.CreateKitchenSet()
	fmt.Println(h)
}
