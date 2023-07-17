package main

import (
	"fmt"
	"os" // 과거에는 io/util 패키지였는데 변경된거임
	"strings"
)

// decks << ??
type Deck []string

func newDeck() Deck {
	cards := Deck{}
	//슈트랑 벨류스
	cardSuits := []string{"Ace", "Spade", "Clover"}
	cardValues := []string{"King", "Five", "Two", "Three", "One"}

	for _, suit := range cardSuits {
		//for i, suit := range cardSuits {
		for _, value := range cardValues {
			//for j, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

func (cards Deck) print() {
	for i, card := range cards {
		fmt.Println(i, card)
	}
	// go 에서는 self 같은건 없음 , this 같은것도 없음
	// 언어 전체에서 보게 될 형태임 >> 객체와 그 객체가 사용할 메서드까지
}

func deal(d Deck, handSize int) (Deck, Deck) {
	return deck[:handSize], deck[handSize:]
}

func (d Deck) toString() string {
	return strings.Join([]string(deck), ",")
}

//Todo :  리시버랑 아규먼트 차이점 공부하기

func (d Deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}
