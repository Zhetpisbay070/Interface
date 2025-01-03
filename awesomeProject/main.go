package main

import "fmt"

type DNode struct {
	N    int
	Next *DNode
	Prev *DNode
}

func printDlinkedList(head *DNode) {
	for head != nil {
		fmt.Println(head.N)
		head = head.Next

	}
	// 1->2->3->3->2->1
}

func main() {
	p := DNode{
		N:    1,
		Next: nil,
		Prev: nil,
	}
	l := DNode{
		N:    2,
		Next: nil,
		Prev: &p,
	}
	k := DNode{
		N:    3,
		Next: nil,
		Prev: &l,
	}
	l.Next = &k
	p.Next = &l
	printDlinkedList(&p)
}

//import (
//	"errors"
//	"fmt"
//)
//
//type WalletInterface interface {
//	Deposit(amount float64) error
//	Withdraw(amount float64) error
//	Transfer(to WalletInterface, amount float64) error
//}
//type USDWallet struct {
//	Currency string
//	Balance  float64
//}
//type GelWallet struct {
//	Currency string
//	Balance  float64
//}
//type BTCWallet struct {
//	Currency string
//	Balance  float64
//}
//
//func (u *USDWallet) Deposit(amount float64) error {
//	if amount <= 0 {
//		return errors.New("amount must be greater than zero")
//	}
//	u.Balance += amount
//	return nil
//}
//
//func (u *USDWallet) Withdraw(amount float64) error {
//	if amount > u.Balance {
//		return error(errors.New("amount must be less than balance"))
//	}
//	if amount < 0 {
//		return fmt.Errorf("amount must be greater to zero")
//	}
//	u.Balance -= amount
//	return nil
//}
//func (u *USDWallet) Transfer(to *GelWallet, amount float64) error {
//	if amount > u.Balance {
//		return errors.New("no balance to transfer")
//	} else if amount <= 0 {
//		return errors.New("amount must be greater than zero")
//	}
//	if u.Currency != to.Currency {
//		return errors.New("currency not supported")
//	}
//	u.Balance -= amount
//	to.Balance += amount
//
//	return nil
//}
//
//func (u *USDWallet) Transfer2(to *BTCWallet, amount float64) error {
//	if amount > u.Balance {
//		return errors.New("no balance to transfer")
//	} else if amount <= 0 {
//		return errors.New("amount must be greater than zero")
//	}
//	if u.Currency != to.Currency {
//		return errors.New("currency not supported")
//	}
//	u.Balance -= amount
//	to.Balance += amount
//
//	return nil
//}
//
//func main() {
//	u := &USDWallet{
//		Currency: "USD",
//		Balance:  100,
//	}
//	err := u.Deposit(500)
//	if err != nil {
//		fmt.Println("Deposir failed:", err)
//	}
//	fmt.Println(*u)
//
//	errWithd := u.Withdraw(300)
//	if errWithd != nil {
//		fmt.Println("Withdraw failed:", errWithd)
//	}
//	fmt.Println(*u)
//
//	g := &GelWallet{
//		Currency: "GEL",
//		Balance:  200,
//	}
//	errTrans := u.Transfer(g, 200)
//	if errTrans != nil {
//		fmt.Println("Transfer failed:", errTrans)
//	}
//
//	fmt.Println(*g)
//
//	b := &BTCWallet{
//		Currency: "USD",
//		Balance:  300,
//	}
//	fmt.Println("BTC balance", *b)
//
//	errTrans2 := u.Transfer2(b, 100)
//	if errTrans2 != nil {
//		fmt.Println("Transfer2 failed:", errTrans2)
//	}
//	fmt.Println("BTC balance", *b)
//}
