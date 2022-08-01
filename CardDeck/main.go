package main

import "fmt"

func main() {
	deck := Deck{}

	if err := deck.loadFromJSONFile("testFiles/cleanDeck.json"); err == nil {
		hand := Deck{}
		deck.Shuffle()
		if err := deck.Deal(3, &hand); err == nil {
			hand.saveToJSONFile("dealtHand.json")
		} else {
			fmt.Println(err)
		}
	}
}
