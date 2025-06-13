package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PaymentProcessor interface {
	ProcessPayment(amount float64) (float64, float64, error)
	Name() string
}

type PaymentStats struct {
	count       int
	totalAmount float64
	totalFee    float64
}

type PaymentSystem struct {
	processors map[int]PaymentProcessor
	stats      map[string]*PaymentStats
	balance    float64
}

type CreditCard struct {
	balance float64
}

func (c *CreditCard) ProcessPayment(amount float64) (float64, float64, error) {
	if amount > 50000 || amount < 10 {
		fmt.Println("invalid amount")
		os.Exit(1)
	}
	if amount > c.balance {
		fmt.Println("not enough funds")
		os.Exit(1)
	}
	fee := amount * 0.015
	total := fee + amount
	if total > c.balance {
		fmt.Println("not enough for fee")
		os.Exit(1)
	}
	c.balance = c.balance - total
	return amount, fee, nil
}

func (c *CreditCard) Name() string { return "CreditCard" }

type PayPal struct {
	balance float64
}

func (p *PayPal) ProcessPayment(amount float64) (float64, float64, error) {
	if amount < 10 || amount > 50000 {
		fmt.Println("invalid amount")
		os.Exit(1)
	}
	if p.balance < amount {
		fmt.Println("not enough funds")
		os.Exit(1)
	}
	fee := amount * 0.035
	if p.balance < (amount + fee) {
		fmt.Println("not enough for fee")
		os.Exit(1)
	}
	p.balance -= (amount + fee)
	return amount, fee, nil
}

func (p *PayPal) Name() string { return "PayPal" }

type Cash struct {
	balance float64
}

func (c *Cash) ProcessPayment(amount float64) (float64, float64, error) {
	if amount > 50000 || amount < 10 {
		fmt.Println("invalid amount")
		os.Exit(1)
	}
	if c.balance < amount {
		fmt.Println("not enough cash")
		os.Exit(1)
	}
	c.balance -= amount
	return amount, 0, nil
}

func (c *Cash) Name() string { return "Cash" }

type BankTransfer struct {
	balance float64
}

func (b *BankTransfer) ProcessPayment(amount float64) (float64, float64, error) {
	if amount < 10 || amount > 50000 {
		fmt.Println("invalid amount")
		os.Exit(1)
	}
	if amount > b.balance {
		fmt.Println("not enough funds")
		os.Exit(1)
	}
	fee := amount * 0.02
	total := amount + fee
	if total > b.balance {
		fmt.Println("not enough for fee")
		os.Exit(1)
	}
	b.balance -= total
	return amount, fee, nil
}

func (b *BankTransfer) Name() string { return "BankTransfer" }

func NewPaymentSystem() *PaymentSystem {
	return &PaymentSystem{
		processors: map[int]PaymentProcessor{
			4: &BankTransfer{balance: 15000},
			3: &Cash{balance: 2000},
			2: &PayPal{balance: 5000},
			1: &CreditCard{balance: 10000},
		},
		stats:   make(map[string]*PaymentStats),
		balance: 100,
	}
}

func (ps *PaymentSystem) DisplayMenu() {
	fmt.Println("Payment System")
	fmt.Println("Methods:")
	for i := 4; i >= 1; i-- {
		fmt.Printf("%d. %s\n", i, ps.processors[i].Name())
	}
}

func (ps *PaymentSystem) ProcessPayment(processorID int, amount float64) error {
	processor, exists := ps.processors[processorID]
	if !exists {
		fmt.Println("unknown method")
		os.Exit(1)
	}

	fmt.Printf("Processing %.2f...\n", amount)
	fmt.Printf("Method: %s\n", processor.Name())

	amountPaid, fee, err := processor.ProcessPayment(amount)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if ps.stats[processor.Name()] == nil {
		ps.stats[processor.Name()] = &PaymentStats{}
	}
	ps.stats[processor.Name()].totalFee += fee
	ps.stats[processor.Name()].count++
	ps.stats[processor.Name()].totalAmount += amountPaid

	fmt.Printf("Success!\nAmount: %.2f\nFee: %.2f (%.1f%%)\nTotal: %.2f\n", amountPaid, fee, (fee/amountPaid)*100, amountPaid+fee)
	return nil
}

func (ps *PaymentSystem) DisplayStats() {
	fmt.Println("Stats:")
	for name, stats := range ps.stats {
		fmt.Printf("%s\nTransactions: %d\nAmount: %.2f\nFee: %.2f\n", name, stats.count, stats.totalAmount, stats.totalFee)
	}
}

func (ps *PaymentSystem) RefillBalance(processorID int, amount float64) error {
	if processorID > 4 || processorID < 1 {
		fmt.Println("unknown method")
		os.Exit(1)
	}
	if amount <= 0 {
		fmt.Println("invalid refill amount")
		os.Exit(1)
	}

	if processorID == 1 {
		if creditCard, ok := ps.processors[processorID].(*CreditCard); ok {
			creditCard.balance += amount
		}
	}
	if processorID == 2 {
		if payPal, ok := ps.processors[processorID].(*PayPal); ok {
			payPal.balance += amount
		}
	}
	if processorID == 3 {
		if cash, ok := ps.processors[processorID].(*Cash); ok {
			cash.balance += amount
		}
	}
	if processorID == 4 {
		if bankTransfer, ok := ps.processors[processorID].(*BankTransfer); ok {
			bankTransfer.balance += amount
		}
	}

	fmt.Printf("Refilled %.2f for %s\n", amount, ps.processors[processorID].Name())
	return nil
}

func main() {
	ps := NewPaymentSystem()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		ps.DisplayMenu()
		fmt.Print("Choose (1-4 pay, 5 refill, 6 stats, 7 exit): ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Enter number!")
			continue
		}

		if choice == 7 {
			fmt.Println("Exit")
			break
		}

		if choice == 5 {
			fmt.Print("Method (1-4): ")
			scanner.Scan()
			processorID, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("Enter number!")
				continue
			}
			fmt.Print("Refill amount: ")
			scanner.Scan()
			amount, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
			if err != nil {
				fmt.Println("Enter valid amount!")
				continue
			}
			if err := ps.RefillBalance(processorID, amount); err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			continue
		}

		if choice == 6 {
			ps.DisplayStats()
			continue
		}

		if choice > 4 || choice < 1 {
			fmt.Println("Wrong choice!")
			continue
		}

		fmt.Print("Payment amount: ")
		scanner.Scan()
		amount, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
		if err != nil {
			fmt.Println("Enter valid amount!")
			continue
		}

		if err := ps.ProcessPayment(choice, amount); err != nil {
			continue
		}
	}
}
