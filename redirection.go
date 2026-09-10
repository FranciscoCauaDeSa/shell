package main

import (
	"fmt"
	"os"
	"slices"
)

var redirectingStdout, redirectingStderr bool // This will be false by default
var stdoutRedirectFile, stderrRedirectFile *os.File

// Saving the pointer address to return to it later when redirecting
var originalStdout = os.Stdout
var originalStderr = os.Stderr

func RedirectStderr(fields []string) ([]string, error) {
	var redirectStderrIndex int

	var append bool

	redirectStderrIndex = slices.Index(fields, "2>")
	if redirectStderrIndex == -1 {
		redirectStderrIndex = slices.Index(fields, "2>>")
		append = slices.Contains(fields, "2>>")
	}

	if redirectStderrIndex == 0 {
		//TODO: turn this into an error
		return fields, fmt.Errorf("Expected a string, but found a redirection")
	}

	// Just to make sure
	if redirectStderrIndex != -1 {
		redirectingStderr = true
	}

	if len(fields)-2 < redirectStderrIndex {
		return fields, fmt.Errorf("Expected a string, but found end of the input")
	}

	var err error

	if append {
		stderrRedirectFile, err = os.OpenFile(fields[redirectStderrIndex+1], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err) // Using pacnic for now. I don't know a better way to do this.
		}
	} else {
		stderrRedirectFile, err = os.Create(fields[redirectStderrIndex+1])
		if err != nil {
			panic(err) // Using pacnic for now. I don't know a better way to do this.
		}
	}

	// Remove everything after the > or 1>
	fields = fields[:redirectStderrIndex]

	// Saving it to retrieve later
	originalStderr = os.Stderr

	// Changing it to point to the file the user inputted
	os.Stderr = stderrRedirectFile

	return fields, nil
}

func RedirectStdout(fields []string) ([]string, error) {
	var redirectStdoutIndex int

	redirectStdoutIndex = slices.Index(fields, ">")

	var append bool

	if redirectStdoutIndex == -1 {
		redirectStdoutIndex = slices.Index(fields, "1>")
	}
	if redirectStdoutIndex == -1 {
		redirectStdoutIndex = slices.Index(fields, ">>")
		append = slices.Contains(fields, ">>")
	}
	if redirectStdoutIndex == -1 {
		redirectStdoutIndex = slices.Index(fields, "1>>")
		append = slices.Contains(fields, "1>>")
	}

	// This will tell if we should append or truncate.
	// append := slices.Contains(fields, ">>")

	if redirectStdoutIndex == 0 {
		//TODO: turn this into an error
		return fields, fmt.Errorf("Expected a string, but found a redirection")
	}

	// Just to make sure
	if redirectStdoutIndex != -1 {
		redirectingStdout = true
	}

	if len(fields)-2 < redirectStdoutIndex {
		return fields, fmt.Errorf("Expected a string, but found end of the input")
	}

	var err error

	if append {
		stdoutRedirectFile, err = os.OpenFile(fields[redirectStdoutIndex+1], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err) // Using pacnic for now. I don't know a better way to do this.
		}
	} else {
		stdoutRedirectFile, err = os.Create(fields[redirectStdoutIndex+1])
		if err != nil {
			panic(err) // Using pacnic for now. I don't know a better way to do this.
		}
	}

	// Remove everything after the > or 1>
	fields = fields[:redirectStdoutIndex]

	// Saving it to retrieve later
	originalStdout = os.Stdout

	// Changing it to point to the file the user inputted
	os.Stdout = stdoutRedirectFile

	return fields, nil
}

func CloseStdoutRedirection() {
	if redirectingStdout {
		// Giving back os.Stdout it's original value
		os.Stdout = originalStdout

		redirectingStdout = false

		// Closing the file we were writing
		stdoutRedirectFile.Close()
	}
}

func CloseStderrRedirection() {
	if redirectingStderr {
		// Giving back os.Stdout it's original value
		os.Stderr = originalStderr

		redirectingStderr = false

		// Closing the file we were writing
		stderrRedirectFile.Close()
	}
}
