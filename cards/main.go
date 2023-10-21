package main

func main() {

	// fmt.Println("hello world")
	// cards := newDeck()

	// handCards, remainingCards := cards.slice(5)
	// handCards.print()
	// remainingCards.print()
	// cards.print()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	newDeckFromFile("my_cards").print()
}
