package main

import "fmt"

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("Hex code for ", color, " is", hex)
	}
}

func main() {
	colors := make(map[string]string)

	colors["red"] = "#ff0000"
	colors["green"] = "#0000ff"
	colors["white"] = "#ffffff"

	delete(colors, "green")

	printMap(colors)
}
