package main

import (
	"fmt"
	"sync"
	"time"
)

var remainingTickets uint = 50
var bookings = []UserData{}

// var bookings = make([]map[string]string, 0)

type UserData struct {
	firstName      string
	lastName       string
	email          string
	numberOfTckets uint
}

var wg = sync.WaitGroup{}

func main() {

	// for {
	var firstName string
	var lastName string
	var userTickets uint

	fmt.Scan(&firstName)
	fmt.Scan(&lastName)
	fmt.Scan(&userTickets)

	if userTickets <= remainingTickets {
		bookTickets(userTickets, firstName, lastName)

		wg.Add(1) //Synchornize main exit
		go sendTickets(userTickets, firstName, lastName, "email@go.com")

		var firstNames []string
		for _, booking := range bookings {
			firstNames = append(firstNames, booking.firstName)
		}
		fmt.Printf("The first names of bookings are: %v\n", firstNames)
	} else {
		fmt.Printf("We only have %v tickets left, booking not possible", remainingTickets)
	}
	// }
	wg.Wait()
}

func bookTickets(userTickets uint, firstName string, lastName string) {
	remainingTickets = remainingTickets - userTickets
	var userData = UserData{
		firstName:      firstName,
		lastName:       lastName,
		email:          "email@go.com",
		numberOfTckets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v, for booking %v tickets. Confirmation sent to %v. ", firstName, userTickets, "email@go.com")
	fmt.Printf("%v tickets remainging now.\n", remainingTickets)
}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var tickets = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("\n###################")
	fmt.Printf("Sending tickets:\n %v \nto email address %v\n", tickets, email)
	fmt.Println("###################\n")
	wg.Done()
}
