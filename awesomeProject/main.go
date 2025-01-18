package main

import "fmt"

//type Node struct {
//	N    int
//	Next *Node

//type DNode struct {
//	N    int
//	Next *DNode
//	Prev *DNode
//}

//type Queue struct {
//	s []any
//}
//
//func (q *Queue) Push(v any) {
//	q.s = append(q.s, v)
//}
//
//func (q *Queue) Pop() (any, bool) {
//	if len(q.s) == 0 {
//
//		return nil, false
//	}
//	res := q.s[0]
//
//	q.s = q.s[1:]
//
//	return res, true
//}

//func reverse(n *Node)*Node{
//	//has to reverse linked list
//
//	// 1->2->3
//	//3->2->1
//}

//type Node struct {
//	N    int
//	Next *Node
//}

//func makeCircle(head *Node) *Node {
//		temp := head
//	for head.Next!= nil{
//		fmt.Println(head.N)
//		head = head.Next
//	}
//		head.Next = temp
//
//		return temp
//	// makes circle out of linked list
//}

//
//func printLinkedListReversed(head *Node) {
//	for head.Next != nil{
//		fmt.Println(head.N)
//		head = head.Next
//	}
//
//
//	//4->3->2->1
//}

//func findMiddle(head *Node) int {
//	slow := head
//	fast := head.Next
//	for fast != nil && fast.Next != nil {
//		slow = slow.Next
//		fast = fast.Next.Next
//	}
//	return slow.N
//}

// finds n/2+1 value

// 3->2->1 WITH LOOP
//}

// 1 - 2 - 3 - 3 - 2 - 1
//func printDlinkedList(head *DNode) {
//	for head.Next != nil {
//		fmt.Println(head.N)
//		head = head.Next
//	}
//	fmt.Println(head.N)

//for head != nil {
//	fmt.Println(head.N)
//	head = head.Prev
//}
//1->2->3->3->2->1
//}

//func PrintLinkedListReversed(node *Node) {
//	if node == nil {
//		return
//	}
//	PrintLinkedListReversed(node.Next)
//	fmt.Println(node.N)
//}

func main() {
	s := `
	This string is on
	multiple lines
	within three single
	quotes on either side`
	fmt.Println(s)
}

//	s := Queue{}
//
//	s.Push(1)
//	s.Push(2)
//	s.Push(3)
//
//	fmt.Println(s.Pop())
//	fmt.Println(s.Pop())
//	fmt.Println(s.Pop())
//	fmt.Println(s.Pop())
//}

//p := DNode{
//	N:    1,
//	Next: nil,
//	Prev: nil,
//}
//l := DNode{
//	N:    2,
//	Next: nil,
//	Prev: &p,
//}
//k := DNode{
//	N:    3,
//	Next: nil,
//	Prev: &l,
//}
//l.Next = &k
//p.Next = &l
//
//printDlinkedList(&p)

//
//	n := Node{
//		N: 1,
//		Next: &Node{
//			N: 2,
//			Next: &Node{
//				N: 3,
//				Next: &Node{
//					N:    4,
//					Next: nil,
//				},
//			},
//		},
//	}
//	fmt.Println(&n)
//
//}

//
////import (
////	"errors"
////	"fmt"
////)
////
////type WalletInterface interface {
////	Deposit(amount float64) error
////	Withdraw(amount float64) error
////	Transfer(to WalletInterface, amount float64) error
////}
////type USDWallet struct {
////	Currency string
////	Balance  float64
////}
////type GelWallet struct {
////	Currency string
////	Balance  float64
////}
////type BTCWallet struct {
////	Currency string
////	Balance  float64
////}
////
////func (u *USDWallet) Deposit(amount float64) error {
////	if amount <= 0 {
////		return errors.New("amount must be greater than zero")
////	}
////	u.Balance += amount
////	return nil
////}
////
////func (u *USDWallet) Withdraw(amount float64) error {
////	if amount > u.Balance {
////		return error(errors.New("amount must be less than balance"))
////	}
////	if amount < 0 {
////		return fmt.Errorf("amount must be greater to zero")
////	}
////	u.Balance -= amount
////	return nil
////}
////func (u *USDWallet) Transfer(to *GelWallet, amount float64) error {
////	if amount > u.Balance {
////		return errors.New("no balance to transfer")
////	} else if amount <= 0 {
////		return errors.New("amount must be greater than zero")
////	}
////	if u.Currency != to.Currency {
////		return errors.New("currency not supported")
////	}
////	u.Balance -= amount
////	to.Balance += amount
////
////	return nil
////}
////
////func (u *USDWallet) Transfer2(to *BTCWallet, amount float64) error {
////	if amount > u.Balance {
////		return errors.New("no balance to transfer")
////	} else if amount <= 0 {
////		return errors.New("amount must be greater than zero")
////	}
////	if u.Currency != to.Currency {
////		return errors.New("currency not supported")
////	}
////	u.Balance -= amount
////	to.Balance += amount
////
////	return nil
////}
////
////func main() {
////	u := &USDWallet{
////		Currency: "USD",
////		Balance:  100,
////	}
////	err := u.Deposit(500)
////	if err != nil {
////		fmt.Println("Deposir failed:", err)
////	}
////	fmt.Println(*u)
////
////	errWithd := u.Withdraw(300)
////	if errWithd != nil {
////		fmt.Println("Withdraw failed:", errWithd)
////	}
////	fmt.Println(*u)
////
////	g := &GelWallet{
////		Currency: "GEL",
////		Balance:  200,
////	}
////	errTrans := u.Transfer(g, 200)
////	if errTrans != nil {
////		fmt.Println("Transfer failed:", errTrans)
////	}
////
////	fmt.Println(*g)
////
////	b := &BTCWallet{
////		Currency: "USD",
////		Balance:  300,
////	}
////	fmt.Println("BTC balance", *b)
////
////	errTrans2 := u.Transfer2(b, 100)
////	if errTrans2 != nil {
////		fmt.Println("Transfer2 failed:", errTrans2)
////	}
////	fmt.Println("BTC balance", *b)
////}
