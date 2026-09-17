package main

import "fmt"

func main() {
	// most used construct in go
	// uninitialized slice is nil
	var nums []int

	fmt.Println(nums == nil)

	fmt.Println(len(nums))

	

}