package main

import "fmt"

// decks << ??
type Deck []string

func newDeck() Deck {
	cards := Deck{}
	//슈트랑 벨류스
	cardSuits := []string{"Ace", "Two", "Three"}
	cardValues := []string{"King", "Five", "Two", "Three"}

	for _, suit := range cardSuits {
		//for i, suit := range cardSuits {
		for _, value := range cardValues {
			//for j, value := range cardValues {
			cards = append(cards, suit+"of"+value)
		}
	}
	return cards
}

func (cards Deck) print() {
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

// go 에서는 self 같은건 없음 , this 같은것도 없음
// 언어 전체에서 보게 될 형태임 >> 객체와 그 객체가 사용할 메서드까지
