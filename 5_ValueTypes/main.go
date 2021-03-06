package main

import "fmt"

type StructVal struct {
	IntVal int
}

// Pass By Value
//• Pass by value will pass the value of the variable into the method,
//or we can say that the original variable ‘copy’ the value into another memory location
//and pass the newly created one into the method.
//So, anything that happens to the variable inside the method will not affect the original variable value.
func (s StructVal) Add() {
	fmt.Println("######################")
	fmt.Printf("Func Add\n")
	fmt.Println("######################")
	fmt.Printf("Memory Location %p\n", &s)
	fmt.Printf("Value Before: %+v\n", s)
	s.IntVal++
	fmt.Printf("Value After: %+v\n\n", s)
}

// Pass By Reference
//• Pass by reference will pass the memory location instead of the value.
//In other words, it passes the ‘container’ of the variable to the method so,
//anything that happens to the variable inside the method will affect the original variable.
func (s *StructVal) AddPtr() {
	fmt.Println("######################")
	fmt.Printf("Func AddPtr\n")
	fmt.Println("######################")
	fmt.Printf("Memory Location %p\n", s)
	fmt.Printf("Value Before: %+v\n", s)
	s.IntVal++ // Need to use a pointer here inorder to update!
	//(*s).IntVal++ //SAME
	fmt.Printf("Value After: %+v\n\n", s)
}

func main() {
	init := StructVal{
		IntVal: 0,
	}

	fmt.Println("\n######################")
	fmt.Printf("MAIN\n")
	fmt.Println("######################")
	fmt.Printf("Memory Location: %p\n", &init)
	fmt.Printf("Value: %+v\n\n", init)

	init.Add()
	fmt.Println("######################")
	fmt.Printf("AFTER Func Add()\n")
	fmt.Println("######################")
	fmt.Printf("Value: %+v\n\n", init)

	init.AddPtr()
	fmt.Println("######################")
	fmt.Printf("AFTER Func AddPtr()\n")
	fmt.Println("######################")
	fmt.Printf("Value: %+v\n\n", init)

	fmt.Println("######################")
	fmt.Printf("FINAL\n")
	fmt.Println("######################")
	fmt.Printf("Value: %+v\n", init)

}
