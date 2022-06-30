package main

import "fmt"

func main() {
	fmt.Printf("\n[ Basic Pointer Notation: ] ############################# \n")

	var i int
	i = 42

	var j int
	j = 2701

	var p *int // the type *T is a pointer to a T value (int as a base)
	p = &i     // point to i

	fmt.Printf("Memmory Location of p: %p\n", p)    // the memory address of the value this variable is pointing at
	fmt.Printf("Dereferenced Value of p: %d\n", *p) // the value that this memory address is pointing at (dereferencing)

	*p = 21                                            // update the value that this memory address is pointing at
	fmt.Printf("Updated original variable i: %d\n", i) // the value that this memory address is pointing at

	p = &j                                               // point to j
	*p = *p / 37                                         // divide j through the pointer ==> 2701 / 37
	fmt.Printf("Updated original variable j: %d\n\n", j) // see the new value of j

	fmt.Printf("######################################################### \n")
}
