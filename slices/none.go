package slices

// None returns true if none of the elements match the given predicate.
func None[S ~[]E, E any](vals S, predicate func(E) bool) bool {
	for _, val := range vals {
		if predicate(val) {
			return false
		}
	}
	return true
}
