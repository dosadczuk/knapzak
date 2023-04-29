package slices

// Reduce returns accumulated value starting with the first element
// and applying operation from left to right to current accumulator
// value.
func Reduce[S ~[]E, E any](vals S, operation func(acc E, val E) E) (out E) {
	if len(vals) == 0 {
		return
	}

	out = vals[0]
	for i := 1; i < len(vals); i++ {
		out = operation(out, vals[i])
	}
	return
}
