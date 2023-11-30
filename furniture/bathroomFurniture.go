package furniture

type bathtub struct {
	Furniture Furniture
}

type sinkCabinet struct {
	Furniture    Furniture
	MaterialSink string
	SinkDepth    float32
	BoxCount     int
}

type mirrorCabinet struct {
	Furniture Furniture
	ifMirror  bool
}

type toilet struct {
	Furniture Furniture
}

type BathroomSet struct {
	Bathtub       bathtub
	SinkCabinet   sinkCabinet
	MirrorCabinet mirrorCabinet
	Toilet        toilet
}

func CreateBathroomSet() BathroomSet {

	return BathroomSet{Bathtub: bathtub{Furniture: Furniture{Name: "Ванна", Length: 170, Width: 70, Height: 65, Material: "Чугун", Colour: "Белый"}},
		SinkCabinet:   sinkCabinet{Furniture: Furniture{Name: "Шкаф с раковиной", Length: 83, Width: 49, Height: 64, Material: "Дерево", Colour: "Светло-серый"}, MaterialSink: "Керамика", SinkDepth: 20, BoxCount: 3},
		MirrorCabinet: mirrorCabinet{Furniture: Furniture{Name: "Шкафчик с зеркалом", Length: 40, Width: 20, Height: 70, Material: "Дерево", Colour: "Белый"}, ifMirror: true},
		Toilet:        toilet{Furniture: Furniture{Name: "Skibidi Toilet", Length: 74, Material: "Керамика", Colour: "Белый"}}}

}
