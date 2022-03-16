package main

import (
	"fmt"
)

func main() {
	var applicationName = "Booking application"
	const conferenceTickets int = 50
	var remainingTickets int = 50

	fmt.Printf("Initializing: %v \n", applicationName)
	fmt.Printf("Total tickets: %v \n", conferenceTickets)
	fmt.Printf("Remaining tickets: %v \n", remainingTickets)

	// user information
	var userName string
	var userTickets int

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User: %v -> Tickets: %v \n", userName, userTickets)

}
