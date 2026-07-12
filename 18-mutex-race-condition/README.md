 # Mutex & Race Conditions in Go

Concurrency makes Go applications fast and scalable, but it also introduces one of the most common problems in software development: **race conditions**.

Whenever multiple goroutines access and modify the same data simultaneously, the final result can become unpredictable.

Go provides synchronization primitives such as `sync.Mutex` and `sync.RWMutex` to ensure that shared data is accessed safely.

Understanding mutexes is essential for writing reliable concurrent programs.

---

# Table of Contents

1. What is a Race Condition?
2. Why Race Conditions Happen
3. What is a Mutex?
4. Lock & Unlock
5. Read-Write Mutex
6. Deadlocks
7. JavaScript vs Go
8. Memory Representation
9. Real-World Use Cases
10. Best Practices
11. Common Mistakes
12. Mini Project
13. Interview Questions
14. Key Takeaways

---

# What is a Race Condition?

A race condition occurs when two or more goroutines read or modify the same variable at the same time, leading to unpredictable results.

Imagine two employees updating the same bank account balance simultaneously.

```
Balance = 1000

Employee A

Withdraw 200

Employee B

Deposit 500
```

If both operations happen at the same time without synchronization, the final balance may be incorrect.

---

# Why Race Conditions Happen

Suppose:

```go
counter++
```

This looks like one operation.

Actually it consists of:

```
Read counter

↓

Add 1

↓

Write counter
```

If another goroutine modifies the value during these steps, the result becomes inconsistent.

---

# What is a Mutex?

A mutex (mutual exclusion) allows **only one goroutine at a time** to access a shared resource.

Think of a public restroom with one key.

```
Person A

↓

LOCK

↓

Uses Room

↓

UNLOCK

↓

Person B
```

Only one person can enter at a time.

---

# Lock & Unlock

```go
mutex.Lock()

counter++

mutex.Unlock()
```

The lock protects the critical section.

While one goroutine holds the lock, others must wait.

---

# Read-Write Mutex (`sync.RWMutex`)

Sometimes many goroutines only need to read data.

Using a normal mutex blocks everyone.

`RWMutex` improves performance.

Read:

```go
rwMutex.RLock()

fmt.Println(counter)

rwMutex.RUnlock()
```

Write:

```go
rwMutex.Lock()

counter++

rwMutex.Unlock()
```

Rules:

- Multiple readers are allowed simultaneously.
- Only one writer is allowed.
- Writers block readers and other writers until finished.

---

# Deadlocks

A deadlock occurs when goroutines wait forever for locks that are never released.

Example:

```go
mutex.Lock()

// Forgot Unlock()
```

Every other goroutine attempting to acquire the lock will wait indefinitely.

A common practice is:

```go
mutex.Lock()

defer mutex.Unlock()
```

This ensures the lock is always released, even if the function returns early.

---

# JavaScript vs Go

## JavaScript

```javascript
let counter = 0;

counter++;
```

JavaScript typically runs user code on a single thread, so race conditions on ordinary variables are less common (though asynchronous workflows introduce different coordination challenges).

---

## Go

```go
go increment()

go increment()
```

Multiple goroutines can execute concurrently.

Without synchronization, race conditions become possible.

---

# Comparison

| JavaScript | Go |
|------------|----|
| Single-threaded event loop | Multiple goroutines |
| Shared state less common | Shared memory possible |
| Promises | Goroutines |
| Mutex rarely needed | Mutex frequently needed |

---

# Memory Representation

Without Mutex

```
Worker 1

↓

Counter

↑

Worker 2
```

Both access the same memory simultaneously.

---

With Mutex

```
Worker 1

↓

LOCK

↓

Counter

↓

UNLOCK

↓

Worker 2
```

Access is serialized, preventing race conditions.

---

# Real-World Use Cases

Mutexes are commonly used for:

- Bank account balances
- Shopping carts
- Inventory systems
- Session storage
- Caches
- API rate limiting
- Shared configuration
- Counters and metrics

Whenever multiple goroutines share mutable data, synchronization is required.

---

# Best Practices

- Protect shared mutable state with a mutex.
- Keep critical sections as short as possible.
- Prefer `defer mutex.Unlock()` after locking.
- Use `RWMutex` when reads greatly outnumber writes.
- Consider channels when communication is more appropriate than shared memory.

---

# Common Mistakes

## Forgetting to Unlock

```go
mutex.Lock()

// Missing Unlock()
```

This can cause a deadlock.

---

## Locking Too Much Code

Bad:

```go
mutex.Lock()

// Long-running work

mutex.Unlock()
```

Keep only the shared data access inside the lock.

---

## Using RWMutex for Heavy Writes

If writes happen frequently, `RWMutex` may not provide much benefit over a regular `Mutex`.

---

# Mini Project

Build an Inventory Management System.

Requirements:

- Product stock
- Concurrent purchases
- Concurrent restocking
- Prevent negative stock
- Protect shared data using `sync.Mutex`

Bonus:

Add multiple readers using `sync.RWMutex`.

---

# Interview Questions

### What is a race condition?

A race condition occurs when multiple goroutines access shared data concurrently and at least one modifies it, leading to unpredictable results.

---

### What is a mutex?

A synchronization primitive that allows only one goroutine to access a critical section at a time.

---

### What is the difference between `Mutex` and `RWMutex`?

- `Mutex` allows one reader or writer.
- `RWMutex` allows multiple readers but only one writer.

---

### What is a deadlock?

A situation where goroutines wait indefinitely for resources or locks, preventing progress.

---

### How can you detect race conditions?

Run your program with Go's race detector:

```bash
go run -race main.go
```

or

```bash
go test -race
```

The race detector reports potential data races during execution.

---

# Key Takeaways

- Race conditions occur when shared data is accessed concurrently without proper synchronization.
- `sync.Mutex` protects critical sections.
- `sync.RWMutex` allows multiple readers and one writer.
- Always release locks, preferably with `defer`.
- Keep critical sections small.
- Use the race detector to identify synchronization issues during development.

---
 
 