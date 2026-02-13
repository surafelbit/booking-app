package main
import "fmt"
import "strings"
func main(){

	var conferenceName = "Go Confrenece"
var numberOfTickets = 50
var remainigTickets = 50
var firstName string
var lastName string
var email string
// var bookings= []string{}
var bookings [] string
// bookings[0]=firstName + " " + lastName
for {
	fmt.Printf("Welcome to %v ,User\n", conferenceName)
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets")
	fmt.Scan(&numberOfTickets)
	if(numberOfTickets<remainigTickets){
	
	
	remainigTickets = remainigTickets- numberOfTickets
	fmt.Printf("thank you %v %v for booking %v tickets, and now we have, %v remaining for you. you will receive email on %v\n", firstName,lastName,  numberOfTickets,remainigTickets,  email)
	fmt.Printf("The whole value %v \n", bookings)
bookings = append(bookings, firstName + " "+ lastName)
var firstNames =[]string{}
for _, booking := range bookings{
	var list = strings.Fields(booking)
	 
	firstNames = append(firstNames,list[0])

}
fmt.Printf("The first element %v \n", firstNames)
fmt.Printf("The length of the array %v \n", len(bookings))
} else if numberOfTickets == remainigTickets{
	fmt.Println("you entered the exact so you finished ok man")

	break
} else {
	fmt.Println("error on the loop")
}

}

}