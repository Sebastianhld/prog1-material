package aufgabe1

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ShortestAbc.
MAX. PUNKTE: 10
*/

// ShortestAbc erwartet eine Liste von Strings und liefert
// das kürzeste Element, das mit der Buchstabenfolge "abc" beginnt.
// Liefert den leeren String, falls es kein solches Element gibt.
//
// Hinweis: Die Funktion muss nur mit kurzen Strings der Länge < 100 funktionieren.

func ShortestAbc(list []string) string {
	empty := ""
	for _, s := range list {
		if len(s) >= 3 && s[:3] == "abc" {
			if empty == "" || len(s) < len(empty) {
				empty = s
			}
		}
	}
	return empty
}

// 	short := ""
// 	for _, s := range list {
// 		if len(s) >= 3 && s[:3] == "abc" {
// 			if short == "" || len(s) < len(short) {
// 				short = s
// 			}
// 		}

// 	}
// 	return short
// }
