package rooms

import "goProjects/house/furniture"

type Bedroom struct {
	Room      Room
	Furniture furniture.BedroomSet
}
