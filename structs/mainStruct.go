package structs

type Chip8 struct {
	Opcode byte;
	Memory [4096]byte;
	Reg [16]byte;	// register (memoire temporaire)

	I int;	// index register
	Pc int; //program counter

	Screen [64][32]bool; // true = white pixel, false = black pixel
	Keyboard [16]byte; // not sure about type
}

func (chip8 *Chip8) Load() {
	chip8.Pc = 512
	//WIP
}


func (chip8 *Chip8) EmulateOneCycle() {
	//fetch opcode
	var first uint16 = uint16(chip8.Memory[chip8.Pc])
	var second uint16 = uint16(chip8.Memory[chip8.Pc+1])
	first = first << 8  //décale le byte de 8 / ajoute un byte derrière

	var opcode uint16 = first | second // OR du byte vide  
	
	opcode1 := opcode >> 12  // first hexa number
	opcode2 := (opcode >> 8) & 0x0F  // second...
	opcode3 := (opcode >> 4) & 0x00F
	opcode4 := opcode & 0x000F


	//decode opcode
	switch opcode1 {
	case 0x0 : 
		if (opcode == 0x00E0) {
			// clear screen
		}	
		if (opcode == 0x00EE) {
			// return from subroutine
		}

	case 0x1 :
		//temp := opcode & 0x0FFF
		//goto temp
	
	case 0x2 :
		//temp := opcode & 0x0FFF
		//call subroutine at temp

	case 0x3 :


	}

	//execute opcode
}
