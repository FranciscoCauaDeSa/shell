package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// All builtin functions are in this map for fast search.
var builtins = map[string]struct{}{
	"echo": {},
	"exit": {},
	"type": {},
	"pwd":  {},
	"cd":   {},
}

func Echo(fields []string) {
	for _, word := range fields[1:] {
		fmt.Print(word + " ")
	}
	fmt.Println() // Break line after printing everything
}

func ChangeDirectory(fields []string) {
	if len(fields) < 2 { // Prevent index out of bounds
		if wd, err := os.Getwd(); err == nil {
			fmt.Println(wd) // Just print the current wd if the user did not give arguments
		}
		return
	}

	var destiny string
	if fields[1] == "~" {
		destiny = GetHomeVariable()
	} else {
		destiny = fields[1]
	}
	if err := os.Chdir(destiny); err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", fields[1])
	}
}

func PrintWorkingDirectory() {
	if wd, err := os.Getwd(); err == nil {
		fmt.Println(wd)
	} else {
		log.Fatal("Could not get current working directory.")
	}
}

func Type(fields []string) {
	for _, arg := range fields[1:] {
		// Check if arg exists in builtins
		if _, ok := builtins[arg]; ok {
			fmt.Println(arg + " is a shell builtin")
			continue
		}

		// Check if arg exists in path
		fullpath, err := exec.LookPath(arg)
		if err == nil {
			fmt.Println(arg + " is " + fullpath)
			continue
		}

		// Arg not found
		fmt.Println(arg + ": not found")
	}
}
