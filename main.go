package main

import (
	"chip-8/periph/graphics"
	"chip-8/structs"
	"chip-8/utility"
	"fmt"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	cfgWindow := pixelgl.WindowConfig{
		Title:  "Chip-8 Emulator",
		Bounds: pixel.R(0, 0, 640, 320),
		VSync:  true,
	}
	window, err := pixelgl.NewWindow(cfgWindow)
	if err != nil {
		panic(err)
	}

	chip8 := structs.Chip8{}
	timeStart := time.Now()

	chip8.Display = window

	var fileName string = utility.InputFileName()
	chip8.Load(fileName) // load the game into the memory
	fmt.Println(chip8.Memory)

	for !window.Closed() {
		chip8.EmulateOneCycle() // emulate one cycle

		if time.Now().Sub(timeStart) > time.Second { // when one second has past
			if chip8.DelayTimer > 0 {
				chip8.DelayTimer -= 1
			}
			if chip8.SoundTimer > 0 {
				chip8.SoundTimer -= 1
			}
			timeStart = time.Now()
		}

		if chip8.DrawFlag { //if drawFlag is true
			graphics.UpdateGraph(chip8)
		}
		chip8.DrawFlag = false

	}
}
