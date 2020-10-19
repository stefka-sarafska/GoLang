package main

import "fmt"

func main() {
	//cards := newDeck()
	//fmt.Println(cards.toString())
	//cards.saveToFile("cards")
	// cards := newDeckFromFile("cards")
	cards := newDeck()
	cards.print()
	fmt.Println("-------------------")
	cards.shuffle()
	cards.print()

}
