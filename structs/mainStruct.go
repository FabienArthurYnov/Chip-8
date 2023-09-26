package structs

import "chip-8/utility"

type Chip8 struct {
	Opcode byte
	Memory [4096]byte
	Reg    [16]byte // register (memoire temporaire)

	I  uint16 // 16-bit index register for memory address
	Pc int    //program counter

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
		}

	case 0x1:
		chip8.Pc = int(opcode234) //goto opcode234

	case 0x2:
		//call subroutine at opcode234

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

		}

	}

	chip8.Pc += 2
}
