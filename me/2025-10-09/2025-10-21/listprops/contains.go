package listprops

// Contains liefert true, falls die Liste l
// den String x enthält.
func Contains(l []string, x string) bool {
	for i := 0; i < len(l); i++ {
		//liste durchgehen, solange wie I ist
		if l[i] == x {
			return true
		}
	}
	return false
}
