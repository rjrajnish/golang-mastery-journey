Understanding Pointers in Go
Introduction

Pointers are one of the most important concepts in Go programming.

If you're coming from JavaScript, this topic may initially feel unfamiliar because JavaScript manages memory automatically. In Go, pointers give you controlled access to memory, allowing you to write faster and more efficient programs.

Understanding pointers is essential for building production-grade Go applications, especially when working with large data structures, APIs, databases, and concurrent systems.

What is a Pointer?

A pointer is a variable that stores the memory address of another variable.

Instead of holding the actual value, it points to where that value is stored in memory.

Example:

number := 100
pointer := &number

Here:

number stores the value 100
pointer stores the memory address of number
Why Do We Need Pointers?

Imagine you have a large User struct with dozens of fields.

Passing it by value creates a complete copy every time a function is called.

This:

Uses more memory
Slows down execution
Creates unnecessary copies

Pointers solve this by passing the memory address instead of copying the data.

Benefits include:

Better performance
Lower memory usage
Ability to modify the original value
Efficient handling of large structs
When Should You Use Pointers?

Use pointers when:

You want to modify the original variable.
You're working with large structs.
You want to avoid copying data.
You're sharing data between functions.
You're implementing methods that update object state.

Avoid pointers for small values like integers or booleans unless you specifically need shared access or optional (nil) values.

Where Are Pointers Used in Real Projects?

Pointers appear throughout Go applications.

Database Operations

Database libraries often scan query results into variables using pointers.

var name string
row.Scan(&name)
JSON Decoding
json.Unmarshal(data, &user)
HTTP Request Parsing

Frameworks bind request bodies into structs using pointers.

c.BindJSON(&user)
Updating Records
func UpdateUser(user *User)
How Do Pointers Work?

Suppose we declare:

number := 100

Memory might look like this (illustrative):

Address        Value
0x1001  ----> 100

Now create a pointer:

pointer := &number

The pointer stores the address:

pointer
   |
   v
0x1001 ----> 100

To read the value stored at that address:

fmt.Println(*pointer)

The * operator dereferences the pointer, meaning "go to this address and read the value."

Understanding & and *
Address Operator (&)

Returns the memory address of a variable.

age := 25
fmt.Println(&age)
Dereference Operator (*)

Accesses the value stored at a memory address.

fmt.Println(*pointer)
Pass by Value vs Pass by Pointer
Pass by Value
func update(age int) {
    age = 30
}

The function receives a copy.

Original value remains unchanged.

Pass by Pointer
func update(age *int) {
    *age = 30
}

The function updates the original variable directly.

JavaScript vs Go
JavaScript
let user = {
    name: "Raj"
};

let anotherUser = user;

anotherUser.name = "Rajnish";

console.log(user.name);

Output:

Rajnish

Objects in JavaScript are reference types.

Go
type User struct {
	Name string
}

user := User{Name:"Raj"}

updateUser(&user)

Go makes references explicit using pointers.

This explicitness improves readability and helps developers understand when data is shared.

Pointer Receivers

Methods can use pointer receivers to modify structs.

func (u *User) UpdateName(name string) {
	u.Name = name
}

Without a pointer receiver, the method would modify only a copy.

Nil Pointers

Pointers can be nil.

var ptr *int

Always check before dereferencing:

if ptr != nil {
	fmt.Println(*ptr)
}

Dereferencing a nil pointer causes a runtime panic.

Common Mistakes
Forgetting to Dereference

Incorrect:

pointer = 20

Correct:

*pointer = 20
Dereferencing a Nil Pointer
var ptr *int
fmt.Println(*ptr)

This will panic.

Best Practices
Use pointers for large structs.
Use value types for small immutable data.
Check for nil before dereferencing.
Use pointer receivers when methods modify data.
Keep pointer usage intentional and readable.
Mini Project

Build a Student Management System.

Requirements:

Create a Student struct.
Add methods to update student information using pointer receivers.
Print the student before and after updates.
Add validation to prevent empty names.
Key Takeaways
A pointer stores the memory address of another variable.
& returns an address.
* accesses the value at an address.
Pointers avoid unnecessary copying.
Pointer receivers allow methods to modify structs.
Always check for nil before dereferencing.
Understanding pointers is essential for writing efficient and idiomatic Go code.
 