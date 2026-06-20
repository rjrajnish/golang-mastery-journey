# Functions in Go

## What are Functions?

Functions are reusable blocks of code that perform a specific task.

Instead of writing the same code multiple times, we place the logic inside a function and call it whenever needed.

Example:

```go
func greet() {
	fmt.Println("Hello Go")
}
```

---

## Why Do We Need Functions?

Functions help us:

* Reuse code
* Reduce duplication
* Improve readability
* Simplify maintenance
* Organize business logic

Without functions, programs quickly become difficult to maintain.

---

## When Should We Use Functions?

Use functions when:

* Logic is repeated
* Code becomes lengthy
* Business logic needs separation
* A task can be isolated

Examples:

* Login validation
* User registration
* Tax calculation
* Database queries
* Payment processing

---

## Where Are Functions Used?

Functions are used everywhere in software development.

### REST APIs

```go
func GetUsers() {}
```

### Authentication

```go
func Login() {}
```

### Database Layer

```go
func CreateUser() {}
```

### Payment Gateway

```go
func ProcessPayment() {}
```

### Email Service

```go
func SendEmail() {}
```

Functions are the building blocks of every Go application.

---

# Function Syntax

Basic structure:

```go
func functionName(parameters) returnType {
	// logic
}
```

Example:

```go
func add(a int, b int) int {
	return a + b
}
```

---

# Function with No Parameters

```go
func greet() {
	fmt.Println("Hello")
}
```

Usage:

```go
greet()
```

---

# Function with Parameters

```go
func greet(name string) {
	fmt.Println("Hello", name)
}
```

Usage:

```go
greet("Rajnish")
```

Output:

```text
Hello Rajnish
```

---

# Function with Return Value

```go
func add(a int, b int) int {
	return a + b
}
```

Usage:

```go
result := add(10, 20)
```

Output:

```text
30
```

---

# Multiple Return Values

Go allows multiple return values.

Example:

```go
func divide(a float64, b float64) (float64, error) {

	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}
```

Usage:

```go
result, err := divide(10, 2)
```

This feature is heavily used in Go backend applications.

---

# Variadic Functions

Variadic functions accept unlimited arguments.

Example:

```go
func sum(numbers ...int) int {
	return 0
}
```

Usage:

```go
sum(10, 20)
sum(10, 20, 30)
sum(10, 20, 30, 40)
```

Common use cases:

* Logging
* Aggregations
* Dynamic calculations

---

# Anonymous Functions

Functions without names.

```go
func() {
	fmt.Println("Anonymous Function")
}()
```

Used in:

* Goroutines
* Middleware
* Callbacks

---

# JavaScript vs Go

## JavaScript

```javascript
function add(a, b) {
	return a + b;
}
```

## Go

```go
func add(a int, b int) int {
	return a + b
}
```

---

# Key Differences

| Feature          | JavaScript | Go           |
| ---------------- | ---------- | ------------ |
| Types            | Dynamic    | Static       |
| Return Type      | Optional   | Required     |
| Multiple Returns | No         | Yes          |
| Performance      | Runtime    | Compile Time |

Go functions are more predictable because all types are known during compilation.

---

# Real-World Example

Tax Calculator

```go
func calculateTax(amount float64) float64 {
	return amount * 0.18
}
```

Usage:

```go
tax := calculateTax(1000)
```

Output:

```text
180
```

---

# Mini Project

Build a Calculator.

Functions:

```go
func add()
func subtract()
func multiply()
func divide()
```

Take input from variables and display results.

---

# Advantages of Functions

* Reusability
* Better organization
* Easier testing
* Cleaner code
* Improved readability

---

# Key Takeaways

* Functions are reusable blocks of code.
* Functions can accept parameters.
* Functions can return values.
* Go supports multiple return values.
* Variadic functions accept unlimited arguments.
* Functions are the foundation of Go applications.
