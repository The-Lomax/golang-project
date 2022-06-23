package main

func main() {
	deck := newDeck()

	card := deck.DrawCard()
	card.Show()

	deck.ShowDeck()
}

func newDeck() Deck {
	return Deck{
		Cards: []Card{
			newCard(),
			newCard(),
			newCard(),
			newCard(),
			newCard(),
		},
	}
}

func newCard() Card {
	return Card{
		Face:   "Spades",
		Number: "5",
	}
}
