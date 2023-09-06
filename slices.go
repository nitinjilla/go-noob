//SLICES

package main

import "fmt"

func main() {

	pie := []string{"Zero", "One", "Two"}

	fmt.Printf("Array %s is of type %T and has a length of %d and a capacity of %d. \n", pie, pie, len(pie), cap(pie))

	for _, i := range pie {

		pie = append(pie, i)

		fmt.Printf("Array %s is of type %T and has a length of %d and a capacity of %d. \n", pie, pie, len(pie), cap(pie))

	}

}
