package main

func main() {
	cards := newDeckFromFile("mycards.txt")
	cards.print()
	// cards.saveToFile("mycards.txt")
}
