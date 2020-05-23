package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	n = n % 10
	if n == 3 {
		fmt.Print("bon\n")
	} else if n <= 1 || n == 6 || n == 8 {
		fmt.Print("pon\n")
	} else {
		fmt.Print("hon\n")
	}
}
