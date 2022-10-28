package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "ICMRACPC"
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to %v booking and applications\n", conferenceName)
	fmt.Printf("a total of %v tickets are available.\n", remainingTickets)

	for {
		var firstName string
		var lastName string
		var userTickets uint

		fmt.Scan(&firstName)
		fmt.Scan(&lastName)
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets left, booking not possible", remainingTickets)
			continue
		}

		bookings = append(bookings, firstName+" "+lastName)
		remainingTickets -= userTickets

		var firstNames []string
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		var noTicketsRemaining bool = remainingTickets <= 0
		if noTicketsRemaining {
			break
		}
	}

}
