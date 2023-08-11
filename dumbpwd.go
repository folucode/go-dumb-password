package dumbpwd

import (
	"os"
	"slices"
	"strings"
)

func CheckPassword(p string) (r string, e bool) {
	message := "This password is too common. Please try another!"

	file, err := os.ReadFile("passwordlist.txt")

	if err != nil {
		return "An error occurred while reading the password list.", false
	}

	content := strings.Split(strings.ToLower(string(file)), "\n")

	c := slices.Contains(content, p)

	if c {
		return message, true
	}

	return "This password is good!", false
}
