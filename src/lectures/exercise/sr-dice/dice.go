//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these circumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// roll simulates a dice roll with the given number of sides.
func roll(sides int) int {
	// rand will return a number between 0 and sides.
	// We want a number between 1 and sides, so we add 1.
	return rand.Intn(sides) + 1
}

// isSnakeEyes takes a sum as an integer and dice as an integer
// and evaluates if the sum is 2 and the dice is 2.
func isSnakeEyes(sum, dice int) bool {
	return sum == 2 && dice == 2
}

// isLuckySeven takes a sum as an integer and evaluates if the sum is 7.
func isLuckySeven(sum int) bool {
	return sum == 7
}

// isEven takes a sum as an integer and evaluates if the sum is even.
func isEven(sum int) bool {
	return sum%2 == 0
}

// isOdd takes a sum as an integer and evaluates if the sum is odd.
func isOdd(sum int) bool {
	return sum%2 == 1
}

// printResult prints the result of a dice roll.
func printResult(r, d, rolled, sum, dice int) {
	fmt.Printf("Roll %d, die %d: %d\n", r, d, rolled)
	fmt.Printf("Sum of roll %d: %d\n", r, sum)
	if isSnakeEyes(sum, dice) {
		fmt.Println("Snake eyes!")
	}
	if isLuckySeven(sum) {
		fmt.Println("Lucky 7!")
	}
	if isEven(sum) {
		fmt.Println("Even!")
	}
	if isOdd(sum) {
		fmt.Println("Odd!")
	}
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	dice, sides := 2, 12
	rolls := 2

	for r := 1; r <= rolls; r++ {
		sum := 0
		for d := 1; d <= dice; d++ {
			rolled := roll(sides)
			sum += rolled
			printResult(r, d, rolled, sum, dice)
		}
	}
}
