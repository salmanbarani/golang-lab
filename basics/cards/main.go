package main // Executeable type package

import "fmt"

func main() {
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Life Is Short, Enjoy The Most!!!"
}
