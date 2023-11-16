package main

import "fmt"

type Counter struct {
	hits int
}

func increment(c *Counter) {
	c.hits++
	fmt.Println("Hits:", c.hits)
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)

}

func main() {
	counter := Counter{}
	hello := "Hello"
	world := "World"
	fmt.Println(hello, world)

	replace(&hello, "Hi", &counter)
	fmt.Println(hello, world)

	phrase := []string{hello, world}
	fmt.Println(phrase)

	replace(&phrase[0], "Go!", &counter)
	fmt.Println(phrase)
}
