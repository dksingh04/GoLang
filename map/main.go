package main

import (
	"fmt"
)

func main() {
	color := map[string]string{
		"Red":   "Red",
		"Green": "Green",
	}

	fmt.Println(color)

	printMap(color)

	//other way of creating map

	//var colors map[string]string
	colors := make(map[string]string)
	//Adding value to map
	colors["White"] = "White"

	fmt.Println(colors)

	//deleting an element

	delete(colors, "White")

	fmt.Println(colors)

}

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Println("Key is ", k, "And value is ", v)
	}
}
