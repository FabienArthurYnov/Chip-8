package keyboard

import (
	"fmt"
	"log"
	"os"

	"github.com/eiannone/keyboard"
)

func SetupInput() rune {
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


func DetectedKey(char rune, keyPressed [16]bool) [16]bool {
    // if err := keyboard.Open(); err != nil {
    //     fmt.Println(err)
    //     os.Exit(1)
    // }
    // defer keyboard.Close()

    //     char, _, err := keyboard.GetKey()
    //     if err != nil {
    //         log.Fatal(err)
    //     }

		// var keyPressed [16]bool

        if char == '1' {
            keyPressed[0x1] = true
        } else {
			keyPressed[0x1] = false
		}
        if char == '2' {
            keyPressed[0x2] = true
        } else {
			keyPressed[0x2] = false
		}
        if char == '3' {
            keyPressed[0x3] = true
        } else {
			keyPressed[0x3] = false
		}
        if char == 'q' || char == 'Q' {
            keyPressed[0x4] = true
        } else {
			keyPressed[0x4] = false
		}
        if char == 'w' || char == 'W' {
            keyPressed[0x5] = true
        } else {
			keyPressed[0x5] = true
		}
        if char == 'e' || char == 'E' {
            keyPressed[0x6] = true
        } else {
			keyPressed[0x6] = false
		}
        if char == 'r' || char == 'R' {
            keyPressed[0xd] = true
        } else {
			keyPressed[0xd] = false
		}
        if char == 'a' || char == 'A' {
            keyPressed[0x7] = true
        } else {
			keyPressed[0x7] = false
		}
        if char == 's' || char == 'S' {
            keyPressed[0x8] = true
        } else {
			keyPressed[0x8] = false
		}
        if char == 'd' || char == 'D' {
            keyPressed[0x9] = true
        } else {
			keyPressed[0x9] = false
		}
        if char == 'f' || char == 'F' {
            keyPressed[0xe] = true
        } else {
			keyPressed[0xe] = false
		}
        if char == 'z' || char == 'Z' {
            keyPressed[0xa] = true
        } else {
			keyPressed[0xa] = false
		}
        if char == 'x' || char == 'X' {
            keyPressed[0x0] = true
        } else {
			keyPressed[0x0] = false
		}
        if char == 'c' || char == 'C' {
            keyPressed[0xb] = true
        } else {
			keyPressed[0xb] = false
		}
        if char == 'v' || char == 'V' {
            keyPressed[0xf] = true
        } else {
			keyPressed[0xf] = false
		}

    return keyPressed
}