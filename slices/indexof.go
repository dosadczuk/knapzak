package slices

// IndexOf returns index of the first element matching the target.
// Returns -1 if not in the elements.
func IndexOf[S ~[]E, E comparable](vals S, target E) int {
	for i, val := range vals {
		if val == target {
			return i
		}
	}
	return -1
}
