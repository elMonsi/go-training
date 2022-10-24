package main

import "fmt"

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color)
		fmt.Println(hex)
	}
}

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"black": "#000000",
	}

	//Altrnative ways to declare a map
	//var colors map[string]string
	//colors := make(map[string]string)

	colors["white"] = "#ffffff"

	delete(colors, "white")
	//fmt.Println(colors)

	printMap(colors)
}
