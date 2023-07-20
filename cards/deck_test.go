package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) { //타입캐스팅, 랜덤만들때 씨드 초기화해주는 부분
	d := newDeck()

	if len(d) != 15 {
		t.Errorf("16개가 나와야하는데, 나온게 = %v", len(d))
	}

	if d[0] != "Ace of King" {
		t.Errorf("wrong value = %v", d[0])
	}
}

func TestFileSaveAndLoad(t *testing.T) {
	fileName := "_testdeck"
	os.Remove(fileName)
	d := newDeck()
	d.saveToFile(fileName)
	loadFromFileDeck := newDeckFromFile(fileName)

	if len(d) != len(loadFromFileDeck) {
		t.Errorf("wrrog")
	}

	if len(d) != 15 {
		t.Errorf("wrrog")
	}
	os.Remove(fileName)
}
