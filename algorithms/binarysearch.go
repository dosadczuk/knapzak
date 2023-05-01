package algorithms

import (
	"golang.org/x/exp/constraints"
)

// BinarySearch returns index of the first element matching the target.
// The slice is expected to be sorted in ascending order, otherwise the
// result is undefined.
func BinarySearch[S ~[]E, E constraints.Ordered](vals S, target E) (idx int, found bool) {
	l, r := 0, len(vals)-1
	for l <= r {
		m := l + (r-l)/2

		if val := vals[m]; val == target {
			return m, true // found
		} else if val < target {
			l = m + 1 // search in right half
		} else {
			r = m - 1 // search in left half
		}
	}
	return
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
	l, r := 0, len(vals)-1
	for l <= r {
		m := l + (r-l)/2

		if res := cmp(vals[m], target); res == 0 {
			return m, true // found
		} else if res < 0 {
			l = m + 1 // search in right half
		} else {
			r = m - 1 // search in left half
		}
	}
	return
}
