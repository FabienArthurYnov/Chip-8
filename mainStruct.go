package main

type chip8 struct {
	opcode byte;
	memory []byte;
	reg []byte;	// register (memoire temporaire)

	i byte;	// index register
	pc byte; //program counter

	screen [][]bool; // true = white pixel, false = black pixel
	keyboard []byte; // not sure about type
}
