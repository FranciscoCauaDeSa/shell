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
