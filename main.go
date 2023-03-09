package main

import "fmt"

func main() {

	var conferenceName string = "Go Conference"
	const totalTickets uint = 100
	var remainingTickets uint = totalTickets

	fmt.Printf("Welcome to %v ticket booking system.\n", conferenceName)
	fmt.Printf("Hurry up! only %v tickets are left.\n", remainingTickets)

	for {

		if remainingTickets == 0 {
			break
		}

		var firstName string
		var lastName string
		var conferenceTickets uint
		var email string

		fmt.Printf("Enter First Name: ")
		fmt.Scan(&firstName)
		fmt.Printf("Enter Last Name: ")
		fmt.Scan(&lastName)
		fmt.Printf("Enter Email: ")
		fmt.Scan(&email)
		fmt.Printf("Enter number of ticket(s): ")
		fmt.Scan(&conferenceTickets)

		for conferenceTickets > remainingTickets {
			fmt.Printf("Only %v tickets are available.\nEnter number of ticket(s): ", remainingTickets)
			fmt.Scan(&conferenceTickets)
		}

		fmt.Printf("\nThanks %v for booking %v ticket(s).\n", firstName+" "+lastName, conferenceTickets)
		fmt.Printf("Ticket details has been sent to %v\n", email)

		remainingTickets -= conferenceTickets

		fmt.Printf("\n%v ticket(s) left.\n\n", remainingTickets)

	}

}
