// Summary:
//
//	Print basic information to the terminal using various variable
//	creation techniques. The information may be printed using any
//	formatting you like.
package main

import "fmt"

const (
	firstInitial string = "D"
	lastInitial  string = "H"
)

func main() {
	var fc string = "Blue"
	birthYear, age := 1983, 39
	ageInDays := age * 365
	fmt.Println("My name is", firstInitial, lastInitial, "and I am", age, "years old.")
	fmt.Println("I was born in", birthYear, "and my favorite color is", fc)
	fmt.Println("I am", ageInDays, "days old.")
}
