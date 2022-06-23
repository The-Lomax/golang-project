package main

type Deck struct {
	Cards []Card
}

func (d Deck) DrawCard() Card {
	return d.Cards[0]
}

func (d Deck) ShowDeck() {
	for _, c := range d.Cards {
		c.Show()
	}
}

func (d Deck) Shuffle() {
	// do nothing
}
