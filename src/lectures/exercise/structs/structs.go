//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing a length and width field
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Rectangle struct {
	Length float64
	Width  float64
}

func Area(r Rectangle) float64 {
	return r.Length * r.Width
}

func Perimeter(r Rectangle) float64 {
	return 2*r.Length + 2*r.Width
}

func DoubleSize(r Rectangle) Rectangle {
	r.Length *= 2
	r.Width *= 2
	return r
}

func main() {
	r := Rectangle{Length: 3, Width: 7}
	fmt.Println("Area:", Area(r))

	double := DoubleSize(r)
	fmt.Println("Area:", Area(double))
}
