package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
		"white": "#ffffff",
		"black": "#000000",
	}
	//var colors map[string]string
	fmt.Println(colors)
	//removing grrn color from map
	delete(colors, "green")
	fmt.Println(colors)
	//iterating over the maps
	printMap(colors)

}
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, ":", hex)
	}
}
