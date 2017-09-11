package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := make(map[string][]string)

	for _, file := range os.Args[1:] {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "read file error %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
			files[line] = append(files[line], file)
		}
	}

	for line, n := range counts {
		if n > 1 {
			filenames := files[line]
			fmt.Printf("%d\t%s\t%v", n, line, filenames)
		}
	}
}
