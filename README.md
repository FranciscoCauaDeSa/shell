# Go Shell Implementation

A lightweight, POSIX-style shell implementation written in Go featuring command execution, argument parsing, and standard stream redirections.

---

## Features

- **Interactive REPL**: Read-Eval-Print Loop that displays a prompt (`$ `), evaluates user input, and gracefully handles EOF (Ctrl+D).
- **I/O Stream Redirection**:
  - **Standard Output (stdout)**:
    - Overwrite: `>` and `1>`
    - Append: `>>` and `1>>`
  - **Standard Error (stderr)**:
    - Overwrite: `2>`
    - Append: `2>>`
  - Automatic stream restoration and safe file descriptor cleanup after execution.
- **Builtin Commands**:
  - `exit [code]`: Terminates the shell session.
  - `echo [args...]`: Prints text and arguments to standard output.
  - `type [cmd]`: Distinguishes between shell builtins and executables located in `$PATH`.
  - `pwd`: Prints the current working directory.
  - `cd [dir]`: Changes directory with support for relative paths, absolute paths, and home directory expansion (`~`).
- **External Program Execution**: Resolves external binaries through the system `$PATH` (`exec.LookPath`) and executes them with arguments (`exec.Command`).
- **Command & Quote Parsing**:
  - Handles single quotes (`'...'`) to preserve literal character values.
  - Handles double quotes (`"..."`) with escape sequence support.
  - Handles backslash escaping (`\`) for spaces and special characters.

---

## Project Structure

```text
.
├── builtin.go       # Implementation of shell built-in commands (cd, pwd, echo, type)
├── main.go          # REPL loop, input parsing, command evaluation, and execution
├── redirection.go  # File descriptor manipulation and stream redirection logic (stdout/stderr)
└── README.md
```
---

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) installed (Go 1.22+ recommended).

### Running the Shell Locally

Since the project is organized into multiple files under the `main` package, run the package directly:

```bash
go run .
```

Or build the executable binary:

```bash
go build -o go-shell .
./go-shell
```

---

## Example Usage

### 1. Basic Commands & Builtins

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
```

### 2. Output Redirection (`stdout`)

```bash
# Overwrite stdout to a file (> or 1>)
$ echo "first line" > output.txt
$ cat output.txt
first line

# Append stdout to a file (>> or 1>>)
$ echo "second line" >> output.txt
$ cat output.txt
first line
second line
```

### 3. Error Redirection (`stderr`)

```bash
# Overwrite stderr to a file (2>)
$ ls nonexistent_directory 2> error.log
$ cat error.log
ls: cannot access 'nonexistent_directory': No such file or directory

# Append stderr to an existing file (2>>)
$ ls another_missing_dir 2>> error.log
```

### 4. Quoting and Escaping

```bash
$ echo 'hello    world'
hello    world

$ echo "hello \"world\""
hello "world"

$ echo hello\ \ \ world
hello   world
```

### 5. Exiting

```bash
$ exit
```
