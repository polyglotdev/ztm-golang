package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	elijah := Passenger{
		Name:         "Elijah",
		TicketNumber: 1,
		Boarded:      true,
	}
	fmt.Println(elijah)

	var (
		ezra      = Passenger{Name: "Ezra", TicketNumber: 2, Boarded: true}
		stephanie = Passenger{Name: "Stephanie", TicketNumber: 3, Boarded: true}
	)

	fmt.Println(ezra)
	fmt.Println(stephanie)

	var littleDom Passenger
	littleDom.Name = "Dominic"
	littleDom.TicketNumber = 4
	littleDom.Boarded = true
	fmt.Println(littleDom)

	bus := Bus{FrontSeat: elijah}
	fmt.Println(bus.FrontSeat.Name)
}
