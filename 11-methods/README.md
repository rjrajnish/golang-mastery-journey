 # Methods in Go

Methods are one of the most important features of Go. They allow you to attach behavior to a struct, making your code more organized, reusable, and easier to understand.

If a struct represents **data**, then methods represent **actions** that can be performed on that data.

For example:

- A `BankAccount` can deposit money.
- A `User` can log in.
- A `Product` can calculate its discount.

These actions are implemented using methods.

---

# Table of Contents

1. What is a Method?
2. Why Do We Need Methods?
3. Method Syntax
4. Value Receivers
5. Pointer Receivers
6. Methods Returning Values
7. Methods vs Functions
8. JavaScript vs Go
9. Memory Representation
10. Real-World Use Cases
11. Best Practices
12. Common Mistakes
13. Mini Project
14. Interview Questions
15. Key Takeaways

---

# What is a Method?

A method is a function that is associated with a specific type (usually a struct).

Unlike regular functions, methods belong to a struct and can access its fields directly.

Example:

```go
func (b BankAccount) DisplayAccount() {

	fmt.Println(b.AccountHolder)

}
```

The part before the method name:

```go
(b BankAccount)
```

is called the **receiver**.

---

# Why Do We Need Methods?

Imagine a banking application.

Without methods:

```go
DisplayAccount(account)

Deposit(account)

Withdraw(account)
```

These are just separate functions.

With methods:

```go
account.DisplayAccount()

account.Deposit(1000)

account.Withdraw(500)
```

The code is more readable because the behavior belongs to the object itself.

---

# Method Syntax

```go
func (receiver Type) MethodName() {

}
```

Example:

```go
func (b BankAccount) DisplayAccount() {

	fmt.Println(b.Balance)

}
```

---

# Value Receiver

A value receiver works on a **copy** of the struct.

```go
func (b BankAccount) DisplayAccount() {

	fmt.Println(b.Balance)

}
```

Changes inside this method do not affect the original struct.

Use value receivers when:

- You only need to read data.
- The struct is small.
- No modifications are required.

---

# Pointer Receiver

A pointer receiver works on the original struct.

```go
func (b *BankAccount) Deposit(amount float64) {

	b.Balance += amount

}
```

Changes are reflected in the original object.

Use pointer receivers when:

- Updating data
- Working with large structs
- Avoiding unnecessary copies

---

# Methods Returning Values

Methods can return values.

```go
func (b BankAccount) GetBalance() float64 {

	return b.Balance

}
```

Usage:

```go
balance := account.GetBalance()
```

---

# Methods vs Functions

| Function | Method |
|----------|---------|
| Independent | Belongs to a type |
| Called directly | Called using an object |
| Doesn't have a receiver | Has a receiver |

Function:

```go
CalculateTax(amount)
```

Method:

```go
account.Deposit(1000)
```

---

# JavaScript vs Go

## JavaScript

```javascript
class BankAccount {

	constructor(balance){

		this.balance = balance;

	}

	deposit(amount){

		this.balance += amount;

	}

}
```

---

## Go

```go
type BankAccount struct {

	Balance float64

}

func (b *BankAccount) Deposit(amount float64){

	b.Balance += amount

}
```

### Comparison

| JavaScript | Go |
|------------|----|
| Class | Struct |
| Class Method | Method |
| `this` keyword | Receiver |
| Inheritance | Composition |
| Dynamic | Static |

Go keeps things simpler by separating data (structs) from behavior (methods).

---

# Memory Representation

```
BankAccount

+------------------------+

Account Holder

Balance

+------------------------+

        │

        ▼

Deposit()

Withdraw()

DisplayAccount()

GetBalance()
```

The methods operate on the struct through the receiver.

---

# Real-World Use Cases

Methods are used in almost every Go project.

Examples:

- User Authentication
- Product Pricing
- Shopping Cart
- Bank Account
- Employee Management
- Order Processing
- Payment Systems
- Inventory Management

---

# Best Practices

- Use methods for behavior related to a struct.
- Use value receivers for read-only operations.
- Use pointer receivers for updates.
- Keep methods focused on one responsibility.
- Choose receiver types consistently across methods.

---

# Common Mistakes

## Using Value Receiver When Updating Data

Incorrect:

```go
func (b BankAccount) Deposit(amount float64){

	b.Balance += amount

}
```

This updates only the copy.

Correct:

```go
func (b *BankAccount) Deposit(amount float64){

	b.Balance += amount

}
```

---

## Mixing Receiver Types

Avoid using value receivers for some methods and pointer receivers for others without a reason. Consistency makes your API easier to understand.

---

# Interview Questions

### What is a method?

A method is a function with a receiver.

---

### What is a receiver?

A receiver is the type a method belongs to.

---

### Difference between value receiver and pointer receiver?

Value receiver works on a copy.

Pointer receiver works on the original object.

---

### When should you use pointer receivers?

- Updating data
- Large structs
- Better performance

---

### Can methods return values?

Yes.

Methods can return any type, including multiple values.

---

# Key Takeaways

- Methods define behavior for structs.
- Methods are functions with receivers.
- Value receivers work on copies.
- Pointer receivers modify original data.
- Methods improve code organization and readability.
- Methods are used extensively in backend development.

---

 