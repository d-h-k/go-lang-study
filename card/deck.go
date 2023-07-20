package main

import (
	"fmt"
	"math/rand"
	"time"

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
	return d[:handSize], d[handSize:]
}

func (d Deck) toString() string {
	return strings.Join([]string(d), ",")
}

//Todo :  리시버랑 아규먼트 차이점 공부하기

func (d Deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) Deck {
	data, err := os.ReadFile(fileName)

	if err != nil {
		//고의 오류처리는 쉽지않다
		//여기서 해야할 선택지 두개
		// Option1 : 새로운 덱을 만들어버림
		// Option2 : 프로그램 중단시켜버리기
		fmt.Println("cannot Read")
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}

	dataStrings := strings.Split(string(data), ",")
	//return dataStrings;
	return Deck(dataStrings)

	//nil << 널같은건데 널은 아닌거같고 값이 없다는뜻이긴 한듯
}

func (d Deck) shuffle() {

	myRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	for index := range d {
		newPosition := myRand.Intn(len(d) - 1)
		d[index], d[newPosition] = d[newPosition], d[index] // swap 을 아주 쉽게 할 수 있다
	}
}
