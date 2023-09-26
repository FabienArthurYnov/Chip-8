package utility

import "fmt"

func InputFileName() string {
	var name string

	fmt.Println("Mettez le nom du jeu")

	fmt.Scanln(&name)
	return name
}