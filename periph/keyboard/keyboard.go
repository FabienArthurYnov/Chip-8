package keyboard

import (
	"fmt"
	"log"
	"os"

	"github.com/eiannone/keyboard"
	"github.com/faiface/pixel/pixelgl"
)

func SetupInputPaused() rune {
	if err := keyboard.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer keyboard.Close()

	char, _, err := keyboard.GetKey()
	if err != nil {
		log.Fatal(err)
	}
	return char
}

func DetectedKey(display *pixelgl.Window, keyPressed []bool) []bool {

	keyPressed = append(keyPressed, display.Pressed(pixelgl.KeyX))
	keyPressed = append(keyPressed, display.Pressed(pixelgl.Key1))
	keyPressed = append(keyPressed, display.Pressed(pixelgl.Key2))
	keyPressed = append(keyPressed, display.Pressed(pixelgl.Key3))
	keyPressed = append(keyPressed, display.Pressed(pixelgl.KeyQ))
	keyPressed = append(keyPressed, display.Pressed(pixelgl.KeyW))
	keyPressed = append(keyPressed, display.Pressed(pixelgl.KeyE))
	keyPressed = append(keyPressed, display.Pressed(pixelgl.KeyA))
	keyPressed = append(keyPressed, display.Pressed(pixelgl.KeyS))
	keyPressed = append(keyPressed, display.Pressed(pixelgl.KeyD))
	keyPressed = append(keyPressed, display.Pressed(pixelgl.KeyZ))
	keyPressed = append(keyPressed, display.Pressed(pixelgl.KeyC))
	keyPressed = append(keyPressed, display.Pressed(pixelgl.Key4))
	keyPressed = append(keyPressed, display.Pressed(pixelgl.KeyR))
	keyPressed = append(keyPressed, display.Pressed(pixelgl.KeyF))
	keyPressed = append(keyPressed, display.Pressed(pixelgl.KeyV))

	display.Update()

	return keyPressed
}

func DetectedKeyPaused() int {
	real := SetupInputPaused()
	final := -1

	switch real {
	case rune(pixelgl.KeyX):
		final = 0x0
	case rune(pixelgl.Key1):
		final = 0x1
	case rune(pixelgl.Key2):
		final = 0x2
	case rune(pixelgl.Key3):
		final = 0x3
	case rune(pixelgl.KeyQ):
		final = 0x4
	case rune(pixelgl.KeyW):
		final = 0x5
	case rune(pixelgl.KeyE):
		final = 0x6
	case rune(pixelgl.KeyA):
		final = 0x7
	case rune(pixelgl.KeyS):
		final = 0x8
	case rune(pixelgl.KeyD):
		final = 0x9
	case rune(pixelgl.KeyZ):
		final = 0xa
	case rune(pixelgl.KeyC):
		final = 0xb
	case rune(pixelgl.Key4):
		final = 0xc
	case rune(pixelgl.KeyR):
		final = 0xd
	case rune(pixelgl.KeyF):
		final = 0xe
	case rune(pixelgl.KeyV):
		final = 0xf
	}

	return final
}
