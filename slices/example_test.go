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

func ExampleFilterNot() {
	filtered := slices.FilterNot(
		[]int{1, 2, 3, 4, 5},
		func(num int) bool {
			return num%2 == 0
		},
	)
	fmt.Printf("filtered: %v", filtered)

	// Output:
	// filtered: [1 3 5]
}

func ExampleFind() {
	value, found := slices.Find(
		[]int{1, 2, 3, 4, 5},
		func(num int) bool {
			return num%2 == 0
		},
	)
	fmt.Printf("value: %v, found: %t", value, found)

	// Output:
	// value: 2, found: true
}

func ExampleFindIndex() {
	index, found := slices.FindIndex(
		[]int{1, 2, 3, 4, 5},
		func(num int) bool {
			return num%2 == 0
		},
	)
	fmt.Printf("index: %v, found: %t", index, found)

	// Output:
	// index: 1, found: true
}

func ExampleFindLast() {
	value, found := slices.FindLast(
		[]int{1, 2, 3, 4, 5},
		func(num int) bool {
			return num%2 == 0
		},
	)
	fmt.Printf("value: %v, found: %t", value, found)

	// Output:
	// value: 4, found: true
}

func ExampleFindLastIndex() {
	index, found := slices.FindLastIndex(
		[]int{1, 2, 3, 4, 5},
		func(num int) bool {
			return num%2 == 0
		},
	)
	fmt.Printf("index: %v, found: %t", index, found)

	// Output:
	// index: 3, found: true
}

func ExampleFirst() {
	value, found := slices.First([]int{1, 2, 3, 4, 5})
	fmt.Printf("value: %v, found: %t", value, found)

	// Output:
	// value: 1, found: true
}

func ExampleIndexOf() {
	index := slices.IndexOf([]int{1, 2, 3, 4, 5}, 3)
	fmt.Printf("index: %v", index)

	// Output:
	// index: 2
}

func ExampleLast() {
	value, found := slices.Last([]int{1, 2, 3, 4, 5})
	fmt.Printf("value: %v, found: %t", value, found)

	// Output:
	// value: 5, found: true
}
