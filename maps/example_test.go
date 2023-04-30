package maps_test

import (
	"fmt"

	"github.com/dosadczuk/knapzak/maps"
)

func ExampleAll() {
	values := map[int]int{
		1: 2,
		2: 4,
		3: 6,
	}
	result := maps.All(values, func(key int, val int) bool {
		return 2*key == val
	})

	fmt.Printf("match: %t", result)

	// Output:
	// match: true
}

func ExampleAny() {
	values := map[int]int{
		1: 2,
		2: 4,
		3: 6,
	}
	result := maps.Any(values, func(key int, val int) bool {
		return val%2 == 0
	})

	fmt.Printf("match: %t", result)

	// Output:
	// match: true
}

func ExampleContainsKey() {
	values := map[int]int{
		1: 2,
		2: 4,
		3: 6,
	}
	result := maps.ContainsKey(values, 2)

	fmt.Printf("contains: %t", result)

	// Output:
	// contains: true
}

func ExampleNone() {
	values := map[int]int{
		1: 2,
		2: 4,
		3: 6,
	}
	result := maps.None(values, func(key int, val int) bool {
		return val%2 == 1
	})

	fmt.Printf("match: %t", result)

	// Output:
	// match: true
}
