// Package deck implements a tiny playing‑card library.
package deck

import "fmt"

// Deck is a slice of card strings (e.g. "Ace of Spades").
type Deck []string

// NewDeck returns a fresh, ordered 52‑card deck.
func NewDeck() Deck {
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six",
		"Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}

	var d Deck
	for _, s := range suits {
		for _, v := range values {
			d = append(d, fmt.Sprintf("%s of %s", v, s))
		}
	}
	return d
}

// Print shows one card per line.
func (d Deck) Print() {
	for i, card := range d {
		fmt.Printf("%2d: %s\n", i, card)
	}
}

// Deal splits the deck into <handSize> cards and the remainder.
func Deal(d Deck, handSize int) (Deck, Deck) {
	if handSize > len(d) {
		handSize = len(d)
	}
	return d[:handSize], d[handSize:]
}
