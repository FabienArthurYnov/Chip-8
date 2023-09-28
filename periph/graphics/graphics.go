package graphics

import (
	"chip-8/structs"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

func UpdateGraph(chip8 structs.Chip8) {
	chip8.Display.Clear(color.Black)

	for i, _ := range chip8.ScreenTable {
		for j, _ := range chip8.ScreenTable[i] {
			if chip8.ScreenTable[i][j] { // si pixel true (white)
				imd := imdraw.New(nil)
				imd.Color = color.White
				imd.Push(pixel.V(float64(i*10), float64(j*10)))
				imd.Push(pixel.V(float64(i*10+10), float64(j*10)))
				imd.Push(pixel.V(float64(i*10+10), float64(j*10+10)))
				imd.Push(pixel.V(float64(i*10), float64(j*10+10)))
				imd.Polygon(0)
				imd.Draw(chip8.Display)
			} else { // si pixel true (black)
				imd := imdraw.New(nil)
				imd.Color = color.Black
				imd.Push(pixel.V(float64(i*10), float64(j*10)))
				imd.Push(pixel.V(float64(i*10+10), float64(j*10)))
				imd.Push(pixel.V(float64(i*10+10), float64(j*10+10)))
				imd.Push(pixel.V(float64(i*10), float64(j*10+10)))
				imd.Polygon(0)
				imd.Draw(chip8.Display)
			}
		}
	}

	chip8.Display.Update()
}
