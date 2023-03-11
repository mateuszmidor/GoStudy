package main

import "fmt"

// implements Card
type PokerCard struct {
	Color  string
	Figure string
}

func (c PokerCard) String() string {
	return fmt.Sprintf("%s of %s", c.Figure, c.Color)
}

func NewPokerDeck() *Deck[PokerCard] {
	colors := []string{"Diamonds", "Hearts", "Clubs", "Spades"}
	figures := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	deck := &Deck[PokerCard]{}
	for _, c := range colors {
		for _, f := range figures {
			card := PokerCard{Color: c, Figure: f}
			deck.Add(card)
		}
	}
	return deck
}
