package Struct

import (
	"fmt"
)

type Bank struct {
	cap     int
	balance []int64
}

func ConstructorBank(balance []int64) Bank {
	return Bank{
		cap:     len(balance),
		balance: balance,
	}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if !this.checkAccount(account1) || !this.checkAccount(account2) || !this.checkMoney(account1, money) {
		return false
	}
	if this.minus(account1, money) {
		this.add(account2, money)
	}
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if !this.checkAccount(account) {
		return false
	}
	this.add(account, money)
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if !this.checkAccount(account) {
		return false
	}
	return this.minus(account, money)
}

func (this *Bank) checkAccount(account int) bool {
	if account > this.cap || account < 1 {
		fmt.Println("account: ", account)
		return false
	}
	return true
}
func (this *Bank) checkMoney(account int, money int64) bool {
	return this.balance[account-1] >= money
}

func (this *Bank) add(account int, money int64) {
	this.balance[account-1] += money
}

func (this *Bank) minus(account int, money int64) bool {
	if !this.checkMoney(account, money) {
		fmt.Println("account: money", account, money)
		return false
	}
	this.balance[account-1] -= money
	return true
}
