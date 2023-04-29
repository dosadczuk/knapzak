package slices_test

import (
	"fmt"

	"github.com/dosadczuk/knapzak/slices"
)

// -----------------------------------------------------------------------------
// -- Constants
// -----------------------------------------------------------------------------

var vals = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// -----------------------------------------------------------------------------
// -- Examples
// -----------------------------------------------------------------------------

func ExampleAll() {
	matches := slices.All(vals, func(num int) bool { return num > 5 })
	fmt.Printf("matches: %t", matches)

	// Output:
	// matches: false
}

func ExampleAny() {
	matches := slices.Any(vals, func(num int) bool { return num > 0 })
	fmt.Printf("matches: %t", matches)

	// Output:
	// matches: true
}
