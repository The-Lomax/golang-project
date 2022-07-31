package main

import "fmt"

type Card struct {
	Number string
	Face   string
}

func (c Card) Show() {
	fmt.Println(c.Number + " of " + c.Face)
}

func newCard(face, number string) Card {
	return Card{
		Face:   face,
		Number: number,
	}
}
