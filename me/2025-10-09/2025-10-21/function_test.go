package varnames

import "fmt"

func Example_function() {
	fmt.Println(Baz(4))
	//Output:
}

func Baz(x int) int {
	if x == 0 {
		return 0

	}
	return Baz(x-1) * x // 4*3*2*1
}
