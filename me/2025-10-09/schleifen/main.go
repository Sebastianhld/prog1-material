package main

import "fmt"

func main() {

	x := 1
	if x == 1 {
		fmt.Println("x ist 1")
	}

	if y := 2; y < 3 {
	} else if y < 4 {
	} else {

		fmt.Println("y ist 4 oder größer")
	}

	switch name := "Horst"; name {
	case "Horst":
		fmt.Println("Hallo Horst")
	case "You":
		fmt.Println("Wer You?") // das erste passende case wird ausgeführt, der Rest übersprungen
	default:
		fmt.Println("Hallo Unbekannter") //wird ausgeführt, wenn kein case passt

	}

	switch {
	case 1 > 0:
		fmt.Println("passt, weitermachen")
	case 1 < 0:
		fmt.Println("you have a bug in your code")
	}

	for i := 0; i < 10+1; i++ { //countinloop
		fmt.Println(i)
	}

	cake := 1
	for cake > 27 { //while schleife
		cake += cake * 2
		fmt.Println("noch nicht genug Kuchen")
	}

	for { //while true schleife
		break //endlos schleife, die sofort wieder abgebrochen wird

	}

	fmt.Println("Tutorial Schleifen 5")
}
