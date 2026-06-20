 # Constants in Go

## What are Constants?

Constants are fixed values that cannot be changed during program execution.

Once a constant is assigned a value, it remains the same throughout the program.

Example:

```go
const AppName = "Go Learning"
```

Unlike variables, constants are immutable.

---

## Why Do We Use Constants?

Constants help us store values that should never change.

Examples:

* Application Name
* Version Number
* Tax Percentage
* Pi Value
* API Base URLs
* Configuration Values

Using constants makes code safer and easier to maintain.

---

## When Should We Use Constants?

Use constants when a value:

* Never changes
* Is used multiple times
* Represents a fixed business rule
* Improves code readability
 
### Key Differences

| Feature             | JavaScript | Go           |
| ------------------- | ---------- | ------------ |
| Constant Keyword    | const      | const        |
| Static Typing       | No         | Yes          |
| Compile-Time Checks | No         | Yes          |
| Performance         | Runtime    | Compile Time |

Go validates constants during compilation, making applications safer and faster.

---

## Advantages of Constants

* Prevent accidental changes
* Improve readability
* Simplify maintenance
* Reduce bugs
* Improve code consistency

 
## Key Takeaways

* Constants store fixed values.
* Constants cannot be modified after declaration.
* Use constants for configuration and business rules.
* Constants improve safety and maintainability.
* Use grouped constants for better organization.
