package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	add := ""
	startAdd := time.Now()
	for _, arg := range os.Args {
		add += arg
		add += " "
	}
	timeAdd := time.Since(startAdd).Seconds()
	fmt.Println(timeAdd)

	startJoin := time.Now()
	strings.Join(os.Args, " ")
	timeJoin := time.Since(startJoin).Seconds()
	fmt.Println(timeJoin)
}
