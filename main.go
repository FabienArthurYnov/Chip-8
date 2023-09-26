package main

import (
	"chip-8/periph"
	"chip-8/structs"
	"chip-8/utility"
	"time"
)

func main() {
	chip8 := structs.Chip8{}
	timeStart := time.Now()

	periph.SetupGraphics()
	periph.SetupInput()

	var fileName string = utility.InputFileName()

	chip8.Load(fileName) // load the game into the memory

	for true {
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
			//DrawGraphics()
		}
		chip8.DrawFlag = false

		periph.SetKeys() // set the keys pressed
	}
}
