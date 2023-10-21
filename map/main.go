package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":  "#00011",
		"blue": "#00222",
	}

	colors["white"] = "#00099"

	delete(colors, "red")

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
