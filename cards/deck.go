package main

import "fmt"

// decks << ??
type deck []string

func (cards deck) print() {
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

// go 에서는 self 같은건 없음 , this 같은것도 없음
// 언어 전체에서 보게 될 형태임 >> 객체와 그 객체가 사용할 메서드까지
