package main

import (
	"fmt"
	"os"
	"strings"
)

// create a new type of deck
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}

	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
func (d deck) savetoFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}
