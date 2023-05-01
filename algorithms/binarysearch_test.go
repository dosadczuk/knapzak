package algorithms_test

import (
	"net/netip"
	"testing"

	"github.com/dosadczuk/knapzak/algorithms"
	"github.com/google/go-cmp/cmp"
)

func TestBinarySearch(t *testing.T) {
	tt := map[string]struct {
		// input
		values []int
		target int
		// assert
		wantIndex int
		wantFound bool
	}{
		"empty slice": {
			values:    nil, // zero value
			target:    4,
			wantIndex: 0,     // zero value
			wantFound: false, // zero value
		},
		"slice with matching target": {
			values:    []int{1, 2, 3, 4, 5},
			target:    2,
			wantIndex: 1,
			wantFound: true,
		},
		"slice with not matching target": {
			values:    []int{1, 2, 3, 4, 5},
			target:    10,
			wantIndex: 0,
			wantFound: false,
		},
		"slice with not sorted values and matching target": {
			values:    []int{3, 5, 2, 4, 1},
			target:    1,
			wantIndex: 0,
			wantFound: false,
		},
		"slice with not sorted values and not matching target": {
			values:    []int{3, 5, 2, 4, 1},
			target:    10,
			wantIndex: 0,
			wantFound: false,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			haveIndex, haveFound := algorithms.BinarySearch(tc.values, tc.target)

			if !cmp.Equal(tc.wantIndex, haveIndex) {
				t.Error(cmp.Diff(tc.wantIndex, haveIndex))
			}
			if !cmp.Equal(tc.wantFound, haveFound) {
				t.Error(cmp.Diff(tc.wantFound, haveFound))
			}
		})
	}
}

func TestBinarySearchFunc(t *testing.T) {
	tt := map[string]struct {
		// input
		values []netip.Addr
		target netip.Addr
		// assert
		wantIndex int
		wantFound bool
	}{
		"empty slice": {
			values:    nil, // zero value
			target:    netip.AddrFrom4([4]byte{127, 0, 0, 1}),
			wantIndex: 0,     // zero value
			wantFound: false, // zero value
		},
		"slice with matching target": {
			values: []netip.Addr{
				netip.AddrFrom4([4]byte{1, 0, 0, 0}),
				netip.AddrFrom4([4]byte{2, 0, 0, 0}),
				netip.AddrFrom4([4]byte{3, 0, 0, 0}),
				netip.AddrFrom4([4]byte{4, 0, 0, 0}),
				netip.AddrFrom4([4]byte{5, 0, 0, 0}),
			},
			target:    netip.AddrFrom4([4]byte{2, 0, 0, 0}),
			wantIndex: 1,
			wantFound: true,
		},
		"slice with not matching target": {
			values: []netip.Addr{
				netip.AddrFrom4([4]byte{1, 0, 0, 0}),
				netip.AddrFrom4([4]byte{2, 0, 0, 0}),
				netip.AddrFrom4([4]byte{3, 0, 0, 0}),
				netip.AddrFrom4([4]byte{4, 0, 0, 0}),
				netip.AddrFrom4([4]byte{5, 0, 0, 0}),
			},
			target:    netip.AddrFrom4([4]byte{127, 0, 0, 1}),
			wantIndex: 0,
			wantFound: false,
		},
		"slice with not sorted values and matching target": {
			values: []netip.Addr{
				netip.AddrFrom4([4]byte{3, 0, 0, 0}),
				netip.AddrFrom4([4]byte{5, 0, 0, 0}),
				netip.AddrFrom4([4]byte{2, 0, 0, 0}),
				netip.AddrFrom4([4]byte{4, 0, 0, 0}),
				netip.AddrFrom4([4]byte{1, 0, 0, 0}),
			},
			target:    netip.AddrFrom4([4]byte{1, 0, 0, 0}),
			wantIndex: 0,
			wantFound: false,
		},
		"slice with not sorted values and not matching target": {
			values: []netip.Addr{
				netip.AddrFrom4([4]byte{3, 0, 0, 0}),
				netip.AddrFrom4([4]byte{5, 0, 0, 0}),
				netip.AddrFrom4([4]byte{2, 0, 0, 0}),
				netip.AddrFrom4([4]byte{4, 0, 0, 0}),
				netip.AddrFrom4([4]byte{1, 0, 0, 0}),
			},
			target:    netip.AddrFrom4([4]byte{127, 0, 0, 1}),
			wantIndex: 0,
			wantFound: false,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			haveIndex, haveFound := algorithms.BinarySearchFunc(
				tc.values, tc.target,
				func(v1, v2 netip.Addr) int {
					return v1.Compare(v2)
				},
			)

			if !cmp.Equal(tc.wantIndex, haveIndex) {
				t.Error(cmp.Diff(tc.wantIndex, haveIndex))
			}
			if !cmp.Equal(tc.wantFound, haveFound) {
				t.Error(cmp.Diff(tc.wantFound, haveFound))
			}
		})
	}
}
