package slices

// Comparator compares two arguments for order.
//
// Returns:
//
//	 0 - if the arguments are equal,
//	-1 - if the first argument is less than the second,
//	+1 - if the second argument is less than the first.
type Comparator[T any] func(a, b T) int
