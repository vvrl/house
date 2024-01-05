package main

import (
	"fmt"
	"goProjects/house/house"
)

func main() {
	fmt.Println("123")
	h := house.CreateHouse("ул.Кошкина д.2")
	h.GetInfo()
}
