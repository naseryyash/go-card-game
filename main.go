package main

func main() {
	cards := newDeckFromFile("my_cards")

	cards.print()
	// the remaining functions can be called & used as follows if needed:
	// (calls can be in any order, not necessarily in the order written below)

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
	// cards.toString()
	// cards.saveToFile("my_cards") -> the result of this is the file "my_cards", which can be found in the project folder
	// newDeckFromFile("my_cards") -> trying to retreive the file we just created with the above line
	cards.shuffle()

	// receiver functions like toString(), saveToFile() & shuffle()
	// are called with a cards type for e.g. cards.toString()
}
