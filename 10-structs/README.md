 # Structs in Go

Structs are one of the most important features in Go. They allow you to combine multiple related fields into a single custom data type.

In real-world Go applications, almost everything is represented using structs:

- Users
- Products
- Orders
- API Requests
- API Responses
- Database Models
- Configuration
- Services

If arrays and maps store collections of data, **structs model real-world objects**.

---

# Table of Contents

1. What is a Struct?
2. Why Do We Need Structs?
3. Declaring a Struct
4. Creating Struct Instances
5. Accessing Fields
6. Updating Fields
7. Anonymous Structs
8. Nested Structs
9. Slice of Structs
10. Struct Comparison
11. Pointer to Struct
12. Zero Values
13. JavaScript vs Go
14. Memory Representation
15. Real-World Use Cases
16. Best Practices
17. Common Mistakes
18. Mini Project
19. Interview Questions
20. Key Takeaways

---

# What is a Struct?

A struct is a custom data type that groups multiple related fields together.

Instead of storing information in separate variables, a struct keeps everything organized.

Example:

```go
type Student struct {
	Name string
	Age  int
}
```

A `Student` now represents a real-world object with multiple properties.

---

# Why Do We Need Structs?

Imagine storing student information.

Without structs:

```go
name := "Rajnish"
age := 25
course := "Go"
```

Managing multiple students becomes difficult.

With structs:

```go
student := Student{
	Name: "Rajnish",
	Age: 25,
	Course: "Go",
}
```

Everything related to the student stays together.

---

# Declaring a Struct

```go
type Student struct {
	ID     int
	Name   string
	Age    int
	Course string
}
```

Each field has:

- Name
- Data Type

---

# Creating a Struct

```go
student := Student{
	ID: 101,
	Name: "Rajnish",
	Age: 25,
	Course: "Go",
}
```

Output:

```
{101 Rajnish 25 Go}
```

---

# Accessing Fields

Use the dot (`.`) operator.

```go
fmt.Println(student.Name)
fmt.Println(student.Course)
```

Output:

```
Rajnish
Go
```

---

# Updating Fields

Fields can be modified directly.

```go
student.Course = "Advanced Go"
```

---

# Anonymous Structs

Sometimes you need a temporary struct without defining a new type.

```go
user := struct {
	Name string
	Role string
}{
	Name: "Rahul",
	Role: "Admin",
}
```

Anonymous structs are useful for quick, one-time data structures.

---

# Nested Structs

Structs can contain other structs.

```go
type Address struct {
	City string
	State string
}

type Employee struct {
	Name string
	Address Address
}
```

This helps model complex data.

---

# Slice of Structs

You can create collections of structs.

```go
students := []Student{
	{1, "Raj", 24, "Go"},
	{2, "Rahul", 23, "React"},
}
```

Common use cases:

- User lists
- Product catalogs
- API responses

---

# Struct Comparison

Structs are comparable if all their fields are comparable.

```go
s1 := Student{1, "Raj", 24, "Go"}
s2 := Student{1, "Raj", 24, "Go"}

fmt.Println(s1 == s2)
```

Output:

```
true
```

---

# Pointer to Struct

Instead of copying the entire struct, use a pointer.

```go
ptr := &student

ptr.Name = "Rajnish Pandey"
```

Pointers improve performance when working with large structs.

---

# Zero Value

If a struct is declared without initialization, all fields receive zero values.

```go
var student Student
```

Output:

```
{0 "" 0 ""}
```

---

# JavaScript vs Go

## JavaScript

```javascript
const user = {
	name: "Rajnish",
	age: 25,
	course: "Go"
};

console.log(user.name);
```

---

## Go

```go
student := Student{
	Name: "Rajnish",
	Age: 25,
	Course: "Go",
}

fmt.Println(student.Name)
```

### Comparison

| JavaScript Object | Go Struct |
|-------------------|-----------|
| Dynamic properties | Fixed fields |
| Runtime flexibility | Compile-time safety |
| No field validation | Strong typing |
| Easy to modify shape | Fixed structure |

Structs provide better performance and type safety.

---

# Memory Representation

```
Student

+----------------------+

ID       : 101

Name     : Rajnish

Age      : 25

Course   : Go

+----------------------+
```

Each field occupies memory according to its type.

---

# Real-World Use Cases

Structs are used for:

- Database Models
- API Requests
- API Responses
- Configuration Files
- JSON Data
- Authentication
- Product Catalogs
- Employee Records

If you're building a backend application, you'll work with structs every day.

---

# Best Practices

- Group related data together.
- Use meaningful field names.
- Keep structs focused on one responsibility.
- Use nested structs instead of long flat structures.
- Pass pointers to large structs when modifying data.

---

# Common Mistakes

## Forgetting Field Names

Incorrect:

```go
student := Student{}
student.name = "Raj"
```

Fields are case-sensitive.

Correct:

```go
student.Name = "Raj"
```

---

## Using Too Many Fields

Avoid creating huge structs with unrelated data.

Instead, split them into smaller nested structs.

---

## Copying Large Structs

Passing large structs by value copies all fields.

Prefer pointers when updating or passing large objects.

---

# Mini Project

Build a Student Management System.

Requirements:

- Create a Student struct.
- Store multiple students in a slice.
- Display all students.
- Update student details.
- Search by ID.
- Delete a student.

Bonus:

Add an Address struct and make it nested inside Student.

---

# Interview Questions

### What is a struct?

A struct is a custom data type that groups related fields.

---

### Why use structs?

To model real-world entities and organize related data.

---

### Can structs contain other structs?

Yes.

This is called nested structs.

---

### Can structs be compared?

Yes, if all fields are comparable.

---

### Why use pointers with structs?

To avoid copying large amounts of data and allow modifications to the original struct.

---

# Key Takeaways

- Structs group related data into one custom type.
- They are used to model real-world entities.
- Fields are accessed using the dot operator.
- Structs can be nested.
- Slices of structs are common in backend applications.
- Pointers improve efficiency when working with large structs.
- Structs are one of the most frequently used features in Go.

---
 