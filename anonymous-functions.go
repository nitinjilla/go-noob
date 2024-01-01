// Anonymous functions
package main

import (
	"fmt"
)

func main() {

	//Below example shows how functions can be variables
	af := func() {
		fmt.Println("Calling Stepmom from func af")
	}

	//Passing parameters
	sm := func(name string) {
		fmt.Printf("Calling %s from func sm\n", name)
	}
	af()
	sm("Stepmom")

	//Return function from a function

	product := MathExp()
	fmt.Println(product(10, 10))

}

//Return function from a function
func MathExp() func(int, int) int {

	return func(f1 int, f2 int) int {

		return f1 * f2
	}

}
