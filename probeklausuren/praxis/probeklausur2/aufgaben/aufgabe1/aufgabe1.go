package aufgabe1

/* AUFGABENSTELLUNG: Vervollständigen Sie die unten stehende Funktion.
 * ERREICHBARE PUNKTE: 10
 */

// PrefixBelow10 erwartet eine Liste "list" von Zahlen und liefert
// die längste Teil-Liste, mit der "list" beginnt und die nur Zahlen < 10 enthält.
func PrefixBelow10(list []int) []int {
	if len(list) == 0 {
		return []int{}
	}

	element := list[0]
	if element >= 10 {
		return []int{}
	}
	rest := PrefixBelow10(list[1:])
	return append([]int{element}, rest...)
}
