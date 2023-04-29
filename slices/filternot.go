package slices

// FilterNot returns only elements not matching the given predicate.
func FilterNot[S ~[]E, R []E, E any](vals S, predicate func(E) bool) R {
	var out R
	for _, val := range vals {
		if !predicate(val) {
			out = append(out, val)
		}
	}
	return out
}
