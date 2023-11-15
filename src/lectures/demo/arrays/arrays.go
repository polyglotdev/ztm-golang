package main

import "fmt"

type Room struct {
	Name    string
	Cleaned bool
}

func checkCleanliness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.Cleaned {
			fmt.Println(room.Name, "is clean")
		} else {
			fmt.Println(room.Name, "is dirty")
		}
	}
}

func main() {
	rooms := [...]Room{
		{Name: "Office"},
		{Name: "Warehouse"},
		{Name: "Reception"},
		{Name: "Ops"},
	}

	checkCleanliness(rooms)
	fmt.Println("Perform cleaning...")
	rooms[2].Cleaned = true
	rooms[3].Cleaned = true

	checkCleanliness(rooms)
}
