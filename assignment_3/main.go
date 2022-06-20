/**
https://www.udemy.com/course/go-the-complete-developers-guide/learn/practice/9958/introduction#notes
read content from a file
then print to the terminal

file to open is provided as an arg
go run main.go myfile.txt
myfile.txt: an arg
**/

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// os.Args is a slice of arguments passed in
	// os.Args: [temp directory of the compiled file, ...args]
	// os.Args[1]: first arg
	// fmt.Println(os.Args)
	f, err := os.Open(os.Args[1])
	// *File implements the method Read, not File
	// https://stackoverflow.com/questions/46306888/implementing-interface-in-golang-gives-method-has-pointer-receiver
	// 	func (f *File) Read(b []byte) (n int, err error)
	// g := *f
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}