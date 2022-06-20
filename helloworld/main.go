// package = project = workspace
// one project only one package
/*
package:
dont need import if within same package
declare which package this is under so you can import by typing package.filename


can have many related files inside them
need to declare in all files within this directory
type of packages:
1. executable
generate a file we can run

package main -> go build -> main.exe
main: used to make executable package
2. reusable
ie. package build for other people to use
*non main package
code used as helpers
good place to put reusable logic
when run go build -> nothing
any other name
dependency type package

standard lib packages:
golang.org/pkg
eg. fmt, etc

order:
package declaration
import statement
body: function declaration
*/

package main

// import to get access to fmt package
// standard package
// fmt lib (format): used to print, etc

import "fmt"

// must have main function cos main package
// func is function
func main() {
	fmt.Println("Hi there!")
}


