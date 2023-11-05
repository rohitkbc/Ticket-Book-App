package main

import (
	"fmt"
	"strconv"
	"time"
	"sync"
)

var conferenceName = "Go Conference"
const conferenceTickets = 50
var remainingTickets uint = 50
//var bookings []map[string]string				// This is also same as line no 12
var bookings =  make([]map[string]string, 0)
//var bookings =  make([]userData,0)			// This is list of userData type 

type userData struct {
	firstName string
	lastName string
	email string
	userTickets uint
}

var wg = sync.WaitGroup{}

func main(){
	greetUsers()

	for { 
		firstName, lastName, email, userTickets := getUserInput()	

		var isValidName, isValidEmail, isValidTicketsNum bool = validateUserInput(firstName, lastName, email, userTickets)

		if  isValidName && isValidEmail && isValidTicketsNum {
			
			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := printFirstNames()
			fmt.Printf("The firstnames of bookings are %v\n", firstNames)

			var noTicketsRemaining bool = remainingTickets == 0

			if noTicketsRemaining {
				fmt.Printf("No tickets are available for %v. Come back next year\n", conferenceName)
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			
			if !isValidEmail {
				fmt.Println("Please enter valid email address")
			}
			
			if !isValidTicketsNum {
				fmt.Println("Please enter valid tickets numbers")
			}
		}
		wg.Wait()
	}
	
}

func greetUsers(){
	fmt.Printf("Welcome to %v ticket booking app\n", conferenceName)
	fmt.Println("We have total", conferenceTickets, "tickets and", remainingTickets, "are still available for booking")
	fmt.Println("Get your tickets here to attend")
}

func printFirstNames() []string{
	firstNames := []string{}

	for _, booking := range bookings{
		//var names = strings.Fields(booking)
		//var firstName = names[0]
		firstNames = append(firstNames, booking["firstName"] )
	}
	return firstNames
}

func getUserInput()  (string, string, string, uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("How many tickets do you want?: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string){
	remainingTickets = remainingTickets - userTickets

	// var userDataStruct = userData{
	// 	firstName: firstName,
	// 	lastName: lastName,
	// 	email: email,
	// 	userTickets: userTickets,
	// }

	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["userTickets"] = strconv.FormatUint(uint64(userTickets),10)

	bookings = append(bookings, userData)

	fmt.Printf("List of Bookings: %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v %v tickets. You will receive confirmation email on %v soon\n", firstName, lastName, conferenceName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string){
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###################################")
	fmt.Printf("Sending ticket: \n%v \nto email address %v\n", ticket, email)
	fmt.Println("###################################")
	wg.Done()
}