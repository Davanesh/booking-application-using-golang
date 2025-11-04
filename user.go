package main

import (
	"fmt"
	"math/rand"
	"time"
)

type UserData struct {
	ID           string
	FirstName    string
	LastName     string
	Email        string
	NumOfTickets uint
	TicketType   string
	TotalPrice   uint32
	BookedAt     time.Time
}

func CreateUserData(
	firstName string,
	lastName string,
	email string,
	ticketType string,
	numOfTickets uint,
) UserData {
	var price uint32
	switch ticketType {
	case "box":
		price = 1299
	case "premium":
		price = 999
	case "standard":
		price = 499
	default:
		price = 499
	}
	totalPrice := numOfTickets * uint(price)

	return UserData{
		ID:           fmt.Sprintf("GT%04d", rand.Intn(10000)),
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		NumOfTickets: numOfTickets,
		TicketType:   ticketType,
		TotalPrice:   uint32(totalPrice),
		BookedAt:     time.Now(),
	}
}
