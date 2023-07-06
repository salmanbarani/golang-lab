package main // Executeable type package

import "fmt"

func main() {
	cards := newDeck()

	fmt.Println(cards.toString())

}
