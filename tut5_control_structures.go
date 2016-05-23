// tut5_control_structures
package main

import (
	"fmt"
)

func main() {
	fmt.Println("golang-book.com tutorial 5 problems...")

	fmt.Println("\nProblem 1")
	for i := 1; i < 101; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("\n\n Problem 2")
	for i := 1; i < 101; i++ {
		if i%3 == 0 {
			fmt.Print("Fizz")
		}
		if i%5 == 0 {
			fmt.Print("Buzz")
		}
		if (i%3 != 0) && (i%5 != 0) {
			fmt.Print(i)
		}

		fmt.Println("")
	}

}
