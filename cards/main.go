package main

func main() {
	cards := newDeck()

	filename := "full_deck.txt"
	cards.saveToFile(filename)

	cards_new := newDeckFromFile(filename)
	// cards_new.print()
	cards_new.shuffle()
	cards_new.print()
}
