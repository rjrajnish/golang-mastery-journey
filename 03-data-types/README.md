# Data Types in Go

## What are Data Types?

Data types define the kind of value a variable can store.

Every variable in Go has a specific type.

Examples:

```go
var name string = "Rajnish"
var age int = 25
var active bool = true
```

Go uses static typing, meaning the type of a variable is checked during compilation.

---

## Why Do We Need Data Types?

Data types help Go:

* Store data efficiently
* Use memory correctly
* Prevent programming mistakes
* Improve performance
* Increase code readability

Without data types, Go would not know how to process data.

---

## When Should We Use Data Types?

Every variable requires a data type.

Examples:

* User Name → string
* User Age → int
* Product Price → float64
* Login Status → bool

---

## Where Are Data Types Used?

Data types are used everywhere:

### Authentication System

```go
var username string
var isLoggedIn bool
```

### E-Commerce Application

```go
var productName string
var price float64
```

### Banking Application

```go
var accountBalance float64
```

### Social Media Platform

```go
var likes int
```

---

# Common Data Types in Go

## String

Stores text.

```go
var name string = "Rajnish"
```

Examples:

* Names
* Addresses
* Emails
* Messages

---

## Integer (int)

Stores whole numbers.

```go
var age int = 25
```

Examples:

* Age
* Quantity
* Likes
* User IDs

---

## Float64

Stores decimal values.

```go
var salary float64 = 50000.50
```

Examples:

* Salary
* Product Price
* Interest Rate

---

## Boolean (bool)

Stores true or false.

```go
var isDeveloper bool = true
```

Examples:

* Login Status
* Feature Flags
* Permissions

---

## Byte

Represents ASCII characters.

```go
var grade byte = 'A'
```

Commonly used in:

* Networking
* File Processing
* Binary Data

---

## Rune

Represents Unicode characters.

```go
var heart rune = '❤'
```

Useful for:

* International Languages
* Emoji Support
* Unicode Text Processing

---

# Type Inference

Go can automatically detect variable types.

```go
name := "Rajnish"
age := 25
salary := 50000.50
```

Go automatically assigns:

```text
string
int
float64
```

This is called Type Inference.

---

# Zero Values

Every Go variable has a default value.

```go
var name string
var age int
var active bool
```

Default values:

```text
name = ""
age = 0
active = false
```

Go never leaves variables uninitialized.

---

# Type Conversion

Go does not automatically convert data types.

Incorrect:

```go
var age int = 25
var salary float64 = age
```

This causes a compile error.

Correct:

```go
var age int = 25
var salary float64 = float64(age)
```

Always convert explicitly.

---

# JavaScript vs Go

## JavaScript

```javascript
let age = 25;
let name = "Rajnish";
let active = true;
```

JavaScript determines types at runtime.

---

## Go

```go
age := 25
name := "Rajnish"
active := true
```

Go determines types during compilation.

---

# Major Differences

| Feature           | JavaScript | Go           |
| ----------------- | ---------- | ------------ |
| Typing            | Dynamic    | Static       |
| Performance       | Runtime    | Compile Time |
| Type Safety       | Lower      | Higher       |
| Memory Efficiency | Lower      | Higher       |

Because of static typing, Go applications are generally faster and safer.

---

# Advantages of Strong Data Types

* Fewer runtime bugs
* Better IDE support
* Faster applications
* Better memory management
* Easier maintenance

---
# Key Takeaways

* Every variable has a type.
* Go uses static typing.
* Common types include string, int, float64, and bool.
* Go supports byte and rune for character handling.
* Go does not perform automatic type conversions.
* Strong typing improves performance and reliability.
