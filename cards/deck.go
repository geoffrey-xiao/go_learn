package main

import (
	"fmt"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuit := []string{"summer", "feng", "lei"}
	cardValue := []string{"one", "two", "three"}

	for _, suit := range cardSuit {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) slice(handPosition int) (deck, deck) {
	handCards := d[:handPosition]
	remainingCards := d[handPosition:]

	return handCards, remainingCards
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("error")
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ","))

}
