package main

import (
	"fmt"
)

func main(){
	var x = [3]int{10, 20, 30}
	var y = [...]int{10, 20, 30}

	fmt.Println(x != y)
}
