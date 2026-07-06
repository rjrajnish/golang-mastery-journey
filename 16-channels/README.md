 # Channels in Go

Channels are one of Go's most powerful features. They allow goroutines to communicate safely without sharing memory directly.

Go follows a simple philosophy:

> **Don't communicate by sharing memory; share memory by communicating.**

Instead of multiple goroutines modifying the same variable, they exchange data through channels.

Channels make concurrent programs easier to understand, safer to write, and less prone to race conditions.

---

# Table of Contents

1. What is a Channel?
2. Why Do We Need Channels?
3. Creating a Channel
4. Sending Data
5. Receiving Data
6. Buffered Channels
7. Closing Channels
8. Range over Channels
9. Directional Channels
10. Channels with WaitGroup
11. JavaScript vs Go
12. Memory Representation
13. Real-World Use Cases
14. Best Practices
15. Common Mistakes
16. Mini Project
17. Interview Questions
18. Key Takeaways

---

# What is a Channel?

A channel is a communication mechanism that allows goroutines to send and receive values safely.

Think of a channel as a pipe.

```
Goroutine A

      │

      ▼

  CHANNEL

      │

      ▼

Goroutine B
```

One goroutine sends data.

Another goroutine receives it.

---

# Why Do We Need Channels?

Imagine two goroutines updating the same variable.

```
goroutine1

balance++

goroutine2

balance--
```

Both modify the same memory.

This may produce incorrect results.

Using channels:

```
goroutine1

↓

Channel

↓

goroutine2
```

Data flows safely.

No shared variable is required.

---

# Creating a Channel

```go
ch := make(chan string)
```

This creates an unbuffered channel.

---

# Sending Data

Use the `<-` operator.

```go
ch <- "Hello"
```

The value is sent into the channel.

---

# Receiving Data

```go
message := <-ch
```

The receiver waits until data becomes available.

---

# Channel Flow

```
Sender

↓

Channel

↓

Receiver
```

Channels synchronize goroutines automatically.

---

# Buffered Channels

Normally, a sender waits until another goroutine receives the value.

Buffered channels provide temporary storage.

```go
buffer := make(chan string, 2)
```

Capacity = 2

You can send two values before blocking.

Example:

```go
buffer <- "Go"

buffer <- "Programming"
```

---

# Closing a Channel

When no more values will be sent:

```go
close(ch)
```

Closing tells receivers that transmission is complete.

Do **not** send values after closing a channel.

---

# Range over Channels

After closing a channel:

```go
for value := range ch {

	fmt.Println(value)

}
```

The loop automatically ends when the channel is closed.

---

# Directional Channels

Restrict channel usage.

Send only:

```go
func send(ch chan<- string)
```

Receive only:

```go
func receive(ch <-chan string)
```

Directional channels improve readability and prevent accidental misuse.

---

# Channels with WaitGroup

Channels and `sync.WaitGroup` are often used together.

```
Worker 1

↓

Worker 2

↓

Worker 3

↓

Channel

↓

Main
```

Workers send results.

The main goroutine waits for completion and processes the results.

---

# JavaScript vs Go

## JavaScript

```javascript
async function fetchData() {

	return "Hello";

}
```

Communication is handled with Promises and the event loop.

---

## Go

```go
go worker(ch)

message := <-ch
```

Communication happens through channels.

### Comparison

| JavaScript | Go |
|------------|----|
| Promise | Channel |
| async / await | Goroutine + Channel |
| Event Loop | Go Scheduler |
| Callback Queue | Channel Communication |

---

# Memory Representation

```
Main

│

├── Worker 1

│

├── Worker 2

│

└── Worker 3

        │

        ▼

     Channel

        │

        ▼

     Main
```

Instead of sharing variables, workers communicate through the channel.

---

# Real-World Use Cases

Channels are commonly used for:

- Worker pools
- Background jobs
- Task queues
- HTTP request processing
- Event processing
- Notifications
- Data pipelines
- File downloads

---

# Best Practices

- Close channels when no more values will be sent.
- Use buffered channels only when needed.
- Prefer communication over shared memory.
- Combine channels with goroutines for concurrency.
- Keep channel ownership clear (typically the sender closes the channel).

---

# Common Mistakes

## Forgetting to Close a Channel

Bad:

```go
for value := range ch {

}
```

If the channel is never closed, the loop never ends.

---

## Sending to a Closed Channel

```go
close(ch)

ch <- "Hello"
```

Runtime panic.

---

## Receiving Without a Sender

```go
message := <-ch
```

If no goroutine sends data, the program blocks indefinitely.

---

# Mini Project

Build a Concurrent Task Processor.

Requirements:

- Create five worker goroutines.
- Send tasks through a channel.
- Process tasks concurrently.
- Collect results.
- Print a completion message.

Bonus:

Measure total execution time.

---

# Interview Questions

### What is a channel?

A channel is a communication mechanism between goroutines.

---

### Why use channels?

To safely exchange data between goroutines without sharing memory directly.

---

### Difference between buffered and unbuffered channels?

Unbuffered channels block until both sender and receiver are ready.

Buffered channels allow a limited number of values to be stored before blocking.

---

### Who should close a channel?

Typically, the goroutine that sends values is responsible for closing the channel.

---

### Can you send data after closing a channel?

No.

It causes a runtime panic.

---

# Key Takeaways

- Channels enable safe communication between goroutines.
- Use `make()` to create channels.
- Send values with `ch <- value`.
- Receive values with `<-ch`.
- Close channels when finished sending.
- Use `range` to receive values until a channel is closed.
- Buffered channels reduce blocking in some scenarios.
- Channels are a core part of Go's concurrency model.

---

