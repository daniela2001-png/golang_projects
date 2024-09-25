package main

func main() {

	cards := newDeck()
	handDeck, remainingDeck := deal(cards, 5)
	handDeck.print()
	remainingDeck.print()
}
