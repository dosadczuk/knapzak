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
