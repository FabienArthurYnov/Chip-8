package structs

import (
	"chip-8/utility"
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/faiface/pixel/pixelgl"
)

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

	ScreenTable [64][32]bool // true = white pixel, false = black pixel
	Display     *pixelgl.Window
	Keyboard    [16]byte // not sure about type

	DrawFlag bool // do we update the screen ?  yes when clear screen or draw sprite
}

func (chip8 *Chip8) Load(fileName string) {
	chip8.Pc = 512

	file, err := os.OpenFile("./rom/"+fileName+".ch8", os.O_RDONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fileByte, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(fileByte); i++ {
		chip8.Memory[i+512] = fileByte[i]
	}
	//WIP
}

func (chip8 *Chip8) EmulateOneCycle() {
	//fetch opcode  ex 0xA23B
	var opcode uint16

	// freeze the program if end of memory / program. Or out of range.
	if chip8.Pc >= 4096-1 {
		fmt.Println("End of program")
		opcode = 0x0
		time.Sleep(time.Minute)
	} else { // else get the opcode
		var first uint16 = uint16(chip8.Memory[chip8.Pc])    // 0xA2
		var second uint16 = uint16(chip8.Memory[chip8.Pc+1]) // 0x3B
		first = first << 8                                   // 0xA200  décale le byte de 8 / ajoute un byte derrière

		opcode = first | second // 0xA23B   OR du byte vide
	}

	// DEBUG
	// fmt.Printf("0x%X\n", opcode)
	// time.Sleep(time.Second)

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
			chip8.Display.Clear(color.Black)
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
		/*Draws a sprite at coordinate (VX, VY) that has a width of 8 pixels and a height of N pixels.
		Each row of 8 pixels is read as bit-coded starting from memory location I; I value does not change after the execution of this instruction.
		As described above, VF is set to 1 if any screen pixels are flipped from set to unset when the sprite is drawn, and to 0 if that does not happen.*/
		chip8.DrawFlag = true
		tempI := chip8.I
		x := chip8.Reg[opcode2]
		y := chip8.Reg[opcode3]
		n := opcode4

		fmt.Printf("%X\n", n)

		nInt := int(n)

		fmt.Println(nInt)

		setToTrue := false

		// fmt.Printf("0x%X\n", opcode) //DEBUG opcode and x,y,n,I
		// fmt.Println(x, y, n, tempI)

		for i := 0; i < nInt; i++ { // par ligne
			rowByte := chip8.Memory[tempI]
			// fmt.Printf("0b%b\n", rowByte)  //debug view all bits
			tempI += 1
			for j := 8; j > 0; j-- { // par pixel
				bit := (rowByte & (1 << j)) != 0
				if !chip8.ScreenTable[x][32-y] && bit { // si pixel set to true
					setToTrue = true
				}
				chip8.ScreenTable[x][32-y] = bit
				x += 1
			}
			y += 1
			x = chip8.Reg[opcode2] // to make a new line, need to go back to start X

		}
		if setToTrue {
			chip8.Reg[0xF] = 1
		} else {
			chip8.Reg[0xF] = 0
		}

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
