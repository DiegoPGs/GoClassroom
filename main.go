package main

import (
	"fmt"

	"example.com/cards/deck"
)

func main() {
	d := deck.NewDeck()

	hand, remaining := deck.Deal(d, 5)

	fmt.Println("Hand:")
	hand.Print()

	fmt.Printf("\nRemaining cards: %d\n", len(remaining))
}
