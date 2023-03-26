package main

func main() {
	cards := newDeck()

	cards.saveToFile("my_cards")
	d := newDeckFromFile("my_cards")
	//cards.shuffle()
	d.print()

}
