package main

import "fmt"

func main() {
	fmt.Printf("\n[ Basic Pointer Notation: ] ############################# \n")

	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(p)  // the memory address of the value this variable is pointing at
	fmt.Println(*p) // the value that this memory address is pointing at

	*p = 21        // update the value that this memory address is pointing at
	fmt.Println(i) // the value that this memory address is pointing at

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer ==> 2701 / 37
	fmt.Println(j) // see the new value of j

	fmt.Printf("######################################################### \n")
}
