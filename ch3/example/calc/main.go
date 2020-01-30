package main

import (
	"fmt"
	"math"
)

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func itob(i int) bool {
	return i != 0
}

func main() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}

	fmt.Printf("%08b\n", x<<1)
	fmt.Printf("%08b\n", x>>1)

	for x := 0; x < 10; x++ {
		fmt.Printf("x = %d, EX = %.4f\n", x, math.Exp(float64(x)))
	}

	var z float64

	fmt.Println(z, -z, 1/z, -1/z)

	fmt.Println(btoi(true))
	fmt.Println(itob(1))

	s := "你好，中国"
	fmt.Println(len(s))
	fmt.Println(s[0])
}
