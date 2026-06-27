# Arrays in Go

Arrays are one of the fundamental data structures in Go. They allow you to store multiple values of the same data type in a single variable. Although arrays are not used as frequently as slices in modern Go applications, understanding them is essential because slices are built on top of arrays.

If you're coming from JavaScript, you may notice that Go arrays behave very differently. In JavaScript, arrays are dynamic and act like objects, while in Go, arrays have a fixed size and are treated as value types.

This guide explains everything you need to know about arrays in Go with practical examples and comparisons.

---

# Table of Contents

- What is an Array?
- Why Do We Need Arrays?
- Array Declaration
- Array Initialization
- Accessing Array Elements
- Updating Array Elements
- Looping Through Arrays
- Array Length
- Multidimensional Arrays
- Arrays are Value Types
- Passing Arrays to Functions
- Arrays vs Slices
- JavaScript vs Go
- Real-World Use Cases
- Best Practices
- Common Mistakes
- Mini Project
- Interview Questions
- Key Takeaways

---

# What is an Array?

An array is a collection of elements of the **same data type** stored in **contiguous memory locations**.

Each element has a fixed position called an **index**, starting from **0**.

Example:

```go
var numbers [5]int
```

This creates an array capable of storing **5 integers**.

---

# Why Do We Need Arrays?

Imagine you're building a student management system.

Without arrays:

```go
student1 := 85
student2 := 90
student3 := 78
student4 := 88
student5 := 95
```

Managing many variables quickly becomes difficult.

With an array:

```go
marks := [5]int{85, 90, 78, 88, 95}
```

Now all student marks are stored together, making your code cleaner and easier to maintain.

---

# Characteristics of Arrays

- Fixed size
- Same data type
- Stored in contiguous memory
- Fast element access
- Index starts from 0
- Value type (copied when assigned)

---

# Array Declaration

There are multiple ways to declare arrays.

## Empty Array

```go
var numbers [5]int
```

Output:

```
[0 0 0 0 0]
```

Go automatically initializes each element with its zero value.

---

## Array with Values

```go
numbers := [5]int{10, 20, 30, 40, 50}
```

Output

```
[10 20 30 40 50]
```

---

## Let Go Count the Elements

```go
numbers := [...]int{10, 20, 30, 40}
```

Go automatically determines the array length.

---

# Accessing Elements

Every element is accessed using its index.

```go
marks := [5]int{80, 85, 90, 75, 95}

fmt.Println(marks[0])

fmt.Println(marks[4])
```

Output

```
80
95
```

Remember:

```
Index

0 1 2 3 4
```

---

# Updating Elements

Array elements can be modified using their index.

```go
marks[1] = 100
```

Before

```
[80 85 90 75 95]
```

After

```
[80 100 90 75 95]
```

---

# Looping Through Arrays

## Using a for Loop

```go
for i := 0; i < len(marks); i++ {
    fmt.Println(marks[i])
}
```

---

## Using range

```go
for index, value := range marks {
    fmt.Println(index, value)
}
```

Output

```
0 80
1 100
2 90
3 75
4 95
```

`range` is the preferred way to iterate over arrays in Go.

---

# Finding Array Length

Use the built-in `len()` function.

```go
fmt.Println(len(marks))
```

Output

```
5
```

---

# Multidimensional Arrays

Arrays can contain other arrays.

Example:

```go
matrix := [2][3]int{
    {1,2,3},
    {4,5,6},
}
```

Output

```
[[1 2 3]
 [4 5 6]]
```

Useful for:

- Matrix calculations
- Board games
- Image processing

---

# Arrays are Value Types

One of the biggest differences between Go and JavaScript.

Example:

```go
a := [3]int{1,2,3}

b := a

b[0] = 100
```

Output

```
a = [1 2 3]

b = [100 2 3]
```

Why?

Because Go copies the entire array.

The original array remains unchanged.

---

# Passing Arrays to Functions

Passing an array creates a copy.

```go
func update(arr [3]int){
    arr[0] = 999
}
```

Original array is not modified.

If you want to modify the original array, pass a pointer.

```go
func update(arr *[3]int){
    arr[0] = 999
}
```

---

# Arrays vs Slices

| Arrays | Slices |
|---------|---------|
| Fixed Size | Dynamic Size |
| Value Type | Reference-like Descriptor |
| Rarely Used | Used Everywhere |
| Size Cannot Change | Size Can Grow |

Arrays are mainly used to understand slices and in situations where the size is known beforehand.

---

# JavaScript vs Go

## JavaScript

```javascript
let numbers = [10,20,30]

let copy = numbers

copy[0] = 100

console.log(numbers)
```

Output

```
[100,20,30]
```

Arrays are reference types.

---

## Go

```go
numbers := [3]int{10,20,30}

copy := numbers

copy[0] = 100

fmt.Println(numbers)
```

Output

```
[10 20 30]
```

Go copies the array instead of sharing it.

This behavior prevents accidental modifications.

---

# Memory Representation

```
Index

0      1      2      3

+------+------+------+------+

10     20     30     40

+------+------+------+------+
```

Arrays are stored in continuous memory, making element access extremely fast.

---

# Real-World Use Cases

Although slices are more common, arrays are still useful in:

- Image Processing
- Cryptography
- Fixed-size Buffers
- Embedded Systems
- Scientific Computing
- Matrix Operations

---

# Best Practices

- Use arrays only when the size is fixed.
- Prefer slices for most applications.
- Use `range` when iterating.
- Pass pointers if modifying large arrays.
- Keep array sizes small and predictable.

---

# Common Mistakes

## Accessing an Invalid Index

```go
numbers[5]
```

Error

```
panic: index out of range
```

---

## Wrong Array Size

```go
var numbers [3]int

numbers = [4]int{1,2,3,4}
```

Compile Error

Different array sizes are different types.

---

## Assuming Arrays Behave Like JavaScript

Many beginners expect copied arrays to share data.

They do not.

Every assignment creates a new copy.

---

# Mini Project

Create a Student Marks Management System.

Requirements:

- Store marks of 5 students
- Print all marks
- Calculate total marks
- Calculate average
- Find highest mark
- Find lowest mark

Bonus:

Display grade based on average marks.

---

# Interview Questions

### What is an array?

An array is a fixed-size collection of elements of the same type.

---

### Are arrays value types?

Yes.

Every assignment creates a copy.

---

### Difference between arrays and slices?

Arrays have fixed size.

Slices have dynamic size.

---

### Why are arrays rarely used?

Because slices provide more flexibility.

---

### Can array sizes change?

No.

Their size is fixed during declaration.

---

### Are two arrays with different sizes the same type?

No.

```go
[3]int
```

and

```go
[4]int
```

are different types.

---

# Key Takeaways

- Arrays store elements of the same data type.
- Arrays have a fixed size.
- Arrays are stored in contiguous memory.
- Arrays are value types.
- Assigning an array creates a copy.
- Use `range` to iterate over arrays.
- Use pointers to modify arrays inside functions.
- Prefer slices for most real-world applications.

---
 