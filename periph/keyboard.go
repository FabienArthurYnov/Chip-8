package periph

import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
)

func SetupInput() {

}

func SetKeys() {

}

func detectedKey() []bool {
	if err := keyboard.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer keyboard.Close()

	fmt.Println("Press 'a' key to check if it's pressed (press 'q' to quit)")

	var keyPressed []bool

	for {
		char, _, err := keyboard.GetKey()
		if err != nil {
			fmt.Println(err)
			break
		}

		if char == '1' {

			keyPressed = append(keyPressed, true)
		} else if char == '2' {

			keyPressed = append(keyPressed, true)
		} else if char == '3' {

			keyPressed = append(keyPressed, true)
		} else if char == '4' {

			keyPressed = append(keyPressed, true)
		} else if char == 'q' || char == 'Q' {

			keyPressed = append(keyPressed, true)
		} else if char == 'w' || char == 'W' {

			keyPressed = append(keyPressed, true)
		} else if char == 'e' || char == 'E' {

			keyPressed = append(keyPressed, true)
		} else if char == 'r' || char == 'R' {

			keyPressed = append(keyPressed, true)
		} else if char == 'a' || char == 'A' {

			keyPressed = append(keyPressed, true)
		} else if char == 's' || char == 'S' {

			keyPressed = append(keyPressed, true)
		} else if char == 'd' || char == 'D' {

			keyPressed = append(keyPressed, true)
		} else if char == 'f' || char == 'F' {

			keyPressed = append(keyPressed, true)
		} else if char == 'z' || char == 'Z' {

			keyPressed = append(keyPressed, true)
		} else if char == 'x' || char == 'X' {

			keyPressed = append(keyPressed, true)
		} else if char == 'c' || char == 'C' {

			keyPressed = append(keyPressed, true)
		} else if char == 'v' || char == 'V' {

			keyPressed = append(keyPressed, true)
		} else {

			keyPressed = append(keyPressed, false)
		}
	}
	return keyPressed

}
