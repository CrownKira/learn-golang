package main

import "fmt"

type contactInfo struct {
	email string 
	zipCode int
}

// type: type declaration 
// struct {}: the type value
type person struct {
	firstName string
	lastName string
	// contactInfo (alone): === contactInfo contactInfo
	contactInfo
}

func main() {
	// rely on order of the field 
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)
	// zero values: 
	/**
	string: ""
	int: 0
	float: 0 
	bool: false (not undefined)
	**/
	// var alex person 

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// /**
	// Printf()

	// %+v: print all field names and values in struct
	// **/
	// fmt.Printf("%+v", alex)


	jim := person {
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo {
			email: "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// jimPointer := &jim
	// jim can act like &jim: syntactic sugar
	jim.updateName("jimmy")
	jim.print()
}


/**
go: pass by value
struct: pass by value
slice: pass by value also (but it is the slice data structure that got copied) (slice data structure contains the address to the underlying array which also got copied but not the array itself)
make a new copy of struct when pass to a function 

&variable: give me the memory address
*pointer give me the value at this memory address

*person: is a type for person address, address type 
*pointerToPerson: convert address to value
**/

/**
array vs slice:
slice: used 99% 
slice data structure contains: ptr to head, capacity, length
mySlice: pointing to slice data structure not the underlying array
wrapper data structure
**/
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// can declare receiver method for struct also
func (p person) print() {
	fmt.Printf("%+v", p)
}


/**
map: hash, object, dict 
map: key value pairs 
key: same type 
value: same type
**/