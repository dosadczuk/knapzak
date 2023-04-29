package slices

// Any returns true if any of the elements match the given predicate.
// If function is `nil`, returns true if there is at least one element.
func Any[S ~[]E, E any](vals S, predicate func(E) bool) bool {
	if predicate == nil {
		return len(vals) > 0
	}

	for _, val := range vals {
		if predicate(val) {
			return true
		}
	}
	return false
}
