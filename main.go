package main

import (
	"fmt"
	"slices"
)

func main() {
	// most used construct in go
	// uninitialized slice is nil
	

	// var nums =  make([]int, 2, 5)
	// // fmt.Println(cap(nums))
	// // // capacity = maximum numbers of elements can fit.
	// // // length = number of ACTUAL elements which are there.
	// // fmt.Println(nums == nil)

	// // nums = append(nums, 1)
	// // nums = append(nums, 2)
	// // nums = append(nums, 3)
	// // nums = append(nums, 4)
	// // fmt.Println(nums)
	// // fmt.Println(cap(nums))

	// nums[0] = 3

// 	var nums = make([]int, 0, 5)
// 	nums = append(nums, 2)
// 	var nums2 = make([]int, len(nums))

// 	copy(nums2, nums)

// 	fmt.Println(nums, nums2)
// 	fmt.Println(cap(nums), cap(nums2))
// 	fmt.Println(len(nums), cap(nums2))
// 

// slice operator
	// var nums = []int{1,2,3}

	// fmt.Println(nums[1:2])
	var nums1 = []int{1,2,3}
	var nums2 = []int{1,2,3}

	fmt.Println(slices.Equal(nums1, nums2))

}