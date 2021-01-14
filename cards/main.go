package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()
	cards.print()

	d := cards.toString()
	fmt.Println(d)

	// cards.saveToFile("my_cards")

	savedCards := newDeckFromFile("my_cards")

	savedCards.print()

}
