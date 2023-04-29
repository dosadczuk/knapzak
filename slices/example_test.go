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
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	notEmpty := slices.Any(values, nil)
	fmt.Printf("not empty: %t\n", notEmpty)

	matches := slices.Any(values, func(num int) bool { return num > 0 })
	fmt.Printf("matches: %t", matches)

	// Output:
	// not empty: true
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

func ExampleFlatten() {
	flattened := slices.Flatten([][]int{{1}, {1, 2}, {2, 3}})
	fmt.Printf("flattened: %v", flattened)

	// Output:
	// flattened: [1 1 2 2 3]
}

func ExampleFold() {
	order := slices.Fold(
		[]string{"1", "2", "3", "4", "5"},
		"order: ",
		func(acc string, val string) string {
			return acc + val
		},
	)
	fmt.Println(order)

	// Output:
	// order: 12345
}

func ExampleFoldRight() {
	order := slices.FoldRight(
		[]string{"1", "2", "3", "4", "5"},
		"order: ",
		func(acc string, val string) string {
			return acc + val
		},
	)
	fmt.Println(order)

	// Output:
	// order: 54321
}

func ExampleForEach() {
	slices.ForEach(
		[]int{1, 2, 3, 4, 5},
		func(idx int, num int) {
			fmt.Printf("index: %v, value: %v\n", idx, num)
		},
	)

	// Output:
	// index: 0, value: 1
	// index: 1, value: 2
	// index: 2, value: 3
	// index: 3, value: 4
	// index: 4, value: 5
}

func ExampleGroupBy() {
	grouped := slices.GroupBy(
		[]int{1, 2, 3, 4},
		func(num int) int {
			return num % 2
		},
	)
	fmt.Printf("grouped by remainder: %v", grouped)

	// Output:
	// grouped by remainder: map[0:[2 4] 1:[1 3]]
}

func ExampleIndexOf() {
	index := slices.IndexOf([]int{1, 4, 3, 4, 5}, 4)
	fmt.Printf("index: %v", index)

	// Output:
	// index: 1
}

func ExampleLast() {
	value, found := slices.Last([]int{1, 2, 3, 4, 5})
	fmt.Printf("value: %v, found: %t", value, found)

	// Output:
	// value: 5, found: true
}

func ExampleLastIndexOf() {
	index := slices.LastIndexOf([]int{1, 4, 3, 4, 5}, 4)
	fmt.Printf("index: %v", index)

	// Output:
	// index: 3
}

func ExampleMap() {
	mapped := slices.Map(
		[]int{1, 2, 3, 4, 5},
		func(num int) int {
			return 10 - num
		},
	)
	fmt.Printf("mapped: %v", mapped)

	// Output:
	// mapped: [9 8 7 6 5]
}

func ExampleMax() {
	max := slices.Max([]int{3, 2, 1, 4, 5})
	fmt.Printf("max: %v", max)

	// Output:
	// max: 5
}

func ExampleMaxFunc() {
	max := slices.MaxFunc(
		[]int{3, 2, 1, 4, 5},
		func(a, b int) int {
			if a == b {
				return 0
			}
			if a < b {
				return -1
			} else {
				return +1
			}
		},
	)
	fmt.Printf("max: %v", max)

	// Output:
	// max: 5
}

func ExampleMin() {
	min := slices.Min([]int{3, 2, 1, 4, 5})
	fmt.Printf("min: %v", min)

	// Output:
	// min: 1
}

func ExampleMinFunc() {
	min := slices.MinFunc(
		[]int{3, 2, 1, 4, 5},
		func(a, b int) int {
			if a == b {
				return 0
			}
			if a < b {
				return -1
			} else {
				return +1
			}
		},
	)
	fmt.Printf("min: %v", min)

	// Output:
	// min: 1
}

func ExampleNone() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	empty := slices.None(values, nil)
	fmt.Printf("empty: %t\n", empty)

	matches := slices.None(values, func(num int) bool { return num > 0 })
	fmt.Printf("matches: %t", matches)

	// Output:
	// empty: false
	// matches: false
}

func ExampleReduce() {
	order := slices.Reduce(
		[]string{"1", "2", "3", "4", "5"},
		func(acc string, val string) string {
			return acc + val
		},
	)
	fmt.Printf("order: %v", order)

	// Output:
	// order: 12345
}

func ExampleReduceRight() {
	order := slices.ReduceRight(
		[]string{"1", "2", "3", "4", "5"},
		func(acc string, val string) string {
			return acc + val
		},
	)
	fmt.Printf("order: %v", order)

	// Output:
	// order: 54321
}

func ExampleReverse() {
	values := []int{1, 2, 3, 4, 5}
	slices.Reverse(values)
	fmt.Printf("reversed: %v", values)

	// Output:
	// reversed: [5 4 3 2 1]
}

func ExampleReversed() {
	reversed := slices.Reversed([]int{1, 2, 3, 4, 5})
	fmt.Printf("reversed: %v", reversed)

	// Output:
	// reversed: [5 4 3 2 1]
}
