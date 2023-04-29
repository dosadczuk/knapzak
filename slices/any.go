package slices

// Any returns true if any of the elements match the given predicate.
func Any[S ~[]E, E any](vals S, predicate func(E) bool) bool {
	for _, val := range vals {
		if predicate(val) {
			return true
		}
	}
	return false
}
