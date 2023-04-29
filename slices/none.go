package slices

// None returns true if none of the elements match the given predicate.
//
// If function is `nil`, returns true if there is no elements.
func None[S ~[]E, E any](vals S, predicate func(E) bool) bool {
	if predicate == nil {
		return len(vals) == 0
	}

	for _, val := range vals {
		if predicate(val) {
			return false
		}
	}
	return true
}
