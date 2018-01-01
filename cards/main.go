package main

func main() {
	cards := newDeckFromFile("my_cards")
	cards.suffle()
	hand, _ := deal(cards, 5)
	hand.print()
}
