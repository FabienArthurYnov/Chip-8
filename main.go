package main

import (
	"chip-8/periph"
	"chip-8/structs"
)

func main() {
	chip8 := structs.Chip8{}

	periph.SetupGraphics()
	periph.SetupInput()

	chip8.Load() // load the game into the memory

	for true {
		chip8.EmulateOneCycle() // emulate one cycle

		if chip8.DrawFlag { //if drawFlag is true
			//DrawGraphics()
		}
		chip8.DrawFlag = false

		periph.SetKeys() // set the keys pressed
	}
}
