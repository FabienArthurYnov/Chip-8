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

func DetectedKeyReleased(display *pixelgl.Window, keyPressed []bool) []bool {

	keyPressed = append(keyPressed, display.JustReleased(pixelgl.KeyX))
	keyPressed = append(keyPressed, display.JustReleased(pixelgl.Key1))
	keyPressed = append(keyPressed, display.JustReleased(pixelgl.Key2))
	keyPressed = append(keyPressed, display.JustReleased(pixelgl.Key3))
	keyPressed = append(keyPressed, display.JustReleased(pixelgl.KeyQ))
	keyPressed = append(keyPressed, display.JustReleased(pixelgl.KeyW))
	keyPressed = append(keyPressed, display.JustReleased(pixelgl.KeyE))
	keyPressed = append(keyPressed, display.JustReleased(pixelgl.KeyA))
	keyPressed = append(keyPressed, display.JustReleased(pixelgl.KeyS))
	keyPressed = append(keyPressed, display.JustReleased(pixelgl.KeyD))
	keyPressed = append(keyPressed, display.JustReleased(pixelgl.KeyZ))
	keyPressed = append(keyPressed, display.JustReleased(pixelgl.KeyC))
	keyPressed = append(keyPressed, display.JustReleased(pixelgl.Key4))
	keyPressed = append(keyPressed, display.JustReleased(pixelgl.KeyR))
	keyPressed = append(keyPressed, display.JustReleased(pixelgl.KeyF))
	keyPressed = append(keyPressed, display.JustReleased(pixelgl.KeyV))

	display.Update()

	return keyPressed
}
