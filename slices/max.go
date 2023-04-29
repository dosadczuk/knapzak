package slices

import (
	"golang.org/x/exp/constraints"
)

// Max returns the largest value from the elements.
func Max[S ~[]E, E constraints.Ordered](vals S) (max E) {
	if len(vals) == 0 {
		return
	}

	max = vals[0]
	for i := 1; i < len(vals); i++ {
		if val := vals[i]; val > max {
			max = val
		}
	}
	return
}

// MaxFunc returns the largest value from the elements, comparing
// with the given Comparator function.
func MaxFunc[S ~[]E, E any](vals S, compare Comparator[E]) (max E) {
	if len(vals) == 0 {
		return
	}

	max = vals[0]
	for i := 1; i < len(vals); i++ {
		if val := vals[i]; compare(val, max) > 0 {
			max = val
		}
	}
	return
}
