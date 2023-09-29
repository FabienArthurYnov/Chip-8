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


func DetectedKey(char rune, keyPressed []bool) []bool {
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


        for {
            char, _, err := keyboard.GetKey()
            if err != nil {
                fmt.Println(err)
                break
            }
    
            if char == 'x' || char == 'X' {
                keyPressed = append(keyPressed, true)
                // fmt.Println("0")
    
            } else {
                keyPressed = append(keyPressed, false)
                // fmt.Println("mauvaise touche")
            }
            if char == '1' {
                keyPressed = append(keyPressed, true)
                // fmt.Println("1")
    
            } else {
                keyPressed = append(keyPressed, false)
                // fmt.Println("mauvaise touche")
            }
            if char == '2' {
                keyPressed = append(keyPressed, true)
                // fmt.Println("2")
    
            } else {
                keyPressed = append(keyPressed, false)
                // fmt.Println("mauvaise touche")
            }
            if char == '3' {
                keyPressed = append(keyPressed, true)
                // fmt.Println("3")
    
            } else {
                keyPressed = append(keyPressed, false)
                // fmt.Println("mauvaise touche")
            }
            if char == '4' {
                keyPressed = append(keyPressed, true)
                // fmt.Println("c")
    
            } else {
                keyPressed = append(keyPressed, false)
                // fmt.Println("mauvaise touche")
            }
            if char == 'q' || char == 'Q' {
                keyPressed = append(keyPressed, true)
                // fmt.Println("4")
    
            } else {
                keyPressed = append(keyPressed, false)
                // fmt.Println("mauvaise touche")
            }
            if char == 'w' || char == 'W' {
                keyPressed = append(keyPressed, true)
                // fmt.Println("5")
    
            } else {
                keyPressed = append(keyPressed, false)
                // fmt.Println("mauvaise touche")
            }
            if char == 'e' || char == 'E' {
                keyPressed = append(keyPressed, true)
                // fmt.Println("6")
    
            } else {
                keyPressed = append(keyPressed, false)
                // fmt.Println("mauvaise touche")
            }
            if char == 'r' || char == 'R' {
                keyPressed = append(keyPressed, true)
                // fmt.Println("d")
    
            } else {
                keyPressed = append(keyPressed, false)
                // fmt.Println("mauvaise touche")
            }
            if char == 'a' || char == 'A' {
                keyPressed = append(keyPressed, true)
                // fmt.Println("7")
    
            } else {
                keyPressed = append(keyPressed, false)
                // fmt.Println("mauvaise touche")
            }
            if char == 's' || char == 'S' {
                keyPressed = append(keyPressed, true)
                // fmt.Println("8")
    
            } else {
                keyPressed = append(keyPressed, false)
                // fmt.Println("mauvaise touche")
            }
            if char == 'd' || char == 'D' {
                keyPressed = append(keyPressed, true)
                // fmt.Println("9")
    
            } else {
                keyPressed = append(keyPressed, false)
                // fmt.Println("mauvaise touche")
            }
            if char == 'f' || char == 'F' {
                keyPressed = append(keyPressed, true)
                // fmt.Println("e")
    
            } else {
                keyPressed = append(keyPressed, false)
                // fmt.Println("mauvaise touche")
            }
            if char == 'z' || char == 'Z' {
                keyPressed = append(keyPressed, true)
                // fmt.Println("a")
    
            } else {
                keyPressed = append(keyPressed, false)
                // fmt.Println("mauvaise touche")
            }
            
            if char == 'c' || char == 'C' {
                keyPressed = append(keyPressed, true)
                // fmt.Println("b")
    
            } else {
                keyPressed = append(keyPressed, false)
                // fmt.Println("mauvaise touche")
            }
            if char == 'v' || char == 'V' {
                keyPressed = append(keyPressed, true)
                // fmt.Println("f")
    
            } else {
                // fmt.Println("mauvaise touche")
                keyPressed = append(keyPressed, false)
            }
    
            keyPressed = nil
        }
    return keyPressed
}