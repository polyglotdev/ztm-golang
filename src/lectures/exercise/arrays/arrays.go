//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	Name  string
	Price float64
}

// getTotal takes in an array and returns the total price
// of all the products in the array
func getTotal(products [4]Product) float64 {
	var total float64
	for _, product := range products {
		total += product.Price
	}
	return total
}

func main() {
	var shoppingList [4]Product
	shoppingList[0] = Product{Name: "Milk", Price: 2.99}
	shoppingList[1] = Product{Name: "Eggs", Price: 1.99}
	shoppingList[2] = Product{Name: "Bread", Price: 1.99}
	shoppingList[3] = Product{Name: "Butter"}
	li := len(shoppingList) - 1
	ti := len(shoppingList)
	total := getTotal(shoppingList)
	fmt.Println("Last item:", shoppingList[li].Name)
	fmt.Println("Total items:", ti)
	fmt.Println("Total cost:$", fmt.Sprintf("%.2f", total))
}
