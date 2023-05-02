package algorithms

// BinarySearch returns index of the first element matching the target.
// The slice is expected to be sorted in ascending order, otherwise the
// result is undefined.
func BinarySearch[S ~[]E, E orderable](vals S, target E) (idx int, found bool) {
	n := len(vals)

	l, r := 0, n
	for l < r {
		m := l + (r-l)/2

		if vals[m] < target {
			l = m + 1 // search in right half
		} else {
			r = m // search in left half
		}
	}
	return l, l < n && vals[l] == target
}

// BinarySearchFunc returns index of the first element matching the given
// compare function. The slice is expected to be sorted in ascending order,
// otherwise the result is undefined.
//
// Function compares two arguments for order and returns:
//
//	 0 - if the arguments are equal,
//	-1 - if the first argument is less than the second,
//	+1 - if the second argument is less than the first.
func BinarySearchFunc[S ~[]E, E, T any](vals S, target T, cmp func(E, T) int) (idx int, found bool) {
	n := len(vals)

	l, r := 0, n
	for l < r {
		m := l + (r-l)/2

		if cmp(vals[m], target) < 0 {
			l = m + 1 // search in right half
		} else {
			r = m // search in left half
		}
	}
	return l, l < n && cmp(vals[l], target) == 0
}
