package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/fstanis/go-hive/hive"
	"golang.org/x/crypto/ssh/terminal"
)

const (
	maxInputLength = 256
)

// Asks the user for input. If password is true, prevents echo.
func ask(message string, password bool) (result string) {
	fmt.Print(message)

	if password {
		bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
		fmt.Print("\n")
		result = string(bytePassword)
	} else {
		reader := bufio.NewReader(os.Stdin)
		result, _ = reader.ReadString('\n')
	}
	result = strings.TrimSpace(result)
	if len(result) > maxInputLength {
		result = ""
	}
	return result
}

// Asks the user for username, password and login URL.
func askCredentials() *hive.Credentials {
	var creds hive.Credentials

	for creds.Username == "" {
		creds.Username = ask("Enter username: ", false)
		if creds.Username == "" {
			fmt.Println("Username mustn't be empty.")
		}
	}

	for creds.Password == "" {
		creds.Password = ask("Enter password: ", true)
		if creds.Password == "" {
			fmt.Println("Password mustn't be empty.")
		}
	}

	for creds.URL == "" {
		creds.URL = ask("Enter login URL: ", false)
		if creds.URL == "" {
			fmt.Println("URL mustn't be empty.")
		}
	}
	return &creds
}
