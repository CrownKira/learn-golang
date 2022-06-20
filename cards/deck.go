package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
// extend the functionality of string
// deck extends the behavior of a slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// replace with underscore to mark no use
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			// qn: why need to reassign
			// append() returns a copy
			cards = append(cards, value+" of " +suit)
		}
	}

	return cards
}

/**
(d deck): receiver of type deck
any new var with type deck 
gets access to the print method 
ie. add on the print method to deck type

d: copy of the deck (convention: single character) (it is like 'this' or 'self', ie. this instance)
deck: receiver type (attach this method to the type)

eg. 
cards.print() 
=== 
print(cards)
where print() method is defined as: 
print(cards deck)

in python: 
to create Array object
- create inner array
- loop thru inner array 
**/
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}


// arguments in function call
// multiple return types
/**
why no use receiver:
cards.deal(5): implies that cards would have been modified, would have 5 fewer cards
deal(...)
**/
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")	
}

// error: error type
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

/**
err: nil if no error 
**/
func newDeckFromFile(filename string) deck {
	bs, err :=	ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log error and return call to newDeck()
		// Option #2 - log error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1) // 1: something went wrong
	}
	s :=	strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	/**
	https://github.com/golang/go/wiki/CodeReviewComments#variable-names
	Variable names in Go should be short rather than long. This is especially true for local variables with limited scope. Prefer c to lineCount. Prefer i to sliceIndex. 
	The basic rule: the further from its declaration that a name is used, the more descriptive the name must be. For a method receiver, one or two letters is sufficient. Common variables such as loop indices and readers can be a single letter (i, r). More unusual things and global variables need more descriptive names.
	**/
	for i := range d {
		// pseudo random generator 
		// seed is used in generator 
		// if use same seed, then randomness is the same
		newPosition := r.Intn(len(d) - 1)

		// swap syntax, without temp
		// newPosition -> i
		// i -> newPosition
		// no temp needed
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}