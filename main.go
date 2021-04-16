package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
