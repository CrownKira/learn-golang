package main

import "fmt"


func main() {
	// key: type string
	// value: type string
	colors := map[string]string {
		"red": "#ff000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// var colors map[string]string
	// make(): built in syntax 
	// equivalent to above, init zero value: ie. no value in map
	// colors := make(map[string]string)

	// add key value pair
	// cannot do colors.white
	// colors["white"] = "#ffffff"

	// delete(colors, "white")

	// fmt.Println(colors)
	printMap(colors)
}


func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

/**
Differences between map and struct:

see notion
**/