package main

import (
	"goProjects/house/house"
)

func main() {
	h := house.CreateHouse("ул.Кошкина д.2")
	h.GetInfo()
}
