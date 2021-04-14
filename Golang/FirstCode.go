package main

import (
	"fmt"
)


func main()  {
	var (
		name[2]string
		age[2]int
	)

	for i:=0 ;  i<2; i++ {
		fmt.Println("Enter the name:")
		fmt.Scanln(&name[i])

		fmt.Println("Enter the Age:")
		fmt.Scanln(&age[i])
	}

	fmt.Println(name)
	fmt.Println(age)
}