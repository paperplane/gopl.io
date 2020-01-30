package main

import (
	"fmt"
	"strings"
)

func comma(s string) string {
	a := strings.Trim(s, "de")
	n := len(a)
	if n < 3 {
		return a
	}
	return a[:n-3] + "," + a[n-3:]
}

func main() {
	a := "124345dede"
	fmt.Println(comma(a))
}
