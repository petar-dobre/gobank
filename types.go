package main

import "math/rand"

type Account struct {
	ID int
	FirstName string
	LastName string
	number int64
	balance int64
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID: rand.Intn(10000),
		FirstName: firstName,
		LastName: lastName,
		number: int64(rand.Intn(100000)),
	}
}