package util

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Data(part int, sep string, args ...bool) []string {
	if sep == "" {
		sep = "\n"
	}

	test := false
	var b []byte
	var err error

	if len(args) > 0 {
		test = args[0]
	}

	if test {
		b, err = os.ReadFile("test_data.csv")
	} else {
		b, err = os.ReadFile("data.csv")
	}

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

func ConvertArrayOfStringsToInts(arr []string) ([]int, error) {
	var err error
	b := make([]int, len(arr))
	for i, s := range arr {
		b[i], err = strconv.Atoi(strings.Trim(s, " "))
		if err != nil {
			return nil, err
		}
	}
	return b, nil
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

func Min(values []int) (min int, e error) {
	if len(values) == 0 {
		return 0, errors.New("Cannot detect a minimum value in an empty slice")
	}

	min = values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}

	return min, nil
}

func Max(values []int) (max int, e error) {
	if len(values) == 0 {
		return 0, errors.New("Cannot detect a maximum value in an empty slice")
	}

	max = values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}

	return max, nil
}

func Contains[T comparable](s []T, e T) bool {
    for _, v := range s {
        if v == e {
            return true
        }
    }
    return false
}


// MapKeys returns a slice of all the keys in m.
// The keys are not returned in any particular order.
func MapKeys[Key comparable, Val any](m map[Key]Val) []Key {
    s := make([]Key, 0, len(m))
    for k := range m {
        s = append(s, k)
    }
    return s
}

