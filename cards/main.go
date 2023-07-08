package main // Executeable type package

// import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()

}
