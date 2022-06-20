package main

import "fmt"

func main(){
	numbers := []int{0,1,2,3,4,5,6,7,8,9,10}

	for i := range numbers {
		var x string
		if i % 2 == 0 {
			x = "even"
		} else {
			x = "odd"
		}
		fmt.Println(i, x)
	}
}