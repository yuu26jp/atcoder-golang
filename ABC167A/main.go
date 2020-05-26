package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, t string
	fmt.Scanf("%s", &s)
	fmt.Scanf("%s", &t)

	if strings.HasPrefix(t, s) {
		fmt.Print("Yes\n")
	} else {
		fmt.Print("No\n")
	}
}
