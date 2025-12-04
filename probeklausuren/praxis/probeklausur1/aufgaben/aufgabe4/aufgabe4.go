package aufgabe4

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion MaxElements.
MAX. PUNKTE: 10
ZUSATZBEDINGUNG: Die Funktion muss rekursiv sein!
*/

// MaxElements erwartet zwei int-Listen und liefert eine Liste, die an jeder Position
// jeweils das größere der beiden Elemente enthält.
// Falls eine Position nur in einer Liste vorkommt, gilt dieses Element als das größere.
func MaxElements(l1, l2 []int) []int {
	if len(l1) == 0 {
		return l2
	}
	if len(l2) == 0 {
		return l1
	}
	john := l1[0]
	if l2[0] > john {
		john = l2[0]
	}
	return append([]int{john}, MaxElements(l1[1:], l2[1:])...)
}

// 	if len(l1) == 0 {
// 		return l2
// 	}
// 	if len(l2) == 0 {
// 		return l1
// 	}
// 	//Wenn Liste 1 leer, dann gib 2 aus oder umgekehrt

// 	greater := l1[0]
// 	if l2[0] > greater {
// 		greater = l2[0]
// 	} // Liste 1 Position 0  wird zu greater, wenn greater an Position 0 kleiner ist als als Liste 2 an Position 0, dann speicher diese als neues greater.
// 	return append([]int{greater}, MaxElements(l1[1:], l2[1:])...)
// 	// Rekursiv, Prüfe greaterfür den rest der Liste
// }
