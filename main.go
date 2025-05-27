package main

import (
	"bufio"
	"fmt"
	"os"

	"banking-application/bank"
	"banking-application/utils"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n--- Welcome to GoBank CLI ---")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")

		input := utils.ReadLine(scanner)
		switch input {
		case "1":
			bank.RegisterUser(scanner)
		case "2":
			user := bank.LoginUser(scanner)
			if user != nil {
				bank.HandleUserMenu(scanner, user)
			}
		case "3":
			fmt.Println("Thank you. Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
