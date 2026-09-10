package main

import (
	"bufio"
	"fmt"
	"github.com/chzyer/readline"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"unicode"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$ ")

		// Handle EOF (Ctrl+D or closed input stream)
		if !scanner.Scan() {
			break
		}

		line := strings.TrimFunc(scanner.Text(), unicode.IsSpace)
		if line == "" {
			continue
		}

		fields := TreatString(line)

		// Redirecting stderr
		if slices.Contains(fields, "2>") || slices.Contains(fields, "2>>") {
			var err error
			fields, err = RedirectStderr(fields)
			if err != nil {
				fmt.Fprintln(originalStderr, err)
				continue
			}
		}

		// TODO: make it the same as above.
		if slices.Contains(fields, ">") || slices.Contains(fields, "1>") || slices.Contains(fields, ">>") || slices.Contains(fields, "1>>") {
			var err error
			fields, err = RedirectStdout(fields)
			if err != nil {
				fmt.Fprintln(originalStdout, err)
				continue
			}
		}

		switch fields[0] {
		case "exit":
			os.Exit(0)
		case "echo":
			Echo(fields)
		case "cd":
			ChangeDirectory(fields)
		case "pwd":
			PrintWorkingDirectory()
		case "type":
			Type(fields)
		default:
			// Check if arg exists in path
			fullpath, err := exec.LookPath(fields[0])
			if err == nil {
				cmd := exec.Command(
					filepath.Base(fullpath),
					fields[1:]...,
				)

				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				_ = cmd.Run()
			} else {
				// Arg not found
				fmt.Println(fields[0] + ": not found")
			}
		}

		// Close their files if they are open.
		CloseStdoutRedirection()
		CloseStderrRedirection()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func TreatString(line string) []string {
	fields := make([]string, 0)

	// This will the contain the character we're "inside"
	var specialCharacter string
	scaping := false

	tempField := ""

	for _, char := range line {

		// If the special character is a backslash
		// We'll add the next character literally to the tempfield
		if specialCharacter == "\\" {
			specialCharacter = "" // Reset variable

			tempField += string(char)
			continue
		}

		if specialCharacter == "\"" {
			if scaping {
				tempField += string(char)
				scaping = false
				continue
			}
		}

		switch char {

		// Space Character
		case ' ':
			// If we're inside a single quote, then ignore the space
			if specialCharacter != "" {
				tempField += string(char)
				continue
			}

			// If tempField is empty we just ignore the space
			if strings.TrimSpace(tempField) == "" {
				continue
			}

			// Append to the fields, since a space marks a separation
			fields = append(fields, tempField)

			tempField = "" // Reset variable
			continue

		// Single Quote Character (Apostrophe)
		case '\'':
			if char == '\'' {
				if specialCharacter == "\"" {
					tempField += string(char)
					continue
				}

				if specialCharacter != "'" {
					// Set the variable to true because we just entered a single quote
					specialCharacter = "'"
					continue
				}

				// At this point we are exiting a single quote
				specialCharacter = ""
				continue
			}

		// Double Quote Character
		case '"':
			if specialCharacter == "'" {
				tempField += string(char)
				continue
			}

			if specialCharacter != "\"" {
				// Set the variable to true because we just entered a quote quote
				specialCharacter = "\""
				continue
			}

			// At this point we are exiting a single quote
			specialCharacter = ""
			continue

		// Blackslash character
		case '\\':
			if specialCharacter == "\"" {
				scaping = true
				continue
			}

			if specialCharacter == "'" {
				tempField += string(char)
				continue
			}

			if char == '\\' {
				specialCharacter = "\\"
				continue
			}

		// Any other character
		default:
			// Add the current character to the temporary field
			// if they're not found in the switch
			tempField += string(char)
		}

	}

	// Add what's still left in tempField after getting out of the for loop
	if tempField != "" {
		fields = append(fields, tempField)
	}

	return fields
}

func GetHomeVariable() (home string) {
	if runtime.GOOS == "linux" {
		return os.Getenv("HOME") // Linux
	} else {
		return os.Getenv("USERPROFILE") // Windows
	}
}
