 # Maps in Go

Maps are one of the most powerful and frequently used data structures in Go. They allow you to store data as **key-value pairs**, making it easy to retrieve, update, and organize information efficiently.

If you've worked with JavaScript, Go maps will feel familiar because they are similar to JavaScript objects or the ES6 `Map` type. However, Go maps are strongly typed, making your programs safer and easier to maintain.

Maps are used in almost every Go backend application, from storing user information to handling configuration, caching, and JSON data.

---

# Table of Contents

1. What is a Map?
2. Why Do We Need Maps?
3. Creating Maps
4. Accessing Values
5. Adding and Updating Data
6. Deleting Keys
7. Checking Key Existence
8. Looping Through Maps
9. The `make()` Function
10. Nested Maps
11. Nil Maps
12. JavaScript vs Go
13. Real-World Use Cases
14. Best Practices
15. Common Mistakes
16. Mini Project
17. Interview Questions
18. Key Takeaways

---

# What is a Map?

A map is a collection of **key-value pairs**.

Each key must be unique, and each key points to a value.

Think of a map like a dictionary.

```
Key            Value

name     --->  Rajnish

age      --->  25

city     --->  Ghaziabad
```

Example:

```go
student := map[string]string{
    "name": "Rajnish",
    "city": "Ghaziabad",
}
```

---

# Why Do We Need Maps?

Suppose you're storing user information.

Using variables:

```go
name := "Rajnish"
city := "Ghaziabad"
course := "Go"
```

This becomes difficult to manage.

Using a map:

```go
student := map[string]string{
    "name": "Rajnish",
    "city": "Ghaziabad",
    "course": "Go",
}
```

Everything is organized in one place.

---

# Creating a Map

There are two common ways.

## Method 1

```go
student := map[string]string{
    "name": "Rajnish",
}
```

---

## Method 2

Using `make()`.

```go
student := make(map[string]string)
```

Then add values:

```go
student["name"] = "Rajnish"
```

---

# Accessing Values

Retrieve data using the key.

```go
fmt.Println(student["name"])
```

Output

```
Rajnish
```

---

# Adding New Data

```go
student["city"] = "Ghaziabad"
```

Maps grow dynamically.

---

# Updating Existing Values

```go
student["city"] = "Noida"
```

The previous value is replaced.

---

# Deleting Keys

Use the built-in `delete()` function.

```go
delete(student, "city")
```

If the key doesn't exist, Go does nothing.

---

# Checking if a Key Exists

```go
value, exists := student["course"]
```

If `exists` is `true`, the key was found.

This is much safer than assuming a key is always present.

---

# Looping Through Maps

Use `range`.

```go
for key, value := range student {

    fmt.Println(key, value)

}
```

> **Note:** Go does not guarantee iteration order. The order of keys may be different each time you run the program.

---

# Map Length

Use `len()`.

```go
len(student)
```

Returns the total number of key-value pairs.

---

# Nested Maps

Maps can contain other maps.

```go
employee := map[string]map[string]string{
    "emp1": {
        "name": "Ravi",
        "city": "Delhi",
    },
}
```

Useful for representing structured data.

---

# Nil Maps

A map can be `nil`.

```go
var users map[string]string
```

Reading from a nil map is safe.

Writing to a nil map causes a panic.

Always initialize a map before adding values.

---

# JavaScript vs Go

## JavaScript

```javascript
const user = {
  name: "Rajnish",
  city: "Ghaziabad"
};

console.log(user.name);
```

---

## Go

```go
user := map[string]string{
    "name": "Rajnish",
    "city": "Ghaziabad",
}

fmt.Println(user["name"])
```

### Comparison

| JavaScript Object | Go Map |
|-------------------|---------|
| Dynamic keys | Strongly typed keys |
| Dynamic values | Strongly typed values |
| Dot notation supported | Use square brackets |
| Flexible | Safer and more predictable |

---

# Memory Representation

```
Map

+--------+-----------+

name ----> Rajnish

city ----> Ghaziabad

course --> Go

+--------+-----------+
```

Unlike arrays and slices, maps are hash tables internally, allowing fast lookups.

---

# Real-World Use Cases

Maps are widely used for:

- User profiles
- Configuration settings
- HTTP headers
- JSON objects
- Caching
- Session storage
- API query parameters
- Product catalogs

---

# Best Practices

- Use meaningful key names.
- Check if a key exists before using it.
- Initialize maps with `make()` when adding values dynamically.
- Don't rely on iteration order.
- Keep value types consistent.

---

# Common Mistakes

## Writing to a Nil Map

```go
var users map[string]string

users["name"] = "Rajnish"
```

This causes a runtime panic.

Correct:

```go
users := make(map[string]string)
```

---

## Assuming Order

Incorrect:

```go
for key := range student {
    // Expecting keys in insertion order
}
```

Go maps are unordered.

---

## Ignoring Key Existence

Bad:

```go
fmt.Println(student["course"])
```

Better:

```go
value, ok := student["course"]

if ok {
    fmt.Println(value)
}
```

---

# Mini Project

Build a Student Directory.

Requirements:

- Add students
- Update student details
- Delete students
- Search by roll number
- Display all students

Use maps to store the data.

---

# Interview Questions

### What is a map?

A map stores key-value pairs.

---

### Are maps ordered?

No.

Iteration order is not guaranteed.

---

### Can duplicate keys exist?

No.

Keys must be unique.

---

### What happens when accessing a missing key?

Go returns the zero value for the value type.

Use the `value, ok := map[key]` pattern to distinguish between a missing key and an actual zero value.

---

### Can you write to a nil map?

No.

Writing to a nil map causes a panic.

---

# Key Takeaways

- Maps store data as key-value pairs.
- Keys must be unique.
- Maps are dynamically sized.
- Use `make()` to initialize writable maps.
- Check for key existence using the comma-ok idiom.
- Don't rely on iteration order.
- Maps are one of the most commonly used data structures in Go.

---
 