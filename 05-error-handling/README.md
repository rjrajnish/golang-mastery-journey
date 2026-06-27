# Error Handling in Go

---

# Introduction

Errors are a normal part of software development.

Files may not exist.

Network requests may fail.

Database connections may be unavailable.

User input may be invalid.

Instead of hiding these failures, Go makes error handling an essential part of writing reliable software.

Unlike many programming languages, Go does **not** rely on exceptions for normal application flow. Instead, functions return an `error` value that developers are expected to check and handle.

This design makes Go applications easier to understand, debug, and maintain.

---

# What is an Error?

An error is simply a value that indicates something went wrong during program execution.

In Go, the built-in `error` interface represents an error.

```go
type error interface {
    Error() string
}
```

Any type that implements the `Error()` method satisfies this interface.

---

# Why Does Go Use Errors Instead of Exceptions?

Go's creators intentionally avoided exception-driven programming.

Their philosophy is:

> "Errors are expected. Panics are exceptional."

Returning errors makes the flow of a program explicit.

Instead of hidden exceptions jumping between functions, every possible failure is visible in the function signature.

Example:

```go
func ReadFile(path string) ([]byte, error)
```

Immediately you know this function may fail.

---

# JavaScript vs Go

### JavaScript

```javascript
try {
    const data = fs.readFileSync("file.txt");
} catch(err) {
    console.log(err);
}
```

### Go

```go
data, err := os.ReadFile("file.txt")

if err != nil {
    fmt.Println(err)
}
```

### Comparison

| JavaScript            | Go                    |
| --------------------- | --------------------- |
| Uses exceptions       | Uses error values     |
| try/catch             | if err != nil         |
| Hidden control flow   | Explicit control flow |
| Exceptions can bubble | Errors are returned   |

Go encourages developers to think about failure as part of normal program execution.

---

# Creating Errors

Using `errors.New()`:

```go
return errors.New("invalid username")
```

Use this when you need a simple error message.

---

# Formatted Errors

Sometimes you need additional context.

```go
return fmt.Errorf("login failed: %w", err)
```

This wraps the original error while adding more information.

---

# Error Wrapping

Error wrapping helps preserve the original cause of a failure.

Example:

```
Database Error
↓

Repository Error
↓

Service Error
↓

HTTP Handler
```

Each layer adds context while keeping the root error intact.

---

# Custom Errors

Large applications often require custom error types.

Example:

```go
type ValidationError struct {
    Message string
}
```

Custom errors allow applications to distinguish between different failure scenarios.

Examples:

* ValidationError
* AuthenticationError
* DatabaseError
* PaymentError

---

# Panic

A panic immediately stops normal execution.

```go
panic("database connection lost")
```

Panic should **not** be used for normal business logic.

Use panic only when the application reaches an unrecoverable state.

---

# Recover

`recover()` prevents the application from crashing after a panic.

```go
defer func() {

    if r := recover(); r != nil {
        fmt.Println(r)
    }

}()
```

Recover is commonly used in:

* HTTP middleware
* Worker pools
* Background jobs

---

# Best Practices

Always check returned errors.

```go
if err != nil {
    return err
}
```

Wrap errors with meaningful context.

```go
fmt.Errorf("failed to load config: %w", err)
```

Do not ignore errors.

Bad:

```go
result, _ := divide(10, 2)
```

Good:

```go
result, err := divide(10, 2)

if err != nil {
    return
}
```

Use panic only for unrecoverable situations.

---

# Real-World Example

Imagine an API request.

```
HTTP Request
↓

Validate Input

↓

Read Database

↓

Return Response
```

Every step can fail.

Go returns an error at each stage, allowing the caller to decide how to respond instead of crashing the application.

 

# Key Takeaways

* Errors are values in Go.
* Handle every returned error.
* Use `errors.New()` for simple errors.
* Use `fmt.Errorf()` for additional context.
* Create custom error types for complex applications.
* Reserve `panic` for unrecoverable situations.
* Use `recover` to prevent unexpected crashes.

---
 
