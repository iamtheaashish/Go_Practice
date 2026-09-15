package main

import "fmt"

func main() {
	// simple switch
	i := 5

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("Other")
	}

}