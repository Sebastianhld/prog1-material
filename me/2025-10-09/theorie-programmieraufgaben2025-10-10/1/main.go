package main

import "fmt"

func Foo2(n int) { // n sagt, wie oft die Schleife läuft
	s := "*"
	for i := 0; i < n; i++ {
		fmt.Println(s)
		s += s
	}
}
func main() {
	Foo2(10) // Sagt, wie oft Schleife läuft

}
