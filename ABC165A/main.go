package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b, k int
	fmt.Scanf("%d", &k)
	fmt.Scanf("%d %d", &a, &b)

	for i := a; i <= b; i++ {
		if i % k == 0 {
			fmt.Print("OK\n")
			os.Exit(0)
		}
	}
	fmt.Print("NG\n")
}