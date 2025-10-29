package main

import "fmt"

func main() {
	var totalTickets uint = 50
	var bookings []string

	fmt.Println("🎟️  Welcome to GoTicket — The Simplest Booking App Ever!")
	fmt.Println("---------------------------------------------------------")
	fmt.Printf("We’ve got a total of %v tickets available right now!\n\n", totalTickets)

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

	if(userTickets > totalTickets) {
		fmt.Printf("Sorry %v, only %v tickets are left! Try a smaller number.\n", userFirstName, totalTickets)
		return
	}

	totalTickets -= userTickets

	userName := userFirstName + " " + userLastName
	bookings = append(bookings, userName)

	fmt.Printf("\n✅ Thanks %v! Your booking for %v ticket(s) is confirmed.\n", userName, userTickets)
	fmt.Printf("🎟️  A confirmation email has been sent to: %v\n", userEmail)
	fmt.Printf("📉 Tickets remaining: %v\n\n", totalTickets)

}