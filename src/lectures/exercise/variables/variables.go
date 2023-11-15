// Summary:
//
//	Print basic information to the terminal using various variable
//	creation techniques. The information may be printed using any
//	formatting you like.
package main

import "fmt"

func main() {
	var fc = "blue"
	by, age := 1983, 39
	fi := "D"
	li := "H"

	fmt.Println("My name is", fi, li, "and I am", age, "years old.")
	fmt.Println("I was born in", by, "and my favorite color is", fc)
}
