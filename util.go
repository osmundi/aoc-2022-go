package util

import (
	"fmt"
	"os"
	"strings"
	"unicode"
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

func RuneElement(str string, idx int) rune {
	r := []rune(str)
	return r[idx]
}

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func AddToArrayIndex(arr []rune, index int, value rune) []rune {
	arr = append(arr, 0)
	copy(arr[index+1:], arr[index:])
	arr[index] = value
	return arr
}

func ReverseArray[T any](arr []T) []T {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}


// https://codereview.stackexchange.com/questions/223438/check-whether-a-string-has-all-unique-characters-time-efficiency
func Unique(arr string) bool {
	m := make(map[rune]bool)
	for _, i := range arr {
		_, ok := m[i]
		if ok {
			return false
		}

		m[i] = true
	}

	return true
}
