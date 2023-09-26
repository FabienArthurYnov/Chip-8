package main 

import (
	"chip-8/structs"
	"chip-8/periph"
)

func main() {
	chip8 := structs.Chip8{}

	periph.SetupGraphics()
	periph.SetupInput()

	chip8.Load() // load the game into the memory

	for (true) {
		chip8.EmulateOneCycle() // emulate one cycle

		if (chip8.Reg[15] == 1) { //if drawFlag is true
			drawGraphics()
		}

		periph.SetKeys()  // set the keys pressed
	}
}

