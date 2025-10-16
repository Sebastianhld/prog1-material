package main

import (
	"fmt"
	"strings"
)

func Foo3(n int) {
	if n == 0 {
		return
	}
	Foo3(n - 1)
	fmt.Println(strings.Repeat("*", n))
}

func main() {
	Foo3(10)
}
