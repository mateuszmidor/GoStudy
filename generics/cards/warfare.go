package main

// implements Card
type WarfareCard struct {
	Vehicle string
}

func (c WarfareCard) String() string {
	return c.Vehicle
}

func NewWarfareDeck() *Deck[WarfareCard] {
	vehicles := []string{"Tank", "Cannon", "Amphibia", "Humv"}
	deck := &Deck[WarfareCard]{}
	for _, v := range vehicles {
		card := WarfareCard{Vehicle: v}
		deck.Add(card)
	}
	return deck
}
