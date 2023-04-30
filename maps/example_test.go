package maps_test

import (
	"fmt"
	"sort"

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
	result := maps.ContainsKey(values, 3)

	fmt.Printf("contains: %t", result)

	// Output:
	// contains: true
}

func ExampleContainsValue() {
	values := map[int]int{
		1: 2,
		2: 4,
		3: 6,
	}
	result := maps.ContainsValue(values, 6)

	fmt.Printf("contains: %t", result)

	// Output:
	// contains: true
}

func ExampleKeys() {
	values := map[int]int{
		1: 2,
		2: 4,
		3: 6,
	}
	result := maps.Keys(values)
	sort.Ints(result) // only for `Output` purpose

	fmt.Printf("keys: %v", result)

	// Output:
	// keys: [1 2 3]
}

func ExampleValues() {
	values := map[int]int{
		1: 2,
		2: 4,
		3: 6,
	}
	result := maps.Values(values)
	sort.Ints(result) // only for `Output` purpose

	fmt.Printf("keys: %v", result)

	// Output:
	// keys: [2 4 6]
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
