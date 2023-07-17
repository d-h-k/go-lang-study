package main

import "fmt"

func main() {

	cards := newDeck()
	//cards = append(cards, "six of spade")
	//fmt.Println(cards)

	// // 반복문
	for i, card := range cards {
		fmt.Println(i, card)
	}
}
