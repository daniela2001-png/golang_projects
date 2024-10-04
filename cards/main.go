package main

import "fmt"

func main() {

	cards := newDeck()
	handDeck, remainingDeck := deal(cards, 5)
	handDeck.print()
	remainingDeck.print()
	myString := cards.toString()
	fmt.Println(myString)
	cards.saveToFile("my_cards")
	newCards := newDeckFromFile("my_cards")
	newCards.print()
}
