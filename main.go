package main

import (
	"fmt"
)

func main() {
	// arrays
	// var nums [4]int
	// nums[0] = 1
	// fmt.Println(nums)

	var vals [4]bool
	vals[2] = true
	vals[1] = true
	fmt.Println(vals)

	var name [3]string
	name[0] = "golang"
	fmt.Println(name)

	nums := [3]int{1,2,3}
	fmt.Println(nums)
}