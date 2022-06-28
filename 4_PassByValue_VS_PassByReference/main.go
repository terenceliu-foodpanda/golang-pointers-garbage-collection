//https://david-yappeter.medium.com/golang-pass-by-value-vs-pass-by-reference-e48aac8b2716
package main

import "fmt"

func main() {
	a, b := 0, 0

	// Initialize Value
	fmt.Printf("\n[ 1 - ORIGINAL] ######################################### \n")
	fmt.Printf("[Memory Location] a: %p, b: %p\n", &a, &b)
	fmt.Printf("[Value] a: %d, b: %d\n", a, b) // 0 0
	fmt.Printf("######################################################### \n")

	// Passing By Value a(int)
	Add(a) // Golang will copy value of 'a' and insert it into argument

	// Passing By Reference b(int), &b(*int) => with '&' we can get the memory location of 'b'
	AddPtr(&b) // Pass Memory Location of 'b' into argument

	fmt.Printf("\n[ 4 - FINAL ] ############################################# \n")
	fmt.Printf("[Memory Location] a: %p, b: %p\n", &a, &b)
	fmt.Printf("[Value] a: %d, b: %d\n", a, b) // 0 1
	fmt.Printf("######################################################### \n")
}

// Pass By Value - will pass the value of the variable into the method,
//or we can say that the original variable ‘copy’ the value into another memory location
//and pass the newly created one into the method.
//So, anything that happens to the variable inside the method will not affect the original variable value.
func Add(x int) {
	fmt.Printf("\n[ 2 - PASS BY VALUE ==> 'Add(a)' ] ######################\n")
	fmt.Printf("Before Add = [Memory Location]: %p, [Value]: %d\n", &x, x)
	x++
	fmt.Printf("After Add = [Memory Location]: %p, [Value]: %d\n", &x, x)
	fmt.Printf("######################################################### \n")
}

// Pass By Reference - will pass the memory location instead of the value.
//In other words, it passes the ‘container’ of the variable to the method so,
//anything that happens to the variable inside the method WILL affect the original variable.
func AddPtr(x *int) {
	fmt.Printf("\n[ 3 - PASS BY REFERENCE ==> 'AddPtr(&b)' ] ##############\n")
	fmt.Printf("Before AddPtr, Memory Location: %p, Value: %d\n", x, *x)
	*x++ // We add * in front of the variable because it is a pointer, * will call value of a pointer
	fmt.Printf("After AddPtr, Memory Location: %p, Value: %d\n", x, *x)
	fmt.Printf("######################################################### \n")
}

//In the first example, we will work on the basic data types. As we can see we have 2 functions:
//
//Add(x int) which take an integer as the parameter.
//AddPtr(x *int) which takes a pointer integer as the parameter.
//The memory location of a and b cannot be predicted but we can track them down by printing their memory location.
//
//In the Add method, the memory location of the value is not the same as a from the main() because Go copy the value of a and initialize a new memory location, so if we change the value x++ , a will still be 0 . The final output of a is 0 because it is Pass by value.
//
//In the AddPtr method, the memory location of the value is the same as b so we know that it shares the same ‘container’, everything that we do to x inside AddPtr will affect b value, we try to add the value of x in *x++ . The final output of b is 1 which changed because of Pass by reference.
//
//Other basic data types like int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, string, bool, byte, rune, Array, Structs . Array and Struct have the same property with basic data types.
