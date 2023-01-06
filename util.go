package util

import (
	"fmt"
	"os"
	"strings"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func Data(part int, sep string) []string {
	if sep == "" {
		sep = "\n"
	}

	b, err := os.ReadFile("data.csv")

	if err != nil {
		fmt.Print(err)
	}

	str := string(b)

	arr := strings.Split(str, sep)

	return arr
}
