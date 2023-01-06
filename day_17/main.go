
package main

import (
    "aoc/util"
    "os"
    "fmt"
)

func main() {
	data := util.Data(1, "\n")
	data = data[:len(data)-1]

	if len(os.Args) < 3 {
		panic("Please provide puzzle part index as parameter `go run main.go -- 1`")
	}

	part := os.Args[2]

    fmt.Printf("Run part: %s\n", part)

	sum := 0

    for _,val := range data {
	    fmt.Println(val)

        if sum == 3 {
            break
        }

        sum++
    }

}
