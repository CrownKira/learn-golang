package main

// vscode auto adds this import statement

/**
go lang does not have OOP

function with a receiver

file structure:
main.go:
create and manipulate deck
deck.go:
deck_test.go:

**/

func main() {

	/*
		variable declaration
		var card string = "Ace of Spades"
		var: about to create a new var
		card: name of the var
		string: type of the variable
		"": value of the variable

		type:
		1. dynamic: js, ruby, python
		2. static: c++, java, go

		types in go:
		- bool: true, false
		- string: "hi"
		- int: 0, -1000
		- float64: 10.0001

		type inference

	*/
	// var card string = "Ace of Spades"
	// := declare new var with type inference; equivalent to above
	// usually use the below
	// card := "Ace of Spades"
	// card := newCard()
	// fmt.Println(card)

	/**
	Array: fixed length
	Slice: can grow and shrink
	**/
	// slice
	// cards := deck{"Ace of Diamonds", newCard()} // slice of type string
	// append string to cards slice
	// cards = append(cards, "Six of Spades")

	// for (let i = 0; i <... ; i++) {}
	// for i, card in enumerate(...)

	/**
	for index, card := range cards {...}
	index: index of this element
	card: current card in iteration
	:= similar to 'in', 'of' syntax in other lang
	range cards: take slice of cards and loop over it
	**/
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// cards := newDeckFromFile("my_cards")
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	cards := newDeck()
	cards.shuffle()
	cards.print()

}

/**
func: func declaration keyword
newCard: name of function
string: return type
**/
func newCard() string {
	return "Five of Diamonds"
}

/**
byte slice:
every char correspond to ascii value

standard lib: expect byte slice instead of string slice
data []byte: essentially a string

type conversion with "go"
turn string into byte slice:
[]byte("Hi there")
**/
