package slices

import (
	"golang.org/x/exp/constraints"
)

// Min returns the smallest value from the elements.
func Min[S ~[]E, E constraints.Ordered](vals S) (min E) {
	if len(vals) == 0 {
		return
	}

	min = vals[0]
	for i := 1; i < len(vals); i++ {
		if val := vals[i]; val < min {
			min = val
		}
	}
	return
}

// MinFunc returns the smallest value from the elements, comparing
// with the given Comparator function.
func MinFunc[S ~[]E, E any](vals S, compare Comparator[E]) (min E) {
	if len(vals) == 0 {
		return
	}

	min = vals[0]
	for i := 1; i < len(vals); i++ {
		if val := vals[i]; compare(val, min) < 0 {
			min = val
		}
	}
	return
}
