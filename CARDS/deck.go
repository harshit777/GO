//we are going to declare deck type for additional functionality

package main

import "fmt"

// Create a new type of deck
//which is a type of string

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// initialising a reciever function "func(d deck) print()" : sets up methods on variables like
//variable deck can have access to print function  "d can be said as "this" in java"

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
