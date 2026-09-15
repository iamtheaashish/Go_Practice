package main

import "fmt"

const age = 20
func main() {
	const name string = "golang"
	
	fmt.Print(age)

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port)
	fmt.Println(host)

}