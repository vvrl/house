package main

import "house/home"

func main() {
	h := home.CreateHouse("ул.Кошкина д.2")
	h.GetInfo()
}
