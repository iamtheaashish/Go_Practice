package main

import "fmt"

func main() {
	age := 16
	
	if age >= 18 {
		fmt.Println("Person is an adult.")
	} else if age >= 12 {
		fmt.Println("Person is a teen")
	} else if age <= 11 {
		fmt.Println("Person is a kid")
	}

}