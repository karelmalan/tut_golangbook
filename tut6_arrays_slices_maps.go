// tut6_arrays_slices_maps

package main

import (
	"fmt"
)

func main() {

	fmt.Println("Golang-book tut6 start \n")

	x := [5]int{3, 4, 5, 6, 7}

	// old school -- similar to C:
	for i := 0; i < len(x); i++ {
		fmt.Printf("x[%d] = %d \n", i, x[i])
	}

	// range method (similar to python)
	for index, value := range x {
		fmt.Println("index: ", index)
		fmt.Println("value: ", value)
	}
}
