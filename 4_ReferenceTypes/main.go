package main

import (
	"fmt"
)

func main() {
	// SLICE
	fmt.Println("\n######################")
	fmt.Println("SLICES")
	fmt.Println("######################")
	var fullSlice = []int{1, 2, 3, 4, 5}

	fmt.Printf("\nInit:\n")
	fmt.Println("-------")
	fmt.Printf("ArrInt: %+v\n\n", fullSlice)
	updateSlice(fullSlice)
	fmt.Printf("After:\n")
	fmt.Println("-------")
	fmt.Printf("ArrInt: %+v", fullSlice)

	// MAP - same effect as a SLICE shown above
	fmt.Println("\n\n######################")
	fmt.Println("MAP")
	fmt.Println("######################")
	var emptyMap = make(map[string]interface{})
	fmt.Println("\nInit:")
	fmt.Println("-------")
	fmt.Printf("emptyMap: %+v\n\n", emptyMap)

	updateMap(emptyMap)
	fmt.Println("After:")
	fmt.Println("-------")
	fmt.Printf("emptyMap: %+v\n", emptyMap)
}

func updateSlice(s []int) { //no pointer was used here and the slice still updated!
	s[0] = 6
}

func updateMap(val map[string]interface{}) { //no pointer was used here and the map still updated!
	val["this is a new value"] = 100
}
