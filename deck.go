package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card
}

func (d *Deck) Deal(n int, hand *Deck) error {
	if len(d.Cards) < n {
		return fmt.Errorf("Not enough cards to deal")
	}
	hand.Cards = d.Cards[:n]
	d.Cards = d.Cards[n:]
	return nil
}

func (d Deck) ShowCards() {
	for _, c := range d.Cards {
		c.Show()
	}
}

func (d *Deck) Shuffle() {
	copyDeck := *d
	drawnNumbers := []int{}
	d.Cards = []Card{}
	for i := range copyDeck.Cards {
		_ = i
		rand.Seed(time.Now().UnixNano())
		for {
			idx := rand.Intn(len(copyDeck.Cards))
			if !hasNumber(drawnNumbers, idx) {
				d.Cards = append(d.Cards, copyDeck.Cards[idx])
				drawnNumbers = append(drawnNumbers, idx)
				break
			}
		}
	}
}

func (d Deck) saveToJSONFile(fileName string) error {
	deckJSON, err := json.Marshal(&d)
	if err != nil {
		return fmt.Errorf("Cannot marshal deck, error: %v", err)
	}
	if err := ioutil.WriteFile(fileName, deckJSON, 0o644); err != nil {
		return fmt.Errorf("Could not save to file, error: %v", err)
	}
	return nil
}

func (d *Deck) loadFromJSONFile(fileName string) error {
	fbytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("Could not read from file, error: %v", err)
	}
	if err := json.Unmarshal(fbytes, d); err != nil {
		return fmt.Errorf("Cannot unmarshal deck, error: %v", err)
	}
	return nil
}

func newDeck() Deck {
	d := Deck{}
	numbers := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}
	faces := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	for _, face := range faces {
		for _, number := range numbers {
			d.Cards = append(d.Cards, newCard(face, number))
		}
	}
	return d
}

func hasNumber(arr []int, n int) bool {
	for _, i := range arr {
		if i == n {
			return true
		}
	}
	return false
}
