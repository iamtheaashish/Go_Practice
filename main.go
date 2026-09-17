package main

import "fmt"

func main() {
	// most used construct in go
	// uninitialized slice is nil
	

	var nums =  make([]int, 2, 5)
	fmt.Println(cap(nums))
	// capacity = maximum numbers of elements can fit.
	// length = number of ACTUAL elements which are there.
	fmt.Println(nums == nil)

	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	fmt.Println(nums)
	fmt.Println(cap(nums))

}