package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/**
type ReadCloser interface {
	Reader
	Closer
}

if you need to satisfy ReadCloser then must also satisfy Closer
**/

type logWriter struct {}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// make() take type of a slice, 99999: number of elements we want the slice to be init with
	// read function is not set up to auto resize the slice if it is already full, so inti with arbitrary large size 

	// bs := make([]byte, 99999) 
	// resp.Body.Read(bs)

	// // can think of byte slice as string
	// fmt.Println(string(bs))


	// []byte -> writer -> source output (eg. outgoing http request, text file on hard drive, terminal, etc) (send to some output channel)
	// Copy(writer, reader) (take from some outside information and copy it to some outside channel)
	// os.Stdout: implement the writer interface (write to some channel)
	// read from some channel, then write to some other channel
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}


/**
Reader interface: 
no matter what the source of input => Reader => []byte (output data that anyone can work with)
every single source of data: going to implement the reader interface

request body implement the reader 
take raw body of respone and inject and push it to byteslice
int: number of bytes read into the slice

**/


func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))	
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}