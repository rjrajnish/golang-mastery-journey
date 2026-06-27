# Slices in Go

Slices are one of the most powerful and frequently used data structures in Go. If arrays are fixed-size containers, slices are flexible views over arrays that can grow and shrink as needed.

In fact, most Go programs use slices instead of arrays because they provide dynamic sizing, easier manipulation, and better flexibility.

If you're a JavaScript developer, you can think of slices as being closer to JavaScript arrays—but with stronger typing and more predictable memory behavior.

---

# Table of Contents

- What is a Slice?
- Why Do We Need Slices?
- Slice vs Array
- Creating a Slice
- Length and Capacity
- The `append()` Function
- Slice Expressions
- Copying Slices
- The `make()` Function
- Nil Slices
- Removing Elements
- JavaScript vs Go
- Real-World Use Cases
- Best Practices
- Common Mistakes
- Mini Project
- Interview Questions
- Key Takeaways

---

# What is a Slice?

A slice is a dynamic, flexible view into an underlying array.

Unlike arrays, slices do **not** have a fixed size.

Example:

```go
numbers := []int{10, 20, 30}
```

This creates a slice containing three integers.

---

# Why Do We Need Slices?

Imagine you're building a Todo application.

You don't know how many tasks the user will create.

Using an array:

```go
var tasks [100]string
```

You must decide the size in advance.

Using a slice:

```go
tasks := []string{}
```

You can add tasks whenever needed.

That's why slices are preferred in most Go applications.

---

# Slice vs Array

| Feature | Array | Slice |
|---------|-------|--------|
| Size | Fixed | Dynamic |
| Can Grow | ❌ | ✅ |
| Common in Projects | Rare | Very Common |
| Declaration | `[5]int` | `[]int` |
| Underlying Storage | Own Memory | References an Array |

---

# Creating a Slice

```go
numbers := []int{10, 20, 30}
```

Output:

```
[10 20 30]
```

---

# Length and Capacity

Every slice has two important properties.

## Length

The number of elements currently in the slice.

```go
len(numbers)
```

## Capacity

The maximum number of elements the slice can hold before Go allocates a new underlying array.

```go
cap(numbers)
```

Understanding capacity is important for writing efficient Go programs.

---

# Adding Elements with `append()`

The `append()` function is used to add elements to a slice.

```go
numbers = append(numbers, 40)
```

Output:

```
[10 20 30 40]
```

If the slice reaches its capacity, Go automatically allocates a larger array behind the scenes.

---

# Slice Expressions

You can create a new slice from an existing slice.

```go
numbers := []int{10, 20, 30, 40, 50}

sub := numbers[1:4]
```

Output:

```
[20 30 40]
```

---

# Copying Slices

Assignment shares the same underlying data.

To create an independent copy:

```go
copySlice := make([]int, len(numbers))

copy(copySlice, numbers)
```

Now changes to one slice won't affect the other.

---

# The `make()` Function

The `make()` function creates slices with a specified length.

```go
users := make([]string, 3)
```

Output:

```
[  ]
```

You can then assign values:

```go
users[0] = "Rajnish"
```

---

# Nil Slices

A slice can be `nil`.

```go
var numbers []int
```

Check it:

```go
numbers == nil
```

Nil slices are common when a slice hasn't been initialized yet.

---

# Removing Elements

Go doesn't have a built-in remove function.

A common pattern is:

```go
index := 2

numbers = append(numbers[:index], numbers[index+1:]...)
```

This removes the element at index 2.

---

# JavaScript vs Go

## JavaScript

```javascript
let numbers = [10, 20, 30];

numbers.push(40);
```

JavaScript arrays are dynamic by default.

---

## Go

```go
numbers := []int{10, 20, 30}

numbers = append(numbers, 40)
```

Both grow dynamically, but Go provides stronger type safety and explicit control over memory.

---

# Real-World Use Cases

Slices are used throughout Go applications:

- API responses
- Database query results
- JSON arrays
- Lists of users
- Product catalogs
- File processing
- Log entries

Whenever you need a collection of items with a variable size, slices are usually the right choice.

---

# Best Practices

- Prefer slices over arrays for most applications.
- Use `append()` to add elements.
- Use `copy()` when you need an independent slice.
- Use `range` to iterate.
- Understand `len()` and `cap()` for performance optimization.

---

# Common Mistakes

## Forgetting to assign the result of `append()`

Incorrect:

```go
append(numbers, 40)
```

Correct:

```go
numbers = append(numbers, 40)
```

---

## Assuming slice assignment creates a copy

```go
a := []int{1,2,3}

b := a
```

Both slices reference the same underlying array.

Changing one may affect the other.

Use `copy()` when you need separate data.

---

# Mini Project

Build a Student Management System.

Requirements:

- Add student names
- Display all students
- Remove a student
- Find total students
- Search for a student

Use slices to store the student list.

---

# Interview Questions

### What is a slice?

A slice is a dynamically sized view into an underlying array.

---

### What is the difference between length and capacity?

Length is the number of elements.

Capacity is the total space available before reallocation.

---

### How do you add elements?

Using `append()`.

---

### How do you copy a slice?

Using the `copy()` function.

---

### Why are slices preferred over arrays?

Because they provide dynamic sizing, flexibility, and are used throughout the Go standard library.

---

# Key Takeaways

- Slices are dynamic collections built on arrays.
- Use `append()` to grow a slice.
- Understand `len()` and `cap()`.
- Use `copy()` for independent copies.
- Prefer slices over arrays in most real-world applications.
- Slices are one of the most important data structures in Go.

---
 