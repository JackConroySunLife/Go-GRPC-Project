package main

import "fmt"

type BankDetails struct {
	Name   string
	Pin    string
	Active bool
}

func main() {
	apiKey := "key314trwgobjbyyc"

	fmt.Println(apiKey)

	serverUsername := "Goodwill"
	serverPassword := "31tjAW@!@"
	serverKey := "18*(^sg"

	fmt.Println(serverUsername, serverPassword, serverKey)

	cardHolderName := "Jack"
	cardNumber := 374245455400126

	fmt.Println(cardHolderName, cardNumber)

	username := "Ryan"
	email := "ryan@gmail.com"
	password := "12GHKis!"

	fmt.Println(username, email, password)

	usersBankDetails := BankDetails{
		Name:   "Jack",
		Pin:    "2123345678",
		Active: true,
	}

	fmt.Println(usersBankDetails)
}
