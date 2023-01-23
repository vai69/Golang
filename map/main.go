package main

import (
	"fmt"
)

func main() {
colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4f5654",
	}
	// delete(colors, "red")


	print(colors)

	fmt.Println(colors)
}

func print(m map[string]string) {
	for color, hex := range m{
		fmt.Println(color, hex)
	}

}