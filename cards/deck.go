package main

import "fmt"

type deck []string

// d is a copy of the deck type
func (d deck) print() {
	// Any variable of type deck can have access to this method
	for i, value := range d {
		fmt.Println(i, value)
	}
}

func newDeck() deck {

	suites := []string{"Spaids", "Hearts", "Diamonds", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four"}

	cards := deck{}

	for _, suite := range suites {
		for _, value := range values {
			cards = append(cards, value+" of "+suite)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
