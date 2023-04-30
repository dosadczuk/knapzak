package maps

// None returns true if none of the key-value pairs match the given predicate.
//
// If function is `nil`, returns true if there is no elements.
func None[M ~map[K]V, K comparable, V any](vals M, predicate func(K, V) bool) bool {
	if predicate == nil {
		return len(vals) == 0
	}

	for key, val := range vals {
		if predicate(key, val) {
			return false
		}
	}
	return true
}
