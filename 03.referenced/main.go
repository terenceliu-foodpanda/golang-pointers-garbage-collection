package main

import (
	"fmt"
)

func main() {
	// Slices
	fmt.Println("\n######################")
	fmt.Println("SLICES")
	fmt.Println("######################")
	var fullSlice []int = []int{1, 2, 3, 4, 5}
	var partialSlice = fullSlice[3:]

	fmt.Printf("Init:\n")
	fmt.Println("-------")
	fmt.Printf("ArrInt: %+v \nSliceInt: %+v\n\n", fullSlice, partialSlice)

	partialSlice[0] = 10
	fmt.Printf("After:\n")
	fmt.Println("-------")
	fmt.Printf("ArrInt: %+v \nSliceInt: %+v\n", fullSlice, partialSlice)

	// MAP
	fmt.Println("\n\n######################")
	fmt.Println("MAP")
	fmt.Println("######################")
	var emptyMap = make(map[string]interface{})
	fmt.Println("Init:")
	fmt.Println("-------")
	fmt.Printf("emptyMap: %+v\n", emptyMap)

	MapFunc(emptyMap)
	fmt.Println("After:")
	fmt.Println("-------")
	fmt.Printf("emptyMap: %+v\n", emptyMap)
}

func MapFunc(val map[string]interface{}) {
	val["this is a new value"] = 100
}

//In this second example, Slices from an array will have the memory location of the array,
//so if we change the value of the Slices it will affect the array value, and map data type is
//Pass by Reference by default, so anything changed inside the function will change the
//original map value. If we want to do a Pass by Value for map we can use copy which will make a
//new variable of the map so the original variable will not take any effect from the method.
//chan or channel is also referenced data type.

//Like many languages, Go has two classes of datatype: those with “value” semantics and
//those with “reference” semantics. Primitive types are all “value” types —
//those include the usual int, string, byte rune, bool types and so on.
//Pointer types are, in effect, “reference” types. However, maps, slices and channels
//are all special types that are or contain references.

//https://neilalexander.dev/2021/08/29/go-pass-by-value
