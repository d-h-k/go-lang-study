package main // 실행가능한 패키지

import "fmt"

func main() {

	var card string = "ACE of spades"
	fmt.Println(card)
	//pair := "dsf"

	gogo := newCard()
	fmt.Println(gogo)

	pair := "Pair card" // 콜론이콜
	fmt.Println(pair)
}

func newCard() string {
	return "five of Dia"
}
