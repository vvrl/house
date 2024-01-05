package furniture

import "fmt"

func (f Furniture) getGeneralFurnitureOptions() {

	fmt.Printf("\vНазвание: \t%s\nМатериал: \t%s\nЦвет: \t%s\nДлина: \t%.2f см\nШирина: \t%.2f см\nВысота: \t%.2f см\n", f.Name, f.Material, f.Colour, f.Length, f.Width, f.Height)

}

//                              Hall furniture
//------------------------------------------------------------------------------

func (p pouf) getInfo() {
	p.Furniture.getGeneralFurnitureOptions()
	storage := "Нет"
	if p.ifStorage {
		storage = "Да"
	}
	fmt.Printf("Наличие ящика: \t%s\n", storage)
}

func (m mirror) getInfo() {
	m.Furniture.getGeneralFurnitureOptions()
	mir := "Нет"
	if m.ifMirror {
		mir = "Да"
	}
	fmt.Printf("Наличие зеркала: \t%s\n", mir)
}

func (s shoeCabinet) getInfo() {
	s.Furniture.getGeneralFurnitureOptions()

	fmt.Printf("Количество ящиков: \t%d\n", s.BoxCount)
}

func (t hallTree) getInfo() {
	t.Furniture.getGeneralFurnitureOptions()

	fmt.Printf("Количество крючков: \t%d\n", t.HooksCount)
}

func (h HallSet) HallFurnitureInfo() {
	h.Mirror.getInfo()
	h.ShoeCabinet.getInfo()
	h.HallTree.getInfo()

	fmt.Printf("Количество пуфиков: \t%d\n", len(h.Poufs))
	h.Poufs[0].getInfo()
}

//--------------------------------------------------------------------------------
//
//
//
//                              Kitchen furniture
//--------------------------------------------------------------------------------

func (t table) getInfo() {
	t.Furniture.getGeneralFurnitureOptions()

	fmt.Printf("Форма: \t%s\n", t.Shape)
}

func (c chair) getInfo() {
	c.Furniture.getGeneralFurnitureOptions()
}

func (w wallCabinet) getInfo() {
	w.Furniture.getGeneralFurnitureOptions()

	fmt.Printf("Количество дверей: \t%d\n", w.DoorCount)
}

func (f floorCabinet) getInfo() {
	f.Furniture.getGeneralFurnitureOptions()

	fmt.Printf("Материал столешницы: \t%s\n", f.TableTopMaterial)
}

func (k KitchenSet) KitchenFurnitureInfo() {
	fmt.Printf("Количество подвесных шкафов: \t%d\n", len(k.WallCabinets))
	k.WallCabinets[0].getInfo()

	fmt.Printf("Количество напольных шкафов: \t%d\n", len(k.FloorCabinets))
	k.FloorCabinets[0].getInfo()

	k.Table.getInfo()

	fmt.Printf("Количество стульев: \t%d\n", len(k.Chairs))
	k.Chairs[0].getInfo()
}

//------------------------------------------------------------------------------------
//
//
//
//                            Bedroom furniture
//------------------------------------------------------------------------------------

func (b bed) getInfo() {
	b.Furniture.getGeneralFurnitureOptions()

	fmt.Printf("Тип кровати: \t%s\n", b.Type)
}

func (bd bedsideTable) getInfo() {
	bd.Furniture.getGeneralFurnitureOptions()
}

func (d dresser) getInfo() {
	d.Furniture.getGeneralFurnitureOptions()

	fmt.Printf("Количество ящиков: \t%d\n", d.BoxCount)
}

func (w wardrobe) getInfo() {
	w.Furniture.getGeneralFurnitureOptions()

	mir := "Нет"
	if w.Mirror {
		mir = "Да"
	}
	fmt.Printf("Наличие зеркала: \t%s\n", mir)
}

func (b BedroomSet) BedroomFurnitureInfo() {
	b.Bed.getInfo()

	fmt.Printf("Количество тумбочек: \t%d\n", len(b.Bedsides))
	b.Bedsides[0].getInfo()

	b.Dresser.getInfo()
	b.Wardrobe.getInfo()
}

//------------------------------------------------------------------------------
//
//
//
//                               Bathroom furniture
//-------------------------------------------------------------------------------

func (b bathtub) getInfo() {
	b.Furniture.getGeneralFurnitureOptions()
}

func (s sinkCabinet) getInfo() {
	s.Furniture.getGeneralFurnitureOptions()

	fmt.Printf("Материал раковины: \t%s\nГлубина раковины: \t%.2f см\nКоличество ящиков: \t%d\n", s.MaterialSink, s.SinkDepth, s.BoxCount)
}

func (m mirrorCabinet) getInfo() {
	m.Furniture.getGeneralFurnitureOptions()

	mir := "Нет"
	if m.ifMirror {
		mir = "Да"
	}
	fmt.Printf("Наличие зеркала: \t%s\n", mir)
}

func (t toilet) getInfo() {
	t.Furniture.getGeneralFurnitureOptions()
}

func (b BathroomSet) BathroomFurnitureInfo() {
	b.Bathtub.getInfo()
	b.SinkCabinet.getInfo()
	b.MirrorCabinet.getInfo()
	b.Toilet.getInfo()
}

//---------------------------------------------------------------------------------
