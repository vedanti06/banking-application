package bank

import (
	"bufio"
	"fmt"
	"banking-application/utils"
)

func RegisterUser(scanner *bufio.Scanner) {
	fmt.Print("Enter username: ")
	username := utils.ReadLine(scanner)
	fmt.Print("Enter password: ")
	password := utils.ReadLine(scanner)

	App.mu.Lock()
	defer App.mu.Unlock()

	if _, exists := App.Users[username]; exists {
		fmt.Println("Username already exists.")
		return
	}

	App.Users[username] = &User{
		Username: username,
		Password: password,
		Balance:  0,
	}
	fmt.Println("Registration successful.")
}

func LoginUser(scanner *bufio.Scanner) *User {
	fmt.Print("Enter username: ")
	username := utils.ReadLine(scanner)
	fmt.Print("Enter password: ")
	password := utils.ReadLine(scanner)

	App.mu.RLock()
	defer App.mu.RUnlock()

	user, exists := App.Users[username]
	if !exists || user.Password != password {
		fmt.Println("Login failed.")
		return nil
	}

	fmt.Println("Login successful.")
	return user
}
