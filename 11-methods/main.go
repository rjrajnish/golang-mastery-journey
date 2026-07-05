 package main

import "fmt"

// ----------------------------------------
// Struct Definition
// ----------------------------------------

type BankAccount struct {
	AccountNumber string
	AccountHolder string
	Balance       float64
}

// ----------------------------------------
// Value Receiver
// ----------------------------------------

func (b BankAccount) DisplayAccount() {

	fmt.Println("Account Number :", b.AccountNumber)
	fmt.Println("Account Holder :", b.AccountHolder)
	fmt.Println("Balance        :", b.Balance)

}

// ----------------------------------------
// Pointer Receiver
// ----------------------------------------

func (b *BankAccount) Deposit(amount float64) {

	b.Balance += amount

	fmt.Printf("₹%.2f deposited successfully.\n", amount)

}

func (b *BankAccount) Withdraw(amount float64) {

	if amount > b.Balance {

		fmt.Println("Insufficient Balance")

		return

	}

	b.Balance -= amount

	fmt.Printf("₹%.2f withdrawn successfully.\n", amount)

}

// ----------------------------------------
// Method Returning Value
// ----------------------------------------

func (b BankAccount) GetBalance() float64 {

	return b.Balance

}

func main() {

	fmt.Println("========================================")
	fmt.Println("       Go Methods - Complete Guide")
	fmt.Println("========================================")

	account := BankAccount{
		AccountNumber: "ACC1001",
		AccountHolder: "Rajnish Pandey",
		Balance:       10000,
	}

	// ----------------------------------------
	// Example 1
	// ----------------------------------------

	fmt.Println("\n1. Display Account")

	account.DisplayAccount()

	// ----------------------------------------
	// Example 2
	// ----------------------------------------

	fmt.Println("\n2. Deposit Money")

	account.Deposit(5000)

	account.DisplayAccount()

	// ----------------------------------------
	// Example 3
	// ----------------------------------------

	fmt.Println("\n3. Withdraw Money")

	account.Withdraw(3000)

	account.DisplayAccount()

	// ----------------------------------------
	// Example 4
	// ----------------------------------------

	fmt.Println("\n4. Check Balance")

	fmt.Printf("Current Balance : ₹%.2f\n", account.GetBalance())

	// ----------------------------------------
	// Example 5
	// ----------------------------------------

	fmt.Println("\n5. Withdraw More Than Balance")

	account.Withdraw(50000)

	// ----------------------------------------
	// Example 6
	// ----------------------------------------

	fmt.Println("\n6. Value Receiver Example")

	account.DisplayAccount()

}