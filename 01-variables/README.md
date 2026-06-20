# Variables in Go

## What are Variables?

Variables are named storage locations used to hold data in memory.

A variable allows a program to store, update, and retrieve information during execution.

Example:

```go
name := "Rajnish"
age := 25
```

---

## Why Do We Use Variables?

Variables help us:

* Store user information
* Store calculation results
* Manage application state
* Pass data between functions

Without variables, we would need to repeatedly write the same values throughout the program.

---

## When Should We Use Variables?

Use variables whenever data needs to:

* Change during program execution
* Be reused multiple times
* Be passed to functions
* Be stored temporarily

Examples:

* User names
* Product prices
* API responses
* Database records

---

## Where Are Variables Used in Real Projects?

Variables are used everywhere.

Examples:

### Authentication System

```go
username := "rajnish"
password := "secret"
```

### E-Commerce Application

```go
productName := "Laptop"
price := 59999.99
```

### REST API

```go
requestID := "REQ123"
```

---

## How Do Variables Work?

Step 1:

Declare a variable.

```go
var name string
```

Step 2:

Assign a value.

```go
name = "Rajnish"
```

Step 3:

Use the value.

```go
fmt.Println(name)
```

---

## Variable Declaration Methods

### Standard Declaration

```go
var name string = "Rajnish"
```

### Type Inference

```go
var name = "Rajnish"
```

### Short Declaration

```go
name := "Rajnish"
```

This is the most commonly used approach in Go.

---

## Basic Example

```go
package main

import "fmt"

func main() {
	name := "Rajnish"
	age := 25

	fmt.Println(name)
	fmt.Println(age)
}
```

---

## JavaScript vs Go

### JavaScript

```javascript
let name = "Rajnish";
let age = 25;
```

### Go

```go
name := "Rajnish"
age := 25
```

Difference:

* JavaScript is dynamically typed.
* Go is statically typed.
* Go checks types at compile time.
* Go programs are generally faster and safer.

---

## Advantages of Variables

* Store reusable data
* Improve readability
* Reduce duplication
* Simplify maintenance
* Enable dynamic applications

---

## Key Takeaways

* Variables store data in memory.
* Go supports multiple declaration styles.
* `:=` is the most common declaration syntax.
* Go uses static typing.
* Variables are fundamental to every Go application.
