//we are going to declare deck type for additional functionality

package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck
//which is a type of string

type deck []string

//creating a new Deck function to call the function of cards

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

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]

}

//to convert a deck into a string
// using both reciever and the function to convert the the deck into the string
//converting slice of deck to slice of string and then converting the slice of string
//into a single string separated by commas

func (d deck) toString() string {
	return strings.Join([]string(d), ",") //strings package used

}

//saving the deck to the hardrive
//0666 is the default permission for any file

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}

//bs is byte slice
//for readng a file we have to do the opposite

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	//Ace of Spades, Two of Spades..

	s := strings.Split(string(bs), ",")
	return deck(s)

}

//shuffling of the cards

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
