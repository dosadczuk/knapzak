package algorithms_test

import (
	"fmt"
	"net/netip"

	"github.com/dosadczuk/knapzak/algorithms"
)

func ExampleBinarySearch() {
	values := []int{1, 2, 3, 4, 5}

	result, found := algorithms.BinarySearch(values, 2)
	fmt.Println(result, found)

	// Output:
	// 1 true
}

func ExampleBinarySearchFunc() {
	values := []netip.Addr{
		netip.AddrFrom4([4]byte{1, 0, 0, 0}),
		netip.AddrFrom4([4]byte{2, 0, 0, 0}),
		netip.AddrFrom4([4]byte{3, 0, 0, 0}),
		netip.AddrFrom4([4]byte{4, 0, 0, 0}),
		netip.AddrFrom4([4]byte{5, 0, 0, 0}),
	}
	target := netip.AddrFrom4([4]byte{2, 0, 0, 0})

	result, found := algorithms.BinarySearchFunc(
		values, target,
		func(v1, v2 netip.Addr) int {
			return v1.Compare(v2)
		},
	)
	fmt.Println(result, found)

	// Output:
	// 1 true
}

func ExampleLevenshteinDistance() {
	src, dst := "intention", "execution"

	result := algorithms.LevenshteinDistance(src, dst)
	fmt.Println(result)

	// Output:
	// 5
}
