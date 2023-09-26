package structs

type Chip8 struct {
	Opcode byte;
	Memory [4096]byte;
	Reg [16]byte;	// register (memoire temporaire)

	I byte;	// index register
	Pc byte; //program counter

	Screen [64][32]bool; // true = white pixel, false = black pixel
	Keyboard [16]byte; // not sure about type
}
