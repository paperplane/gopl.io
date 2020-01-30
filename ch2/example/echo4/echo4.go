package echo4

import (
	"flag"
	"fmt"
	"strings"

	"github.com/paperplane/gopl.io/ch2/example/tempconv"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

var p = f()

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}

func newInt() *int {
	return new(int)
}

func newInt2() *int {
	var dummy int
	return &dummy
}

func delta(new, old int) int {
	return new - old
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}

	var x, y int
	fmt.Println(&x == &y, &x == &x, &x == nil)

	fmt.Println(f() == f())
	fmt.Println(p == f())

	v := 1
	incr(&v)
	fmt.Println(incr(&v))

	p := newInt()
	q := newInt2()

	fmt.Println(p, q, p == q)

	fmt.Printf("%g\n", tempconv.BoilingC-tempconv.FreezingC)
	boilingF := tempconv.CToF(tempconv.BoilingC)
	fmt.Printf("%g\n", boilingF-tempconv.CToF(tempconv.FreezingC))
	fmt.Printf("%g\n", tempconv.FToC(boilingF)-tempconv.FreezingC)

	var c tempconv.Celsius
	var f tempconv.Fahrenheit
	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	fmt.Println(c == tempconv.Celsius(f))
	fmt.Println(c.String())
}
