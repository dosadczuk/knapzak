package slices_test

import (
	"fmt"

	"github.com/dosadczuk/knapzak/slices"
)

func ExampleAll() {
	matches := slices.All(
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		func(num int) bool {
			return num > 5
		},
	)
	fmt.Printf("matches: %t", matches)

	// Output:
	// matches: false
}

func ExampleAny() {
	matches := slices.Any(
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		func(num int) bool {
			return num > 0
		},
	)
	fmt.Printf("matches: %t", matches)

	// Output:
	// matches: true
}

func ExampleAssociate() {
	isNumEven := slices.Associate(
		[]int{1, 2, 3, 4, 5},
		func(num int) (int, bool) {
			return num, num%2 == 0
		},
	)
	fmt.Printf("number is even: %v", isNumEven)

	// Output:
	// number is even: map[1:false 2:true 3:false 4:true 5:false]
}

func ExampleContains() {
	contains := slices.Contains([]int{1, 2, 3, 4, 5}, 2)
	fmt.Printf("contains: %t", contains)

	// Output:
	// contains: true
}

func ExampleFilter() {
	filtered := slices.Filter(
		[]int{1, 2, 3, 4, 5},
		func(num int) bool {
			return num%2 == 0
		},
	)
	fmt.Printf("filtered: %v", filtered)

	// Output:
	// filtered: [2 4]
}

func ExampleIndexOf() {
	index := slices.IndexOf([]int{1, 2, 3, 4, 5}, 3)
	fmt.Printf("index: %v", index)

	// Output:
	// index: 2
}
