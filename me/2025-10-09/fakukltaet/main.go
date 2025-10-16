package main

import "fmt"

func main() {
	var name string = ("Mal wieder ein Go Tutorial und zwar das")

	x := 7
	x = 8
	x = x + 1
	y := 15.5
	var c complex128 = 2i + 13
	var d complex128 = 2i + 15 + c - 200

	fmt.Println(x, y, c, d)
	fmt.Println(name)

	fmt.Println("Go Tutorial 2")

}
