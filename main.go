package main

import (
	"fmt"
	"strings"
)

func main() {
	var totalTickets uint = 50
	var bookings []string
	for {
		greetuser(totalTickets)

		userFirstName, userLastName, userEmail, userTickets := getUser()

		isVaildName := len(userFirstName) >= 2 && len(userLastName) >= 2
		isValidEmail := strings.Contains(userEmail, "@")
		isVaildTicket := userTickets > 0 && userTickets <= totalTickets

		if isVaildName && isValidEmail && isVaildTicket {
			totalTickets -= userTickets

			userName := userFirstName + " " + userLastName
			bookings = append(bookings, userName)

			printFirstNames(bookings)

			fmt.Printf("\n✅ Thanks %v! Your booking for %v ticket(s) is confirmed.\n", userName, userTickets)
			fmt.Printf("🎟️  A confirmation email has been sent to: %v\n", userEmail)
			fmt.Printf("📉 Tickets remaining: %v\n\n", totalTickets)
			fmt.Printf("Current bookings: %v\n", bookings)
			if totalTickets == 0 {
				fmt.Println("😞 Sorry, all tickets are sold out!")
				return 
			}
		} else {
			if !isVaildName {
				fmt.Println("⚠️  First name and last name must be at least 2 characters long.")
			} 
			if !isValidEmail {
				fmt.Println("⚠️  Email address must contain an '@' symbol.")
			}
			if !isVaildTicket {
				fmt.Println("⚠️  Number of tickets must be greater than 0 and less than or equal to the available tickets.")
			}
		}
	}
}

func greetuser(totalTickets uint) {
	fmt.Println("🎟️  Welcome to GoTicket — The Simplest Booking App Ever!")
	fmt.Println("---------------------------------------------------------")
	fmt.Printf("We’ve got a total of %v tickets available right now!\n\n", totalTickets)
}

func getUser() (string, string, string, uint) {
	var userFirstName, userLastName, userEmail string
	var userTickets uint

	fmt.Print("👉 Enter your first name: ")
	fmt.Scan(&userFirstName)

	fmt.Print("👉 Enter your last name: ")
	fmt.Scan(&userLastName)

	fmt.Print("📧 Enter your email ID: ")
	fmt.Scan(&userEmail)

	fmt.Print("🎫 How many tickets do you want to book? ")
	fmt.Scan(&userTickets)

	return userFirstName, userLastName, userEmail, userTickets
}

func printFirstNames(bookings []string) {
	firstName := []string{}
	for _, booking := range bookings {
		var parts = strings.Fields(booking)
		firstName = append(firstName, parts[0])
		fmt.Print(firstName)
	}
}