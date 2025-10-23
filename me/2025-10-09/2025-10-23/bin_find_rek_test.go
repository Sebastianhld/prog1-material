package searching

import "fmt"

func ExampleBinFindRek() {
	l1 := []int{1, 6, 8, 10, 25, 42, 125, 277}
	fmt.Println(BinFindRek(l1, 8))
	fmt.Println(BinFindRek(l1, 277))
	fmt.Println(BinFindRek(l1, 3))

	// Output:
	// 2
	// 7
	//-1
}
