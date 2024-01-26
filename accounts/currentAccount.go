package accounts

import (
	"accountBank/client"
	"math"
)

type Current struct {
	User    client.User
	balance float64
}

func (c *Current) Withdrawal(value float64) (string, bool) {
	if value > 0 && value <= c.balance {
		c.balance -= value
		return "Saque realizado com sucesso", true
	} else {
		return "Valor insuficiente para saque", false
	}
}

func (c *Current) Deposit(value float64) (bool, float64) {
	if value > 0 {
		c.balance += value
		return true, c.balance
	} else {
		return false, c.balance
	}
}

// transfer money entry accounts
func (c *Current) Transfer(otherAccount Current, value float64) bool {
	if value > 0 {
		_, status := otherAccount.Withdrawal(value)
		if status {
			c.Deposit(value)
			return true
		}
	}
	return false
}

// NewCurrent creates a new current account with the given user and initial balance

// get value balance
func (c *Current) Balance() float64 {
	return math.Round(c.balance*100) / 100
}
