package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// cards.saveToFile("mycards.txt")
}
