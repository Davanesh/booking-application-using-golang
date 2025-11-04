package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

const totalTickets uint = 50
var remainingTickets uint = totalTickets

var bookings []UserData

func main() {
	greetUser()

	for remainingTickets > 0 {
		firstName, lastName, email, ticketType, userTickets := getUser()

		if !isValidName(firstName, lastName) {
			fmt.Println("⚠️  First and last names must each have at least 2 characters.\n")
			continue
		}
		if !isValidEmail(email) {
			fmt.Println("⚠️  Invalid email! Must contain '@' and '.'.\n")
			continue
		}
		if userTickets == 0 || userTickets > remainingTickets {
			fmt.Printf("⚠️  You can book between 1 and %v tickets.\n\n", remainingTickets)
			continue
		}

		user := CreateUserData(firstName, lastName, email, ticketType, userTickets)
		bookings = append(bookings, user)
		remainingTickets -= userTickets

		wg.Add(1)
		go sendTicket(user.NumOfTickets, user.FirstName, user.LastName, user.Email)
		wg.Wait()

		bookingConfirmationMsg(user)

		fmt.Printf("🎫 Remaining tickets: %v\n\n", remainingTickets)
		fmt.Println("--------------------------------------------")

		if remainingTickets == 0 {
			fmt.Println("\n🎉 All tickets sold out! Thanks for booking with GoTicket 💖")
			break
		}
	}

	printSummary()
}

func greetUser() {
	fmt.Println(" \n \n 🎟️  Welcome to GoTicket — The Simplest Booking App Ever!")
	fmt.Println("---------------------------------------------------------")
	fmt.Printf("We’ve got a total of %v tickets available right now!\n\n", totalTickets)
}

func getUser() (string, string, string, string, uint) {
	var userFirstName, userLastName, userEmail, ticketType string
	var userTickets uint

	fmt.Print("👉 Enter your first name: ")
	fmt.Scan(&userFirstName)

	fmt.Print("👉 Enter your last name: ")
	fmt.Scan(&userLastName)

	fmt.Print("📧 Enter your email ID: ")
	fmt.Scan(&userEmail)

	fmt.Println("🎟️ Choose ticket type:")
	fmt.Println("1️⃣ Standard (₹499)")
	fmt.Println("2️⃣ Premium (₹999)")
	fmt.Println("3️⃣ Box (₹1299)")
	fmt.Print("👉 Enter choice: ")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		ticketType = "standard"
	case 2:
		ticketType = "premium"
	case 3:
		ticketType = "box"
	default:
		ticketType = "standard"
	}

	for {
		fmt.Print("🎫 How many tickets do you want to book? ")
		fmt.Scan(&userTickets)
		if userTickets <= remainingTickets {
			break
		}
		fmt.Printf("❌ Only %v tickets left, try again!\n", remainingTickets)
	}
	return userFirstName, userLastName, userEmail, ticketType, userTickets
}

func sendTicket(numOfTickets uint, firstName string, lastName string, userEmail string) {
	time.Sleep(2 * time.Second)
	ticket := fmt.Sprintf("%v ticket(s) for %v %v", numOfTickets, firstName, lastName)
	fmt.Println("---------------------------------------------------------")
	fmt.Printf("📨 Sending ticket: \n%v\nTo email: %v\n", ticket, userEmail)
	fmt.Println("---------------------------------------------------------")
	wg.Done()
}

func isValidEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

func isValidName(firstName, lastName string) bool {
	return len(firstName) >= 2 && len(lastName) >= 2
}

func bookingConfirmationMsg(user UserData) {
	fmt.Printf("\n✅ Booking Confirmed for %v %v!\n", user.FirstName, user.LastName)
	fmt.Printf("🆔 Booking ID: %v\n", user.ID)
	fmt.Printf("🎟️ Ticket Type: %v\n", strings.Title(user.TicketType))
	fmt.Printf("🎫 Quantity: %v\n", user.NumOfTickets)
	fmt.Printf("💰 Total Price: ₹%v\n", user.TotalPrice)
	fmt.Printf("📧 Email: %v\n", user.Email)
	fmt.Printf("🕒 Booked at: %v\n\n", user.BookedAt.Format("02-Jan-2006 15:04:05"))
}

func printSummary() {
	fmt.Println("\n📋 Booking Summary:")
	fmt.Println("---------------------------------------------------------")
	for i, booking := range bookings {
		fmt.Printf("%d️⃣ %v %v — %v ticket(s) | %v | ₹%v | ID: %v\n",
		i+1,
		booking.FirstName,
		booking.LastName,
		booking.NumOfTickets,
		strings.Title(booking.TicketType),
		booking.TotalPrice,
		booking.ID)
	}
	fmt.Println("---------------------------------------------------------")
	fmt.Println("💖 Thanks for using GoTicket! Enjoy the show!")
}
