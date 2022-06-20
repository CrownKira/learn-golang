package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.com",
		"http://amazon.com",
	}

	// can treat it just like any other value
	// send data to and fro main routine and child routine
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)	
	}

	// for i := 0; i < len(links); i++ {
	// for { 
	// 	// for {} will loop for eternity; ie. while(true) {}
	// 	// wait for message to come thru channel then print
	// 	// fmt.Println(<-c)
	// 	// call checkLink() immediately after channel got message
	// 	// <-c is a blocking operation
	// 	go checkLink(<-c, c)
	// }

	// equivalent to above
	// range c: wait for the channel to return some value 
	// after get value, assign it to l 
	// range <iterator>: iterator is sth that can produce a value, invoke .next() to get the value (not how go works btw)
	/**

	**/
	for l := range c {
		// dont place sleep in main routine
		// function literal; rmb to invoke it 
		// differnet go routines sharing the same l in the function literal 
		// ^would not happen if using named since pass by value
		// referencing l in function literal is dangerous since l might change b4 passed to the checkLink(); due to change of `l` in next for loop 
		// esp when sleep b4 checkLink()
		// share same var cos for loop var is declared outside for loop scope not inside
		// dont do this: 
		// go func () {
		// 	time.Sleep(5 * time.Second)
		// 	checkLink(l, c)
		// }()	
		
		go func (link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)	
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return 
	}	

	fmt.Println(link, "is up!")
	c <- link
}