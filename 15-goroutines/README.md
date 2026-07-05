# Goroutines in Go

Concurrency is one of Go's biggest strengths. Instead of waiting for one task to finish before starting another, Go allows multiple tasks to run concurrently using **goroutines**.

Goroutines are lightweight functions managed by the Go runtime. They make it easy to build fast, scalable, and responsive applications.

Whether you're processing API requests, downloading files, sending emails, or handling background jobs, goroutines help your application do more work in less time.

---

# Table of Contents

1. What is Concurrency?
2. What is a Goroutine?
3. Why Do We Need Goroutines?
4. Creating a Goroutine
5. Multiple Goroutines
6. Anonymous Goroutines
7. WaitGroup
8. JavaScript vs Go
9. Concurrency vs Parallelism
10. Memory Representation
11. Real-World Use Cases
12. Best Practices
13. Common Mistakes
14. Mini Project
15. Interview Questions
16. Key Takeaways

---

# What is Concurrency?

Concurrency means **handling multiple tasks during the same period of time**.

Imagine a restaurant.

One chef:

```
Cook Food

↓

Serve Food

↓

Wash Dishes

↓

Take New Order
```

Everything happens one after another.

Now imagine three chefs.

```
Chef 1 → Cooking

Chef 2 → Serving

Chef 3 → Washing
```

Multiple tasks progress at the same time.

That's concurrency.

---

# What is a Goroutine?

A goroutine is a lightweight function that runs concurrently with other functions.

Creating one is simple.

```go
go printNumbers()
```

The `go` keyword tells the Go runtime to execute the function concurrently.

---

# Why Do We Need Goroutines?

Without goroutines:

```
Task A

↓

Task B

↓

Task C
```

Each task waits for the previous one.

With goroutines:

```
Task A

Task B

Task C
```

All tasks begin immediately and run concurrently.

Benefits:

- Better CPU utilization
- Faster response time
- Higher throughput
- Improved scalability

---

# Creating a Goroutine

Normal function:

```go
printNumbers()
```

Concurrent execution:

```go
go printNumbers()
```

The caller continues executing without waiting for the goroutine to finish.

---

# Multiple Goroutines

Launching multiple goroutines is straightforward.

```go
go printNumbers()

go printLetters()
```

The Go scheduler decides when each goroutine runs.

---

# Anonymous Goroutines

Anonymous functions can also run as goroutines.

```go
go func() {

	fmt.Println("Hello")

}()
```

Useful for:

- Background jobs
- Short-lived tasks
- Event processing

---

# WaitGroup

If the `main` function exits, all goroutines stop immediately.

Use a `sync.WaitGroup` to wait for them.

```go
var wg sync.WaitGroup

wg.Add(1)

go worker(&wg)

wg.Wait()
```

A WaitGroup ensures all registered goroutines finish before the program exits.

---

# JavaScript vs Go

## JavaScript

```javascript
setTimeout(() => {

	console.log("Hello");

}, 1000);
```

JavaScript uses an event loop and asynchronous callbacks.

---

## Go

```go
go printNumbers()
```

Go uses lightweight goroutines scheduled by the Go runtime.

### Comparison

| JavaScript | Go |
|------------|----|
| Single-threaded event loop | Goroutines managed by Go runtime |
| Promise / async-await | Goroutines |
| Web APIs / libuv | Go Scheduler |
| Callback queue | Goroutine Scheduler |

---

# Concurrency vs Parallelism

These terms are related but different.

## Concurrency

Handling multiple tasks during the same time period.

```
Task A

Task B

Task C
```

Progress overlaps.

---

## Parallelism

Executing multiple tasks at the exact same time on multiple CPU cores.

```
CPU 1 → Task A

CPU 2 → Task B

CPU 3 → Task C
```

Parallelism requires multiple cores.

---

# Memory Representation

```
Main()

│

├── Goroutine 1

├── Goroutine 2

├── Goroutine 3

└── Goroutine 4
```

The Go scheduler efficiently manages thousands of goroutines.

Unlike operating system threads, goroutines have very small initial stack sizes, allowing Go programs to run many concurrent tasks efficiently.

---

# Real-World Use Cases

Goroutines are commonly used for:

- HTTP request handling
- Background jobs
- Email sending
- Image processing
- File downloads
- Log processing
- Data synchronization
- Web scraping
- Worker pools

Almost every production Go backend uses goroutines.

---

# Best Practices

- Keep goroutines focused on a single task.
- Use `sync.WaitGroup` when waiting for multiple goroutines.
- Protect shared data using synchronization primitives (covered in later topics).
- Avoid creating unnecessary goroutines.
- Always consider cancellation for long-running tasks (using `context`, covered later).

---

# Common Mistakes

## Main Function Exits Too Early

Bad:

```go
go printNumbers()
```

The program may exit before the goroutine finishes.

Use a WaitGroup.

---

## Creating Too Many Goroutines

Launching millions of goroutines without limits can consume memory and overwhelm downstream systems.

Use worker pools when appropriate.

---

## Assuming Execution Order

Never assume goroutines execute in the order they were started.

The scheduler determines execution timing.

---

# Mini Project

Build a Concurrent Download Simulator.

Requirements:

- Download multiple files concurrently.
- Show download progress.
- Wait for all downloads to complete.
- Display a final success message.

Bonus:

Measure total execution time using the `time` package.

---

# Interview Questions

### What is a goroutine?

A lightweight function managed by the Go runtime that executes concurrently.

---

### How do you start a goroutine?

Using the `go` keyword.

```go
go worker()
```

---

### What is `sync.WaitGroup`?

A synchronization primitive used to wait for multiple goroutines to finish.

---

### Are goroutines operating system threads?

No.

They are lightweight execution units managed by the Go runtime, which schedules them onto operating system threads.

---

### What is the difference between concurrency and parallelism?

Concurrency is about managing multiple tasks during the same period.

Parallelism is about executing multiple tasks simultaneously on multiple CPU cores.

---

# Key Takeaways

- Goroutines enable lightweight concurrency.
- Start a goroutine using the `go` keyword.
- Use `sync.WaitGroup` to wait for goroutines.
- Goroutines are managed by the Go runtime, not directly by the operating system.
- Goroutines are widely used in production Go applications.
- Concurrency is one of Go's biggest strengths.

---
