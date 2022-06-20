package main

import "fmt"

/**
eg. interface
type bot interface {
	getGreeting(string, int) (string, error)
}
**/
/**
type:
concrete typ: map, struct, int, string
interface: bot (cannot create a value out of this type)
**/
/**
interface:
- write code that is concise
- not generic type
- implicit (dont need explicitly implement)
- contract: help us manage types

**/
type bot interface {
	// ie. any type that has these set of functions
	// implement this interface
	// dun need to explicitly implement the interface
	// used to extract similar logic
	getGreeting() string
}

type englishBot struct {}
type spanishBot struct {}


func main() {
	// declare struct 
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}


func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

/**
go does not support overloading
ie. cannot have same function in one package
**/
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

// (eb englishBot) -> (englishBot) if no use eb
func (englishBot) getGreeting() string {
	 // very custom logic for generating an english greeting
	 return "Hi There"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}