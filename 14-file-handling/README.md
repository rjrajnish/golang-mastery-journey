 # File Handling in Go

File handling is an essential part of backend development. Applications often need to read configuration files, write logs, upload documents, generate reports, or process user data stored in files.

Go provides a rich standard library for working with files. The `os`, `io`, and `bufio` packages make it simple to create, read, write, update, rename, and delete files.

By mastering file handling, you'll be able to build practical applications such as log processors, note-taking tools, file upload services, and report generators.

---

# Table of Contents

1. What is File Handling?
2. Why Do We Need File Handling?
3. Creating a File
4. Writing to a File
5. Reading a File
6. Reading Line by Line
7. Appending Data
8. File Information
9. Renaming Files
10. Deleting Files
11. JavaScript vs Go
12. Real-World Use Cases
13. Best Practices
14. Common Mistakes
15. Mini Project
16. Interview Questions
17. Key Takeaways

---

# What is File Handling?

File handling is the process of creating, reading, writing, updating, renaming, and deleting files.

A file allows applications to store information permanently.

Examples:

- Configuration files
- Log files
- Images
- CSV files
- PDFs
- Reports

---

# Why Do We Need File Handling?

Imagine building a Notes application.

Without files:

```
Notes disappear after the program ends.
```

With files:

```
Notes remain saved even after restarting the application.
```

This persistence makes file handling essential for real-world software.

---

# Creating a File

Use `os.Create()`.

```go
file, err := os.Create("example.txt")
```

If the file already exists, it is truncated (its contents are cleared).

Always check the returned error and close the file when you're done.

---

# Writing to a File

Write an entire file at once:

```go
err := os.WriteFile("example.txt", []byte("Hello, Go!"), 0644)
```

The third argument specifies file permissions.

---

# Reading a File

Read the complete contents:

```go
data, err := os.ReadFile("example.txt")
```

Convert the byte slice to a string for display:

```go
fmt.Println(string(data))
```

---

# Reading Line by Line

For large files, reading line by line is more memory-efficient.

```go
scanner := bufio.NewScanner(file)

for scanner.Scan() {
	fmt.Println(scanner.Text())
}
```

This is commonly used for:

- Log files
- CSV processing
- Text parsing

---

# Appending Data

To add new content without overwriting existing data:

```go
file, err := os.OpenFile(
	"example.txt",
	os.O_APPEND|os.O_WRONLY,
	0644,
)

file.WriteString("New Line\n")
```

---

# File Information

Retrieve metadata:

```go
info, err := os.Stat("example.txt")
```

Useful methods include:

- `info.Name()`
- `info.Size()`
- `info.Mode()`
- `info.ModTime()`

---

# Renaming a File

```go
os.Rename("old.txt", "new.txt")
```

Renaming is useful for versioning, backups, or temporary file management.

---

# Deleting a File

```go
os.Remove("example.txt")
```

Always ensure the file is no longer needed before deleting it.

---

# JavaScript vs Go

## JavaScript (Node.js)

```javascript
const fs = require("fs");

fs.writeFileSync("note.txt", "Hello");
```

---

## Go

```go
os.WriteFile(
	"note.txt",
	[]byte("Hello"),
	0644,
)
```

### Comparison

| JavaScript (Node.js) | Go |
|----------------------|----|
| `fs` module | `os`, `io`, `bufio` packages |
| Synchronous & asynchronous APIs | Standard synchronous APIs |
| Dynamic typing | Static typing |
| Exceptions | Explicit error handling |

Go emphasizes explicit error checking instead of exceptions.

---

# Real-World Use Cases

File handling is used for:

- Application logs
- Configuration files
- File uploads
- Exporting reports
- Importing CSV/Excel
- Backup systems
- Image processing
- Document generation

---

# Best Practices

- Always check returned errors.
- Close files after use (`defer file.Close()` is recommended).
- Use `bufio.Scanner` for large text files.
- Avoid hard-coded file paths.
- Handle missing files gracefully.

---

# Common Mistakes

## Forgetting to Close a File

Bad:

```go
file, _ := os.Open("example.txt")
```

Good:

```go
file, err := os.Open("example.txt")
if err != nil {
	return
}
defer file.Close()
```

---

## Ignoring Errors

Bad:

```go
data, _ := os.ReadFile("example.txt")
```

Always handle the error:

```go
data, err := os.ReadFile("example.txt")
if err != nil {
	fmt.Println(err)
	return
}
```

---

## Reading Huge Files at Once

`os.ReadFile()` loads the entire file into memory.

For very large files, prefer `bufio.Scanner` or buffered readers.

---

# Mini Project

Build a **Notes Manager**.

Requirements:

- Create a notes file.
- Add new notes.
- Display all notes.
- Search notes.
- Delete the notes file.

Bonus:

Add timestamps to each note.

---

# Interview Questions

### What package is commonly used for file handling?

The `os` package, often together with `bufio` and `io`.

---

### How do you create a file?

```go
os.Create("file.txt")
```

---

### How do you append to a file?

Use `os.OpenFile()` with the `os.O_APPEND` flag.

---

### How do you check whether a file exists?

Use `os.Stat()` and inspect the returned error.

---

### Why use `bufio.Scanner`?

To read large text files efficiently, one line at a time.

---

# Key Takeaways

- Go provides excellent file handling support through the standard library.
- Always check errors returned by file operations.
- Close files after opening them.
- Use `os.ReadFile()` for small files.
- Use `bufio.Scanner` for large text files.
- File handling is fundamental for many backend applications.

---

