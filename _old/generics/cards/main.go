package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card interface {
	String() string // Card is required to have String() method
}

type Deck[C Card] struct {
	Cards []C
}

func (d *Deck[C]) Add(card C) {
	d.Cards = append(d.Cards, card)
}

func (d *Deck[C]) Random() C {
	return choice(d.Cards)
}

// inspired by Python's random.choice(); it's so elegant!
func choice[T any](items []T) T {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	index := random.Intn(len(items))
	return items[index]
}

func main() {
	deck1 := NewPokerDeck()
	fmt.Println(deck1.Random())

	deck2 := NewWarfareDeck()
	fmt.Println(deck2.Random())
}
