package main

import "fmt"

func main() {

	// defineing and assigning a string
	var str string = "Hello, world"

	//Comppiler automatically deciding variable type
	s := "This is a string"

	//here %v is used to print value and %T for the Type of variable
	//And to use %v and %T we use Printf like in C not Print / Println
	
	fmt.Printf("%v, %T", str, str)

	fmt.Printf("%v, %T", s, s)

    fmt.Println("Hello, World")
}
