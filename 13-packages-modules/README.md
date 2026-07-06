 # Packages & Modules in Go

As applications grow, keeping all code inside a single file becomes difficult. Packages and modules help organize code into reusable, maintainable, and scalable components.

Almost every production Go application uses packages and modules to separate responsibilities such as API handling, business logic, database access, authentication, and configuration.

Understanding packages is the first step toward building real-world Go applications.

---

# Table of Contents

1. What is a Package?
2. What is a Module?
3. Why Do We Need Packages?
4. Standard Library Packages
5. Importing Packages
6. Exported vs Unexported Identifiers
7. Package Naming Rules
8. Using go.mod
9. JavaScript vs Go
10. Project Structure
11. Real-World Use Cases
12. Best Practices
13. Common Mistakes
14. Mini Project
15. Interview Questions
16. Key Takeaways

---

# What is a Package?

A package is a collection of related Go source files that work together to provide specific functionality.

Think of a package as a toolbox.

For example:

- `fmt` → Printing output
- `math` → Mathematical operations
- `strings` → String manipulation
- `time` → Date and time utilities

Example:

```go
import "fmt"
```

Now you can use:

```go
fmt.Println("Hello")
```

---

# What is a Module?

A module is a collection of Go packages managed together.

Every Go project should have a `go.mod` file.

Example:

```
go mod init github.com/username/go-learning
```

The `go.mod` file records:

- Module name
- Go version
- Project dependencies

---

# Why Do We Need Packages?

Imagine writing everything inside one file.

```
main.go

10000+ lines
```

Difficult to:

- Read
- Maintain
- Test
- Reuse

Instead:

```
Project

api/

database/

models/

services/

config/
```

Everything becomes organized.

---

# Standard Library Packages

Go ships with a powerful standard library.

Common packages:

| Package | Purpose |
|----------|---------|
| fmt | Printing |
| math | Mathematics |
| strings | String manipulation |
| strconv | Type conversion |
| time | Date & Time |
| os | File system |
| net/http | HTTP Server |
| encoding/json | JSON Handling |
| sync | Concurrency |
| context | Request lifecycle |

---

# Importing Packages

```go
import (
	"fmt"
	"math"
	"strings"
)
```

You can import multiple packages using parentheses.

---

# Exported vs Unexported Identifiers

Go controls visibility using capitalization.

### Exported

```go
func Add() {}
```

Accessible from other packages.

---

### Unexported

```go
func add() {}
```

Accessible only within the same package.

This simple rule replaces `public` and `private` keywords used in other languages.

---

# Package Naming Rules

- Use short names.
- Use lowercase.
- Avoid underscores.
- Avoid generic names like `utils` when possible.

Good examples:

- auth
- config
- user
- payment

---

# Using go.mod

Create a module:

```bash
go mod init github.com/username/go-learning
```

Download dependencies:

```bash
go mod tidy
```

View dependencies:

```bash
go list -m all
```

---

# JavaScript vs Go

## JavaScript

```javascript
import express from "express";
```

---

## Go

```go
import "fmt"
```

### Comparison

| JavaScript | Go |
|------------|----|
| npm | go modules |
| package.json | go.mod |
| import/export | import + exported identifiers |
| Default export | No |
| Named export | Exported identifiers |

---

# Recommended Project Structure

A typical Go backend project:

```
project/

cmd/
internal/
api/
handlers/
services/
repository/
models/
config/
middlewares/
utils/
pkg/
go.mod
main.go
```

Each folder has a clear responsibility.

---

# Real-World Use Cases

Packages help separate concerns:

- API handlers
- Database queries
- Authentication
- Payment processing
- Logging
- Email service
- Configuration

This makes applications easier to test and maintain.

---

# Best Practices

- Keep packages focused on one responsibility.
- Use meaningful package names.
- Export only what is necessary.
- Prefer composition over deeply nested packages.
- Keep `main.go` small by moving logic into packages.

---

# Common Mistakes

## Putting Everything in main.go

Bad:

```
main.go

5000+ lines
```

Better:

```
main.go
services/
handlers/
models/
```

---

## Exporting Everything

Not every function needs to be public.

Export only what other packages need.

---

## Poor Package Names

Avoid:

- common
- helper
- misc

Prefer names that describe the package's purpose.

---

# Mini Project

Build a Calculator application.

Suggested structure:

```
calculator/

main.go

math/
    add.go
    subtract.go
```

Call functions from the `math` package in `main.go`.

*(We'll build this structure in a later project after covering package fundamentals.)*

---

# Interview Questions

### What is a package?

A package is a collection of related Go files.

---

### What is a module?

A module is a collection of packages managed together using `go.mod`.

---

### What is go.mod?

It defines the module name, Go version, and dependencies.

---

### How do you export a function?

Start its name with a capital letter.

Example:

```go
func Add() {}
```

---

### Can packages have the same name?

Yes, but within a project they should be organized to avoid import conflicts.

---

# Key Takeaways

- Packages organize related code.
- Modules manage entire Go projects.
- `go.mod` tracks dependencies.
- Capitalized identifiers are exported.
- Use packages to build maintainable applications.
- Every production Go application relies on packages and modules.

---

