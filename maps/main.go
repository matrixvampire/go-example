package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#FF00000",
		"green": "#00FF00",
	}

	// var colors map[string]string

	// colors := make(map[string]string)

	colors["white"] = "#FFFFFF"

	// delete(colors, "white")

	// fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
