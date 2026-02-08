package main
import "fmt"
func main(){

	var conferenceName = "Go Confrenece"
var numberOfTickets = 50
var remainigTickets = 50
var firstName string
var lastName string
var email string
fmt.Printf("Welcome to %v ,User\n", conferenceName)
fmt.Println("Enter your first name")
fmt.Scan(&firstName)
fmt.Println("Enter your last name")
fmt.Scan(&lastName)
fmt.Println("Enter your email")
fmt.Scan(&email)
fmt.Println("Enter number of tickets")
fmt.Scan(&numberOfTickets)
remainigTickets = remainigTickets- numberOfTickets
fmt.Printf("thank you %v %v for booking %v tickets, and now we have, %v remaining for you. you will receive email on %v\n", firstName,lastName,  numberOfTickets,remainigTickets,  email)
var bookings = [50]string{"steve","rogers","captain"}
}