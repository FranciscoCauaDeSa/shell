# Go Shell Implementation

A lightweight, POSIX-style shell implementation written in Go.

---

## Features

- **Interactive REPL**: Read-Eval-Print Loop that displays a prompt (`$ `), evaluates user input, and gracefully handles EOF (Ctrl+D).
- **Builtin Commands**:
  - `exit [code]`: Terminates the shell session.
  - `echo [args...]`: Prints text and arguments to standard output.
  - `type [cmd]`: Distinguishes between shell builtins and executables located in `$PATH`.
  - `pwd`: Prints the current working directory.
  - `cd [dir]`: Changes directory with support for relative paths, absolute paths, and home directory expansion (`~`).
- **External Program Execution**: Resolves external binaries through system `$PATH` (`exec.LookPath`) and executes them with arguments (`exec.Command`).
- **Command & Quote Parsing**:
  - Handles single quotes (`'...'`) to preserve literal values.
  - Handles double quotes (`"..."`) with escape sequence handling.
  - Handles backslash escaping (`\`) for spaces and special characters.

---

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) installed (Go 1.22+ recommended).

### Running the Shell Locally

You can run the shell directly using `go run`:

```bash
go run main.go
```

---

## Example Usage

```bash
$ echo "Hello World"
Hello World 

$ type echo
echo is a shell builtin

$ type ls
ls is /usr/bin/ls

$ pwd
/home/user/shell

$ cd ~
$ pwd
/home/user

$ exit
```