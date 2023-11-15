package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 11
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1
	shoppingList["eggs"] += 1

	fmt.Println(shoppingList)
	delete(shoppingList, "milk")
	fmt.Println("Milk deleted, new list:", shoppingList)
	fmt.Println("need", shoppingList["eggs"], "eggs")

	cereal, found := shoppingList["cereal"]
	if found {
		fmt.Println("Found cereal:", cereal)
	} else {
		fmt.Println("No cereal found")
	}

	totalItems := 0
	for _, count := range shoppingList {
		totalItems += count
	}
	fmt.Println("Total items:", totalItems)
}
