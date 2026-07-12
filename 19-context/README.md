 # Context in Go

`context.Context` is one of the most important packages in Go. It allows different parts of your application to communicate important request-scoped information such as cancellation signals, deadlines, timeouts, and metadata.

Whenever a request starts, the context travels with it through every layer of the application.

Without context, background tasks may continue running even after a request has been cancelled or timed out, wasting CPU, memory, database connections, and network resources.

---

# Table of Contents

1. What is Context?
2. Why Do We Need Context?
3. Background Context
4. Cancellation
5. Timeout
6. Deadline
7. Context Values
8. JavaScript vs Go
9. Memory Representation
10. Real-World Use Cases
11. Best Practices
12. Common Mistakes
13. Mini Project
14. Interview Questions
15. Key Takeaways

---

# What is Context?

A context carries request-scoped information across API boundaries and between goroutines.

It can carry:

- Cancellation signals
- Timeouts
- Deadlines
- Request values

Think of it as a control object that accompanies a request throughout its lifecycle.

---

# Why Do We Need Context?

Imagine a client sends a request to download a large file.

```
User

↓

HTTP Server

↓

Database

↓

File Service

↓

Cloud Storage
```

Now imagine the user closes the browser after one second.

Without context:

```
Download Continues

Database Query Continues

Memory Continues Being Used
```

Resources are wasted.

With context:

```
User Cancels

↓

Context Cancelled

↓

All Operations Stop
```

Everything shuts down cleanly.

---

# Background Context

Every context chain starts with a root context.

```go
ctx := context.Background()
```

This is commonly used in:

- Main functions
- Application startup
- Tests

---

# WithCancel

Create a cancellable context.

```go
ctx, cancel := context.WithCancel(
	context.Background(),
)

defer cancel()
```

Calling:

```go
cancel()
```

immediately notifies all goroutines using that context.

---

# WithTimeout

Automatically cancel after a specified duration.

```go
ctx, cancel := context.WithTimeout(
	context.Background(),
	5*time.Second,
)

defer cancel()
```

Useful for:

- HTTP requests
- Database queries
- External APIs

---

# WithDeadline

Cancel at a specific time.

```go
deadline := time.Now().Add(10 * time.Second)

ctx, cancel := context.WithDeadline(
	context.Background(),
	deadline,
)
```

Unlike `WithTimeout`, you specify the exact deadline.

---

# Context Values

Store request-scoped metadata.

```go
type contextKey string

const UserIDKey contextKey = "userID"

ctx := context.WithValue(
	context.Background(),
	UserIDKey,
	101,
)
```

Retrieve it:

```go
id := ctx.Value(UserIDKey)
```

Use context values only for request-scoped data, not for passing optional parameters or configuration.

---

# JavaScript vs Go

## JavaScript

```javascript
const controller = new AbortController();

fetch(url, {
	signal: controller.signal,
});
```

AbortController cancels requests.

---

## Go

```go
ctx, cancel := context.WithCancel(
	context.Background(),
)

defer cancel()
```

Context provides cancellation, deadlines, timeouts, and request values.

### Comparison

| JavaScript | Go |
|------------|----|
| AbortController | Context |
| signal | ctx.Done() |
| Abort | cancel() |
| Timeout libraries | context.WithTimeout() |

---

# Memory Representation

```
HTTP Request

      │

      ▼

Context

      │

      ├── Database

      ├── Redis

      ├── HTTP Client

      └── Goroutines
```

When the context is cancelled, every dependent operation receives the cancellation signal.

---

# Real-World Use Cases

Context is used in:

- REST APIs
- Database queries
- gRPC
- HTTP clients
- Background workers
- Cloud services
- Message queues
- Microservices

It is a standard part of modern Go APIs.

---

# Best Practices

- Pass `context.Context` as the first parameter to functions that perform request-scoped work.
- Always call the returned `cancel()` function when appropriate.
- Respect `ctx.Done()` in long-running operations.
- Use `context.WithTimeout()` for external operations.
- Use custom key types for context values to avoid collisions.

---

# Common Mistakes

## Ignoring Cancellation

Bad:

```go
for {

	doWork()

}
```

Better:

```go
select {

case <-ctx.Done():

	return

default:

	doWork()

}
```

---

## Storing Large Objects

Avoid storing database connections or configuration inside context values.

Context values are intended only for lightweight request-scoped data.

---

## Forgetting cancel()

When you create a derived context with `WithCancel`, `WithTimeout`, or `WithDeadline`, remember to call the returned `cancel()` to release resources.

---

# Mini Project

Build a File Download Manager.

Requirements:

- Start multiple downloads.
- Add a timeout.
- Allow user cancellation.
- Stop all goroutines when cancelled.
- Display cancellation reason.

Bonus:

Track download progress.

---

# Interview Questions

### What is `context.Context`?

A type used to carry cancellation signals, deadlines, timeouts, and request-scoped values across API boundaries.

---

### Why is Context important?

It prevents wasted work by allowing related operations to stop when a request is cancelled or times out.

---

### Difference between `WithCancel` and `WithTimeout`?

- `WithCancel` cancels manually.
- `WithTimeout` cancels automatically after a duration.

---

### When should you use `WithValue`?

Only for lightweight request-scoped metadata such as request IDs or authenticated user IDs.

---

### Where should Context be passed?

As the first parameter to functions performing request-scoped work.

Example:

```go
func GetUser(ctx context.Context, id int)
```

---

# Key Takeaways

- Context manages cancellation, timeouts, deadlines, and request values.
- Use `context.Background()` as the root context.
- Use `WithCancel()` for manual cancellation.
- Use `WithTimeout()` to limit operation duration.
- Use `WithDeadline()` for absolute deadlines.
- Check `ctx.Done()` in long-running operations.
- Pass context as the first function parameter.

---
 