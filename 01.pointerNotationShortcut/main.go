package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {

	fmt.Printf("\n[ Pointer Notation with Shortcut: ] #################### \n")

	fmt.Printf("Initialize struct:\n")
	v := Vertex{1, 2}
	fmt.Println(v)

	fmt.Printf("After update:\n")
	p := &v

	//Version1: Original
	//(*p).X = 1e9

	//Version2: Shortcut
	p.X = 1e9 //same as the line above

	fmt.Println(v)
	fmt.Printf("######################################################### \n")

}
