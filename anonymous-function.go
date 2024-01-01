// Anonymous functions
package main

import (
	"fmt"
	"math"
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

	sq := squares()
	fmt.Println(sq())

}

// Return function from a function
func MathExp() func(int, int) int {

	return func(f1 int, f2 int) int {

		return f1 * f2
	}

}

// Stateful functions
func squares() func() int64 {
	x := 1.0
	return func() int64 {
		x += 1
		return int64(math.Pow(x, 2))

	}

}
