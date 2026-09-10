package main

import (
	"fmt"
	"math"
)

func main(){
	var b byte = math.MaxUint8
	var smallI int32 = math.MaxInt32
	var bigI uint64 = math.MaxUint64

	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println((bigI))
}
