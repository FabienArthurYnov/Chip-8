package structs

import "chip-8/utility"

type Chip8 struct {
	Opcode uint16
	Memory [4096]byte
	Reg    [16]byte // register (memoire temporaire)

	I  uint16 // 16-bit index register for memory address
	Pc int    //program counter

	DelayTimer int // timer -1/s
	SoundTimer int //timer -1/s

	Stack        [16]int // addresses stacks
	StackPointer int     // where are we in the stack

	Screen   [64][32]bool // true = white pixel, false = black pixel
	Keyboard [16]byte     // not sure about type

	DrawFlag bool // do we update the screen ?  yes when clear screen or draw sprite
}

func (chip8 *Chip8) Load() {
	chip8.Pc = 512
	//WIP
}

func (chip8 *Chip8) EmulateOneCycle() {
	//fetch opcode  ex 0xA23B
	var first uint16 = uint16(chip8.Memory[chip8.Pc])    // 0xA2
	var second uint16 = uint16(chip8.Memory[chip8.Pc+1]) // 0x3B
	first = first << 8                                   // 0xA200  décale le byte de 8 / ajoute un byte derrière

	var opcode uint16 = first | second // 0xA23B   OR du byte vide
	chip8.Opcode = opcode

	opcode1 := opcode >> 12          // 0xA  first hexa number
	opcode2 := (opcode >> 8) & 0x0F  // 0x2  second...
	opcode3 := (opcode >> 4) & 0x00F // 0x3
	opcode4 := opcode & 0x000F       // 0xB

	opcode234 := opcode & 0x0FFF // 0x23B
	opcode34 := byte(opcode)     // 0x3B  truncate to the last 8 bits, like doing opcode & 0x00FF

	//decode opcode
	switch opcode1 {
	case 0x0:
		if opcode == 0x00E0 {
			chip8.DrawFlag = true
			// clear screen
		}
		if opcode == 0x00EE {
			// return from subroutine
			chip8.StackPointer--
			chip8.Pc = chip8.Stack[chip8.StackPointer]
		}

	case 0x1:
		chip8.Pc = int(opcode234) //goto opcode234

	case 0x2:
		//call subroutine at opcode234
		// add the pc now to the stack, and go to the next stack pointer
		chip8.Stack[chip8.StackPointer] = chip8.Pc
		chip8.StackPointer++
		// go to subroutine
		chip8.Pc = int(chip8.Reg[opcode234])

	case 0x3:
		if chip8.Reg[opcode2] == opcode34 {
			chip8.Pc += 2 //skip next instruction
		}

	case 0x4:
		if chip8.Reg[opcode2] != opcode34 {
			chip8.Pc += 2 //skip next instruction
		}

	case 0x5:
		if opcode4 == 0x0 {
			if chip8.Reg[opcode2] == chip8.Reg[opcode3] {
				chip8.Pc += 2 //skip next instruction
			}
		}

	case 0x6:
		chip8.Reg[opcode2] = opcode34

	case 0x7:
		if chip8.Reg[opcode2] != 0xF { // if not flag
			chip8.Reg[opcode2] += opcode34
		}

	case 0x8:
		switch opcode4 {
		case 0x0:
			chip8.Reg[opcode2] = chip8.Reg[opcode3]

		case 0x1:
			chip8.Reg[opcode2] = chip8.Reg[opcode2] | chip8.Reg[opcode3]

		case 0x2:
			chip8.Reg[opcode2] = chip8.Reg[opcode2] & chip8.Reg[opcode3]

		case 0x3:
			chip8.Reg[opcode2] = chip8.Reg[opcode2] ^ chip8.Reg[opcode3]

		case 0x4:
			var temp1 uint16 = uint16(chip8.Reg[opcode2])
			var temp2 uint16 = uint16(chip8.Reg[opcode3])
			result := temp1 + temp2
			var mask uint16 = 0b100000000 // 9th bit
			if (result & mask) != 0 {     // then the 9th bit is 1, so there is a carry
				chip8.Reg[0xF] = 0x1
			} else {
				chip8.Reg[0xF] = 0x0
			}
			chip8.Reg[opcode2] = byte(result) // add

		case 0x5:
			if chip8.Reg[opcode2] < chip8.Reg[opcode3] { // will be negative, needs borrow : flag = 0
				chip8.Reg[0xF] = 0x0
			} else {
				chip8.Reg[0xF] = 0x1
			}
			chip8.Reg[opcode2] -= chip8.Reg[opcode3] //sub reg[op2]-reg[op3]

		case 0x6:
			chip8.Reg[0xF] = chip8.Reg[opcode2] & 0b00000001 // get least significant bit
			chip8.Reg[opcode2] = chip8.Reg[opcode2] >> 1     // shift to the right

		case 0x7:
			if chip8.Reg[opcode3] < chip8.Reg[opcode2] { // will be negative, needs borrow : flag = 0
				chip8.Reg[0xF] = 0x0
			} else {
				chip8.Reg[0xF] = 0x1
			}
			chip8.Reg[opcode2] = chip8.Reg[opcode3] - chip8.Reg[opcode2] //sub reg[op3]-reg[op2]

		case 0xE:
			chip8.Reg[0xF] = chip8.Reg[opcode2] & 0b10000000 // get most significant bit
			chip8.Reg[opcode2] = chip8.Reg[opcode2] << 1     // shift to the left
		}

	case 0x9:
		if opcode4 == 0x0 {
			if chip8.Reg[opcode2] != chip8.Reg[opcode3] {
				chip8.Pc += 2 // skip the next instruction
			}
		}

	case 0xa:
		chip8.I = opcode234 // I register set to opcode234

	case 0xb:
		chip8.Pc = int(opcode234) + int(chip8.Reg[0x0]) // goto opcode234 + reg0

	case 0xc:
		chip8.Reg[opcode2] = utility.RandomByte() & opcode34 // Vx = random number & opcode34

	case 0xd:
		chip8.DrawFlag = true
		/*Draws a sprite at coordinate (VX, VY) that has a width of 8 pixels and a height of N pixels.
		Each row of 8 pixels is read as bit-coded starting from memory location I; I value does not change after the execution of this instruction.
		As described above, VF is set to 1 if any screen pixels are flipped from set to unset when the sprite is drawn, and to 0 if that does not happen.*/

	case 0xe:
		switch opcode34 {
		case 0x9e:
			// if the keycode in chip8.Reg[opcode2] is pressed {
			// 		chip8.Pc += 2 // skip the next instruction
			// }

		case 0xa1:
			// if the keycode in chip8.Reg[opcode2] is NOT pressed {
			// 		chip8.Pc += 2 // skip the next instruction
			// }
		}

	case 0xf:
		switch opcode34 {
		case 0x07:
			chip8.Reg[opcode2] = byte(chip8.DelayTimer) // set Vx to DelayTimer

		case 0x0a:
			// A key press is awaited, and then stored in VX (blocking operation, all instruction halted until next key event).

		case 0x15:
			chip8.DelayTimer = int(chip8.Reg[opcode2]) // set DelayTimer to Vx

		case 0x18:
			chip8.SoundTimer = int(chip8.Reg[opcode2]) // set SoundTimer to Vx

		case 0x1e:
			chip8.I += uint16(chip8.Reg[opcode2])

		case 0x29:
			// font set start at 0 and each one take 5 bytes
			chip8.I = uint16(0x0 + (0x5 * chip8.Reg[opcode2]))

		case 0x33:
			// save l'expression binaire du décimal dans la mémoire
			numInt := int(chip8.Reg[opcode2])
			hundreds := numInt / 100
			tens := (numInt - hundreds*100) / 10
			units := (numInt - hundreds*100 - tens*10)
			chip8.Memory[chip8.I] = byte(hundreds)
			chip8.Memory[chip8.I] = byte(tens)
			chip8.Memory[chip8.I] = byte(units)

		case 0x55:
			// save dans la mémoire
			index := chip8.I
			for i := 0; i > int(opcode2); i++ {
				chip8.Memory[index] = chip8.Reg[i]
				index += 1
			}

		case 0x65:
			// load dans la mémoire
			index := chip8.I
			for i := 0; i > int(opcode2); i++ {
				chip8.Reg[i] = chip8.Memory[index]
				index += 1
			}
		}
	}

	chip8.Pc += 2 // go to next instruction
}
