package main

import (
	"fmt"
)

type Payer interface {
	Pay(amount int) error
}

type CreditCard struct {
	CardNumber int
	Balance    int
}

type Cash struct {
	Money int
}

type CryptoWallet struct {
	Address string
	Balance int
}

func (c *CryptoWallet) Pay(amount int) error {
	if c.Balance >= amount {
		c.Balance -= amount
		fmt.Printf("Crypto transaction of %d completed", amount)
		return nil
	}

	return fmt.Errorf("insufficient crypto balance")
}

func (c *Cash) Pay(amount int) error {
	if c.Money >= amount {
		c.Money -= amount
		fmt.Printf("Cash payment of %d accepted", amount)
		return nil
	}
	return fmt.Errorf("not enough cash")
}

func (p *CreditCard) Pay(amount int) error {
	if amount > p.Balance {
		return fmt.Errorf("insufficient funds on credit card")
	}
	p.Balance -= amount
	fmt.Printf("Card %d: debited %d, remaining balance: %d", p.CardNumber, amount, p.Balance)
	return nil
}

func main() {
	card1 := &CreditCard{CardNumber: 1000200030004000, Balance: 780}
	card2 := &CreditCard{CardNumber: 1111111111111111, Balance: 470}
	card3 := &CreditCard{CardNumber: 2222222222222222, Balance: 4830}
	card4 := &CreditCard{CardNumber: 3333333333333333, Balance: 0}

	cash1 := &Cash{Money: 5000}

	wallet1 := &CryptoWallet{Address: "0900878x5xxvf6", Balance: 321}

	ourPay := []Payer{card1, card2, card3, card4, cash1, wallet1}

	fmt.Println("Processing payments...")

	for _, z := range ourPay {
		err := z.Pay(500)

		if err != nil {
			fmt.Printf(" | Status: Error - %v\n", err)
		} else {
			fmt.Println(" | Status: Success")
		}
	}
}