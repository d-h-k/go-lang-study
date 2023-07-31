package main

import (
	"fmt"
)

func main() {

	// map  초기화 1
	colors := mapMake1()
	fmt.Println(colors)

	// map  초기화 2
	c2 := make(map[string]string)
	c2["white"] = "#ffffff"
	fmt.Println(c2)

	// map  초기화 3
	c3 := mapMake2()
	fmt.Println(c3)

	// map 으로 루프 돌리기
	printMap(mapMake1())
}

func mapMake1() map[string]string {
	return map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
}

func mapMake2() map[int]string {
	return map[int]string{
		1: "w",
		2: "2",
		3: "쓰리",
	}
}

func printMap(stringMap map[string]string) {
	println("\n map 으로 루프 돌리기")
	for color, hexcode := range stringMap {
		fmt.Println("Hex code for", color, "is", hexcode)
	}
	println("\n")
}
