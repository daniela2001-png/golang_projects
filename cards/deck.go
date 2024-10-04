package main

import (
	"fmt"
	"os"
	"strings"
)

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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(data), ",")
	return deck(s)
}
