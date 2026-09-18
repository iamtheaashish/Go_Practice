package main

import (
	"fmt"
	"maps"
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
	// var nums1 = []int{1,2,3}
	// var nums2 = []int{1,2,3}

	// fmt.Println(slices.Equal(nums1, nums2))

	// var nilMap map[string]int
	// fmt.Println(nilMap)

	// totalWins := map[string]int{}
	// fmt.Println(totalWins)

	// teams := map[string][]string {
	// 	"Orcas": []string{"Fred", "Ralph", "Bijou"},
	// 	"Lions": []string{"Sarah", "Peter", "Billie"},
	// 	"Kittens": []string{"Waldo", "Raul", "Ze"},
	// }

	// fmt.Println(teams)

	// totalWins := map[string]int{}
	// 	totalWins["Orcas"] = 1
	// 	totalWins["Lions"] = 2
	// 	fmt.Println(totalWins["Orcas"])
	// 	fmt.Println(totalWins["Kittens"])
	// 	totalWins["Kittens"]++

	// 	fmt.Println(totalWins["Kittens"])
	// 	totalWins["Lions"] = 3
	// 	fmt.Println(totalWins["Lions"])

	// m := map[string]int{
	// 	"hello" : 5,
	// 	"world" : 0,
	// }

	// v, ok := m["hello"]
	// fmt.Println(v, ok)
	
	// v, ok = m["world"]
	// fmt.Println(v, ok)

	// v, ok = m["goodbye"]
	// fmt.Println(v, ok)

	// m := map[string]int{
	// 	"hello" : 5,
	// 	"world" : 10,
	// }

	// fmt.Println(m, len(m))

	// delete(m, "hello")

	// fmt.Println(m)

	// clear(m)

	// fmt.Println(m, len(m))

	// comparing maps
	m := map[string]int{
		"hello" : 5,
		"world" : 10,
	}
	n := map[string]int{
		"world" : 5,
		"hello" : 10,
	}

	fmt.Println(maps.Equal(m, n))

}