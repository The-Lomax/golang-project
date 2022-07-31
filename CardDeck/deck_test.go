package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	t.Run("Check new deck has 52 cards", func(t *testing.T) {
		d := newDeck()
		got := len(d.Cards)
		want := 52

		if got != want {
			t.Errorf("Wanted: %v but got: %v", want, got)
		}
	})
	t.Run("Check first card is 2 of spades", func(t *testing.T) {
		d := newDeck()
		got := d.Cards[0]
		want := Card{
			Number: "2",
			Face:   "Spades",
		}

		if got != want {
			t.Errorf("Wanted first card be %v, got %v", want, got)
		}
	})
}

func TestLoadFromJSONFile(t *testing.T) {
	t.Run("check file loaded successfully", func(t *testing.T) {
		deck := newDeck()
		deck.saveToJSONFile("testFiles/cleanDeck.json")
		d := Deck{}
		d.loadFromJSONFile("testFiles/cleanDeck.json")
		got := len(d.Cards)
		want := 52

		if got != want {
			t.Errorf("Wanted: %v but got: %v", want, got)
		}
	})
}

func TestSaveToJSONFileAndLoadFromJSONFile(t *testing.T) {
	t.Run("Check that file is saved and loaded after properly", func(t *testing.T) {
		os.Remove("test.json")
		deck := newDeck()
		if err := deck.saveToJSONFile("test.json"); err != nil {
			t.Errorf("Failed to save JSON file: %v", err)
		}
		testDeck := Deck{}
		if err := testDeck.loadFromJSONFile("test.json"); err != nil {
			t.Errorf("Failed to load JSON file: %v", err)
		}
		got := len(testDeck.Cards)
		want := 52
		os.Remove("test.json")

		if got != want {
			t.Errorf("Wanted: %v but got: %v", want, got)
		}
	})
}
