package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)
	// you can declare multiple variables at oce
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// variables declared without initialization are zero
	var e int
	fmt.Println(e)
	// := is shortand for initializing a variable
	f := "apple"
	fmt.Println(f)
}
