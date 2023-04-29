package slices

// IndexOf returns index of target, or -1 if not in the elements.
func IndexOf[S ~[]E, E comparable](vals S, target E) int {
	for i, val := range vals {
		if val == target {
			return i
		}
	}
	return -1
}
