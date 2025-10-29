package main

import "fmt"

func main() {
	var totalTickets int = 50
	var userFirstName string
	var userLastName string
	var userEmail string
	var userTickets int

	fmt.Println("Welcome to booking application")
	fmt.Println("You can book tickets here")
	fmt.Printf("Total tickets available: %v\n", totalTickets)

	fmt.Print("Enter your first name: ")
	fmt.Scan(&userFirstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&userLastName)
	var userName string = userFirstName + " " + userLastName 

	fmt.Print("Enter your email ID name: ")
	fmt.Scan(&userEmail)
	
	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	totalTickets -= userTickets
	fmt.Printf("Thanks %v, Your ticket has been sent to %v :)", userName, userEmail)
}