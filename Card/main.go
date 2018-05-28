package main

func main() {
	// Writing data to File.
	cards := newDeck()
	cards.saveToFile("myCards")

	//Reading from File
	cardsReadFromFile := newDeckFromFile("myCards")
	cardsReadFromFile.print()

	//Reading File which does not exist
	//cardsReadFromFile = newDeckFromFile("myCards1")
	//cardsReadFromFile.print()

	// Shuffle the card
	cards.shuffle()
	cards.print()

	/**
	  Experimenting with other methods.
	**/
	//fmt.Println(cards.toString())
	//cards.print()
	//hand, remainingCards := deal(cards, 5)

	//hand.print()
	//remainingCards.print()

}
