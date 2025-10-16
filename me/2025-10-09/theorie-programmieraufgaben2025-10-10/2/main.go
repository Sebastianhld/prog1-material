package main

import "fmt"

func main() {
	x := 8
	y := 7
	k := x + 2*y - 2

	fmt.Println(x, y, k)
	fmt.Printf("%d-%d-%d\n", x, y, k)
	fmt.Println(k + 2 - x)
	fmt.Println(float64(k) / float64(x)) // Geteilte Ganzzahl ergebnis als gleitkommerzahl
}
