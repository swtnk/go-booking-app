package main

import "fmt"

func main() {

	var conferenceName string = "Go Conference"
	const totalTickets uint = 100
	var remainingTickets uint = totalTickets

	fmt.Printf("Welcome to %v ticket booking system.\n", conferenceName)
	fmt.Printf("Hurry up! only %v tickets are left.\n", remainingTickets)

	var firstName string

	fmt.Printf("Enter First Name: ")
	fmt.Scan(&firstName)

	fmt.Printf("FN: %v\n", firstName)

}
