# Go File Handling — Read, Write, Copy & Delete Files

## 1. Introduction

Go provides the `os` package to work with files and directories.

File handling allows us to:

* Open existing files.
* Read file contents.
* Write data into files.
* Create new files.
* Get file information.
* Read directory contents.
* Copy files.
* Delete files.

```go
import (
    "fmt"
    "os"
)
```

The `os` package provides functions such as `Open()`, `Create()`, `ReadFile()`, and `Remove()`.

---

## 2. Opening a File

Use `os.Open()` to open an existing file for reading.

```go
f, err := os.Open("./example.txt")
if err != nil {
    panic(err)
}
defer f.Close()
```

### Explanation

* `os.Open()` opens an existing file in read-only mode.
* `f` is a file handle (`*os.File`).
* `err` contains an error if opening the file fails.
* `defer f.Close()` closes the file when the function returns.

**Important:** `os.Open()` does not create a file if it doesn't exist.

### Why use `defer f.Close()`?

It ensures the file is closed when the function finishes, even if the function returns early.

Closing files releases operating system resources.

---

## 3. Getting File Information

Use the `Stat()` method to get information about an open file.

```go
f, err := os.Open("./example.txt")
if err != nil {
    panic(err)
}
defer f.Close()

fileInfo, err := f.Stat()
if err != nil {
    panic(err)
}

fmt.Println("File name:", fileInfo.Name())
fmt.Println("File size:", fileInfo.Size())
fmt.Println("File time:", fileInfo.ModTime())
fmt.Println("Is directory:", fileInfo.IsDir())
fmt.Println("File mode:", fileInfo.Mode())
```

### Common `FileInfo` methods

| Method      | Description                                    |
| ----------- | ---------------------------------------------- |
| `Name()`    | Returns the file name                          |
| `Size()`    | Returns the file size in bytes                 |
| `ModTime()` | Returns the last modification time             |
| `IsDir()`   | Checks whether the file is a directory         |
| `Mode()`    | Returns file permissions and file type         |
| `Sys()`     | Returns underlying system-specific information |

`Sys()` is platform-dependent. Its returned value may differ between operating systems.

---

## 4. Reading Files Using `Read()`

The `Read()` method reads bytes from an open file into a byte slice.

```go
f, err := os.Open("./example.txt")
if err != nil {
    panic(err)
}
defer f.Close()

buf := make([]byte, 11)

n, err := f.Read(buf)
if err != nil {
    panic(err)
}

fmt.Println("Bytes read:", n)
fmt.Println("Data:", string(buf[:n]))
```

### Explanation

```go
buf := make([]byte, 11)
```

Creates a byte slice with a length of 11.

```go
n, err := f.Read(buf)
```

Reads data into the buffer.

* `n` is the number of bytes actually read.
* `err` indicates whether an error occurred.

```go
string(buf[:n])
```

Converts only the bytes that were actually read into a string.

### Important: `Read()` may read fewer bytes

Even if your buffer has space for 100 bytes, `Read()` may return fewer bytes.

For example:

```go
buf := make([]byte, 100)
n, err := f.Read(buf)
```

The file might contain only 20 bytes, or the read might return fewer bytes even when more data is available.

For this reason, don't assume the buffer is completely filled.

Also, a file read may return both `n > 0` and a non-nil error. Process the bytes returned before handling the error.

---

## 5. Reading a Complete File Using `os.ReadFile()`

Go provides a simpler way to read an entire file.

```go
data, err := os.ReadFile("./example.txt")
if err != nil {
    panic(err)
}

fmt.Println(string(data))
```

### Explanation

`os.ReadFile()` reads the complete file and returns its contents as a byte slice.

```go
data
```

Contains the file's bytes.

```go
string(data)
```

Converts the bytes into a string.

### When should you use it?

Use `os.ReadFile()` when:

* The file is reasonably small.
* You need the complete contents in memory.
* You want simple file-reading code.

For very large files, consider streaming instead of loading everything into memory.

---

## 6. Reading a Directory

You can use `os.Open()` to open a directory and `ReadDir()` to list its entries.

```go
dir, err := os.Open(".")
if err != nil {
    panic(err)
}
defer dir.Close()

fileInfo, err := dir.ReadDir(-1)
if err != nil {
    panic(err)
}

for _, fi := range fileInfo {
    fmt.Println(fi.Name(), fi.IsDir(), fi.Type())
}
```

### Explanation

```go
os.Open(".")
```

Opens the current directory.

The dot (`.`) means the current working directory.

```go
dir.ReadDir(-1)
```

Reads all directory entries and returns them as a slice.

```go
fi.Name()
```

Returns the entry name.

```go
fi.IsDir()
```

Checks whether the entry is a directory.

```go
fi.Type()
```

Returns the file mode type information.

### Important

`ReadDir(-1)` reads all entries into memory.

For very large directories, you can use a positive number to read entries in batches.

You can also use `os.ReadDir()` when you only need a directory listing:

```go
entries, err := os.ReadDir(".")
if err != nil {
    panic(err)
}

for _, entry := range entries {
    fmt.Println(entry.Name())
}
```

---

## 7. Creating a File

Use `os.Create()` to create a file or truncate an existing file.

```go
f, err := os.Create("test.txt")
if err != nil {
    panic(err)
}
defer f.Close()
```

### Important behavior

`os.Create()`:

* Creates the file if it doesn't exist.
* Truncates the file if it already exists.
* Opens it for reading and writing.

**Warning:** If the file already contains data, `os.Create()` removes the existing contents.

---

## 8. Writing to a File

You can use `WriteString()` to write text.

```go
f, err := os.Create("test.txt")
if err != nil {
    panic(err)
}
defer f.Close()

_, err = f.WriteString("Hi Go")
if err != nil {
    panic(err)
}

_, err = f.WriteString(" Hi Python")
if err != nil {
    panic(err)
}
```

The file now contains:

```text
Hi Go Hi Python
```

Each call writes at the current file offset. It doesn't automatically add a newline.

To write a new line:

```go
_, err = f.WriteString("Hi Go\n")
```

### Writing bytes

You can also write a byte slice:

```go
data := []byte("Hello Golang")

_, err := f.Write(data)
if err != nil {
    panic(err)
}
```

`Write()` returns the number of bytes written and an error.

Always check the error when writing important data.

---

## 9. Difference Between `WriteString()` and `Write()`

| Method          | Purpose                                 |
| --------------- | --------------------------------------- |
| `WriteString()` | Writes a string                         |
| `Write()`       | Writes a byte slice                     |
| `WriteByte()`   | Writes one byte using a buffered writer |

Example:

```go
f.WriteString("Hello")
```

And:

```go
f.Write([]byte("Hello"))
```

Both write the same text to a file.

---

## 10. Copying a File Using Streaming

For large files, streaming is useful because it avoids loading the entire file into memory.

Your example uses `bufio.Reader` and `bufio.Writer`.

```go
sourceFile, err := os.Open("./example.txt")
if err != nil {
    panic(err)
}
defer sourceFile.Close()

destFile, err := os.Create("example2.txt")
if err != nil {
    panic(err)
}
defer destFile.Close()
```

Here:

* `sourceFile` is the file we want to copy.
* `destFile` is the new file.

### Using a buffered reader and writer

```go
reader := bufio.NewReader(sourceFile)
writer := bufio.NewWriter(destFile)
```

A buffered reader and writer help process data in chunks instead of making a separate underlying system call for every small operation.

Your original code reads and writes one byte at a time:

```go
for {
    b, err := reader.ReadByte()
    if err != nil {
        if err.Error() != "EOF" {
            panic(err)
        }
        break
    }

    err = writer.WriteByte(b)
    if err != nil {
        panic(err)
    }
}

err = writer.Flush()
if err != nil {
    panic(err)
}
```

### Important correction: Handle `io.EOF`

Instead of checking:

```go
err.Error() != "EOF"
```

use:

```go
errors.Is(err, io.EOF)
```

Error values should be checked using `errors.Is()` when appropriate, rather than comparing their text.

You also need to import `errors` and `io`.

### Simpler and more efficient method: `io.Copy()`

For ordinary file copying, Go provides `io.Copy()`.

```go
package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    source, err := os.Open("example.txt")
    if err != nil {
        panic(err)
    }
    defer source.Close()

    destination, err := os.Create("example2.txt")
    if err != nil {
        panic(err)
    }
    defer destination.Close()

    _, err = io.Copy(destination, source)
    if err != nil {
        panic(err)
    }

    fmt.Println("File copied successfully")
}
```

`io.Copy()` reads from the source and writes to the destination until it reaches the end of the source or encounters an error.

For most simple file-copy tasks, this is the method to learn first.

---

## 11. Why Do We Need `Flush()`?

When using `bufio.Writer`, data may stay in its internal buffer instead of immediately reaching the underlying file.

```go
writer := bufio.NewWriter(destFile)
```

After writing, call:

```go
err := writer.Flush()
if err != nil {
    panic(err)
}
```

`Flush()` sends buffered data to the underlying writer.

Without flushing, some buffered data may not be written before the program finishes.

For important data, also check the error returned by `Close()`, because a final file operation can fail.

---

## 12. Deleting a File

Your code:

```go
err := os.Remove("./example.txt")
if err != nil {
    panic(err)
}

fmt.Println("File Deleted")
```

`os.Remove()` deletes a file or an empty directory.

If the path does not exist, it returns an error.

### Important

Do not print `"File Deleted"` before checking the error.

Your code correctly checks the error first.

Also, `os.Remove()` does not recursively delete a directory and all its contents.

For removing a directory tree, Go provides:

```go
os.RemoveAll("./my-folder")
```

Be careful with `RemoveAll()`: it deletes the specified path and everything inside it.

---

## 13. Complete File Handling Example

This example demonstrates creating, writing, reading, and deleting a file.

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Create a file
    file, err := os.Create("example.txt")
    if err != nil {
        panic(err)
    }

    // Write data
    _, err = file.WriteString("Hello Golang\n")
    if err != nil {
        file.Close()
        panic(err)
    }

    // Close the file before reading it
    if err := file.Close(); err != nil {
        panic(err)
    }

    // Read the file
    data, err := os.ReadFile("example.txt")
    if err != nil {
        panic(err)
    }

    fmt.Println(string(data))

    // Delete the file
    err = os.Remove("example.txt")
    if err != nil {
        panic(err)
    }

    fmt.Println("File deleted")
}
```

Output:

```text
Hello Golang

File deleted
```

This example uses explicit `Close()` so the file is closed before it is read and deleted.

---

## 14. Common File Handling Errors

| Error              | Meaning                                                   |
| ------------------ | --------------------------------------------------------- |
| `os.ErrNotExist`   | File or directory doesn't exist                           |
| `os.ErrPermission` | Permission denied                                         |
| `io.EOF`           | End of file reached                                       |
| `os.ErrExist`      | File already exists in operations that require a new path |

For more reliable error handling, use `errors.Is()`:

```go
if errors.Is(err, os.ErrNotExist) {
    fmt.Println("File does not exist")
}
```

For example:

```go
package main

import (
    "errors"
    "fmt"
    "os"
)

func main() {
    _, err := os.Open("missing.txt")

    if errors.Is(err, os.ErrNotExist) {
        fmt.Println("File does not exist")
        return
    }

    if err != nil {
        panic(err)
    }
}
```

---

## 15. Important File Handling Methods

| Function or method | Description                                |
| ------------------ | ------------------------------------------ |
| `os.Open()`        | Open an existing file for reading          |
| `os.Create()`      | Create or truncate a file                  |
| `os.ReadFile()`    | Read an entire file                        |
| `os.WriteFile()`   | Write data to a file                       |
| `os.Remove()`      | Remove a file or empty directory           |
| `os.RemoveAll()`   | Remove a path and its contents             |
| `os.OpenFile()`    | Open or create a file with specified flags |
| `f.Stat()`         | Get file information                       |
| `f.Read()`         | Read bytes from an open file               |
| `f.Write()`        | Write bytes to an open file                |
| `f.WriteString()`  | Write a string                             |
| `f.Close()`        | Close the file                             |
| `dir.ReadDir()`    | Read directory entries                     |
| `io.Copy()`        | Copy data from a reader to a writer        |
| `writer.Flush()`   | Flush buffered data                        |

---

## 16. Pro Tips

1. **Always check errors.** File operations can fail because of missing files, permissions, or disk problems.
2. **Close files.** Use `defer file.Close()` when the file should remain open until the function returns.
3. **Use `os.ReadFile()` for small files.** It is simple and convenient.
4. **Use streaming for large files.** Avoid loading huge files entirely into memory.
5. **Prefer `io.Copy()` for file copying.** It is simpler than manually copying bytes.
6. **Remember that `os.Create()` truncates existing files.** It does not append.
7. **Use `os.OpenFile()` for more control.** It supports options such as append and exclusive creation.
8. **Check `Read()`'s byte count.** The number of bytes read may be smaller than the buffer length.
9. **Check `Flush()` errors.** Buffered writes can fail when flushed.
10. **Be careful with deletion.** Verify the path before using `os.RemoveAll()`.

---

## 17. What I Learned Today

* I learned how to open files using `os.Open()`.
* I learned how to get file information using `Stat()`.
* I learned how to read bytes using `Read()`.
* I learned how to read complete files using `os.ReadFile()`.
* I learned how to list directory entries using `ReadDir()`.
* I learned how to create files using `os.Create()`.
* I learned how to write strings and byte slices.
* I learned how to copy files using buffered I/O and `io.Copy()`.
* I learned why `Flush()` is important for buffered writers.
* I learned how to delete files using `os.Remove()`.
* I learned how to handle file errors and close files safely.

---

## 18. Quick Reference

```go
// Open
f, err := os.Open("example.txt")

// Create or truncate
f, err := os.Create("example.txt")

// Read entire file
data, err := os.ReadFile("example.txt")

// Write entire file
err = os.WriteFile("example.txt", []byte("Hello"), 0644)

// File information
info, err := f.Stat()

// Read directory
entries, err := os.ReadDir(".")

// Delete
err = os.Remove("example.txt")

// Copy
_, err = io.Copy(destination, source)

// Close
err = f.Close()
```

---

## 🧠 Mental Model

```text
             File Handling
                   |
       ┌───────────┼───────────┐
       |           |           |
      Read        Write       Manage
       |           |           |
  os.Open()    os.Create()   os.Remove()
  os.ReadFile  WriteString   os.OpenFile()
  f.Read()     f.Write()    f.Stat()
       |
       v
  Buffered I/O
       |
  bufio.Reader
  bufio.Writer
       |
  writer.Flush()
       |
       v
   io.Copy()
```

**Remember:** Small files are often easiest to handle with `os.ReadFile()` and `os.WriteFile()`. For large files, use streaming techniques such as `io.Copy()`.
