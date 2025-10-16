package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("Guten Tag")
}

func add(x int, y int) int {
	return x + y

}

func get2() (int, int) {
	return 1, 28
}
func multiply(x int, y int) (result int) {
	result = x * y
	result = result + 10
	return
}

func main() {
	sayHello()
	fmt.Println("go tutorial runde 3.22 ")
	fmt.Println(add(2, 5))
	x := 1
	x = x << 2
	fmt.Println(x)
	fmt.Println(get2())
	fmt.Println(multiply(3, 4))

}
