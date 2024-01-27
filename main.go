package main

import (
	"fmt"
	"strings"
)

func main() {
	// := this can only be used with variables not const or constants
	const conferenceName string = "Go conference"
	const conferenceTickets int8 = 50
	var remainingTickets uint8 = 50

	fmt.Printf("conferenceTicket is %T, remainingTicket is %T , conferenceName is %T \n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking applicationv\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("You can book tickets here")

	for {

		// user inputs

		var firstName string
		var lastName string
		var email string
		var userTickets uint8
		// var bookings [50]string //array
		bookings := []string{} //slice

		fmt.Println("please enter your First Name :")
		fmt.Scan(&firstName) //pointers has to use to take userinput in golang

		fmt.Println("Please enter your last Name")
		fmt.Scan(&lastName)
		fmt.Println("Please enter your email for contact")
		fmt.Scan(&email)
		fmt.Println("Please enter how much tickets you want to book?")
		fmt.Scan(&userTickets)

		isValidname := len(firstName) >= 2 && len(lastName) >= 2
		isValidemail := strings.Contains(email, "@")
		isvalidtickets := userTickets > 0 && userTickets <= remainingTickets

		// check for if remaining tickets is greater than user requested value
		if isValidname && isValidemail && isvalidtickets {

			remainingTickets = remainingTickets - userTickets
			// bookings[0] = firstName + " " + lastName // for array
			bookings = append(bookings, firstName+" "+lastName)
			fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)

			firstNames := []string{}
			// _ is a blank identifier
			for _, booking := range bookings {
				var names = strings.Fields(booking)

				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("Some of other guests names are %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out")
				break
			}
		} else if userTickets == remainingTickets {
			fmt.Printf("%v cannot book same amount of tickets %v as we currenly have, please leave some for other users \n", firstName, remainingTickets)
			continue
		} else {

			fmt.Println("Invalid input detected")
			fmt.Printf("Invalid input , We only have %v tickets left please choose a number matched our tickets count \n", remainingTickets)
			continue
		}
	}

}
