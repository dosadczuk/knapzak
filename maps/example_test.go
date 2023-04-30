package maps_test

import (
	"fmt"
	"net/netip"
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
	fmt.Println(result)

	// Output:
	// true
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
	fmt.Println(result)

	// Output:
	// true
}

func ExampleContainsKey() {
	values := map[int]int{
		1: 2,
		2: 4,
		3: 6,
	}

	result := maps.ContainsKey(values, 3)
	fmt.Println(result)

	// Output:
	// true
}

func ExampleContainsValue() {
	values := map[int]int{
		1: 2,
		2: 4,
		3: 6,
	}

	result := maps.ContainsValue(values, 6)
	fmt.Println(result)

	// Output:
	// true
}

func ExampleEqual() {
	m1 := map[int]int{
		1: 2,
		2: 4,
	}
	m2 := map[int]int{
		1: 2,
		2: 4,
	}

	result := maps.Equal(m1, m2)
	fmt.Println(result)

	// Output:
	// true
}

func ExampleEqualFunc() {
	m1 := map[int]fmt.Stringer{
		127: netip.AddrFrom4([4]byte{127, 0, 0, 1}),
		255: netip.AddrFrom4([4]byte{255, 0, 0, 0}),
	}
	m2 := map[int]fmt.Stringer{
		127: netip.AddrFrom4([4]byte{127, 0, 0, 1}),
		255: netip.AddrFrom4([4]byte{255, 0, 0, 0}),
	}

	result := maps.EqualFunc(m1, m2, func(v1 fmt.Stringer, v2 fmt.Stringer) bool {
		return v1.String() == v2.String()
	})
	fmt.Println(result)

	// Output:
	// true
}

func ExampleFilter() {
	values := map[int]int{
		1: 2,
		2: 4,
		3: 6,
	}

	result := maps.Filter(values, func(key int, _ int) bool {
		return key%2 == 1
	})
	fmt.Println(result)

	// Output:
	// map[1:2 3:6]
}

func ExampleFilterByKey() {
	values := map[int]int{
		1: 2,
		2: 4,
		3: 6,
	}

	result := maps.FilterByKey(values, func(key int) bool {
		return key%2 == 1
	})
	fmt.Println(result)

	// Output:
	// map[1:2 3:6]
}

func ExampleFilterByValue() {
	values := map[int]int{
		1: 2,
		2: 4,
		3: 6,
	}

	result := maps.FilterByValue(values, func(val int) bool {
		return val%4 == 0
	})
	fmt.Println(result)

	// Output:
	// map[2:4]
}

func ExampleFind() {
	values := map[int]int{
		1: 2,
		2: 4,
		3: 6,
	}

	result, found := maps.Find(values, func(key int, _ int) bool {
		return key%2 == 0
	})
	fmt.Println(result, found)

	// Output:
	// 4 true
}

func ExampleFindKey() {
	values := map[int]int{
		1: 2,
		2: 4,
		3: 6,
	}

	result, found := maps.FindKey(values, func(_ int, val int) bool {
		return val > 5
	})
	fmt.Println(result, found)

	// Output:
	// 3 true
}

func ExampleForEach() {
	values := map[int]int{
		1: 2,
		2: 4,
		3: 6,
	}

	var result []string
	maps.ForEach(values, func(key int, val int) {
		result = append(result, fmt.Sprintf("%v:%v", key, val))
	})
	sort.Strings(result) // only for `Output` purpose

	fmt.Println(result)

	// Output:
	// [1:2 2:4 3:6]
}

func ExampleKeyOf() {
	values := map[int]int{
		1: 2,
		2: 4,
		3: 6,
	}

	result := maps.KeyOf(values, 6)
	fmt.Println(result)

	// Output:
	// 3
}

func ExampleKeys() {
	values := map[int]int{
		1: 2,
		2: 4,
		3: 6,
	}

	result := maps.Keys(values)
	sort.Ints(result) // only for `Output` purpose
	fmt.Println(result)

	// Output:
	// [1 2 3]
}

func ExampleValues() {
	values := map[int]int{
		1: 2,
		2: 4,
		3: 6,
	}

	result := maps.Values(values)
	sort.Ints(result) // only for `Output` purpose
	fmt.Println(result)

	// Output:
	// [2 4 6]
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
	fmt.Println(result)

	// Output:
	// true
}
