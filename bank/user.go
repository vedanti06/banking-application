package bank

import (
	"bufio"
	"fmt"
	"banking-application/utils"
)

func HandleUserMenu(scanner *bufio.Scanner, user *User) {
	for {
		fmt.Printf("\n--- Welcome, %s ---\n", user.Username)
		fmt.Println("1. View Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Transfer")
		fmt.Println("5. Logout")
		fmt.Print("Choose an option: ")

		switch utils.ReadLine(scanner) {
		case "1":
			user.mu.Lock()
			fmt.Printf("Balance: $%.2f\n", user.Balance)
			user.mu.Unlock()
		case "2":
			fmt.Print("Enter amount to deposit: $")
			amount := utils.ReadFloat(scanner)
			if amount <= 0 {
				fmt.Println("Amount must be positive.")
				continue
			}
			user.mu.Lock()
			user.Balance += amount
			user.mu.Unlock()
			fmt.Println("Deposit successful.")
		case "3":
			fmt.Print("Enter amount to withdraw: $")
			amount := utils.ReadFloat(scanner)
			user.mu.Lock()
			if amount <= 0 {
				fmt.Println("Invalid amount.")
			} else if amount > user.Balance {
				fmt.Println("Insufficient funds.")
			} else {
				user.Balance -= amount
				fmt.Println("Withdrawal successful.")
			}
			user.mu.Unlock()
		case "4":
			fmt.Print("Enter recipient username: ")
			recipientName := utils.ReadLine(scanner)

			App.mu.RLock()
			recipient, exists := App.Users[recipientName]
			App.mu.RUnlock()
			if !exists {
				fmt.Println("Recipient not found.")
				continue
			}

			fmt.Print("Enter amount to transfer: $")
			amount := utils.ReadFloat(scanner)
			if amount <= 0 {
				fmt.Println("Invalid amount.")
				continue
			}

			go func(sender, receiver *User, amt float64) {
				sender.mu.Lock()
				defer sender.mu.Unlock()

				if amt > sender.Balance {
					fmt.Println("Transfer failed: insufficient balance.")
					return
				}
				sender.Balance -= amt

				receiver.mu.Lock()
				receiver.Balance += amt
				receiver.mu.Unlock()

				fmt.Printf("Transferred $%.2f to %s\n", amt, receiver.Username)
			}(user, recipient, amount)

		case "5":
			fmt.Println("Logged out.")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}
