package main // 실행가능한 패키지

import "fmt"

func main() {

	// var card string = "ACE of spades"
	// fmt.Println(card)
	//pair := "dsf"

	gogo := newCard()
	fmt.Println(gogo)

	// 콜론이콜 변수 할당
	pair := "Pair card"
	fmt.Println(pair)

	// 배열 자료구조
	cards := []string{myCard(), myCard(), newCard(), "sdf"}
	cards = append(cards, "six of spade")
	fmt.Println(cards)

	// 반복문
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

// 함수는 리턴이 있으면 자료구조를 명시해야함
func newCard() string {
	return "five of Dia"
}

func myCard() string {
	return "five of Dia"
}
