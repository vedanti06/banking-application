package bank

import "sync"

type User struct {
	Username string
	Password string
	Balance  float64
	mu       sync.Mutex
}

type Bank struct {
	Users map[string]*User
	mu    sync.RWMutex
}

var App = Bank{
	Users: make(map[string]*User),
}
