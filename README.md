# CLI Banking Application in Go

A command-line banking application written in Go that supports user authentication, concurrent transactions, and basic banking operations with input validation and transaction logic.

## Features

- User registration and login with password authentication
- Concurrent-safe deposit, withdrawal, and transfer operations using mutexes
- Input validation for numeric and string inputs
- Transaction logic ensuring consistency and atomic updates
- Simple CLI interface for user interaction
- Organized code structure with concurrency-safe data models

## Architecture & Concurrency

- `Bank` struct holds all users in a concurrent-safe map protected by a `sync.RWMutex`
- Each `User` struct has its own `sync.Mutex` to protect its balance from race conditions during concurrent operations
- User input handled using `bufio.Scanner` with utility functions for clean input reading and parsing

## Getting Started

1. **Clone the repository**

   ```bash
   git clone https://github.com/yourusername/banking-application.git
   cd banking-application
   go mod tidy
   go run .
