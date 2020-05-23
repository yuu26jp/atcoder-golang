package main

import (
	"fmt"
)

func main() {
	var count = 0
	var str string
	fmt.Scanf("%s", &str)

	for _, c := range str {
		if c == '1' {
			count++
		}
	}
	fmt.Printf("%d\n", count)
}
