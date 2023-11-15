package main

import "fmt"

func Greeting(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

func AnyMessage(message string) string {
	return fmt.Sprintf("Message: %s", message)
}

func AddThreeNumbers(a, b, c int) int {
	return a + b + c
}

func AddNumbers(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func AddNumbersWithMessage(message string, numbers ...int) int {
	fmt.Println(message)
	return AddNumbers(numbers...)
}

func main() {
	fmt.Println(Greeting("World"))
	fmt.Println(AnyMessage("Hello World"))
	fmt.Println(AddThreeNumbers(1, 2, 3))
	fmt.Println(AddNumbers(1, 2, 3, 4, 5))
	fmt.Println(AddNumbersWithMessage("Sum of 1, 2, 3, 4, 5", 1, 2, 3, 4, 5))
}
