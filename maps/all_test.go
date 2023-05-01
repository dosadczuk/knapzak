package maps_test

import (
	"testing"

	"github.com/dosadczuk/knapzak/maps"
	"github.com/google/go-cmp/cmp"
)

func TestAll(t *testing.T) {
	tt := map[string]struct {
		// input
		values    map[int]int
		predicate func(int, int) bool
		// assert
		want bool
	}{
		"empty map": {
			values:    nil, // zero value
			predicate: func(_ int, _ int) bool { return true },
			want:      true,
		},
		"map with values matching predicate": {
			values: map[int]int{
				1: 2,
				2: 4,
				3: 6,
			},
			predicate: func(_ int, val int) bool { return val%2 == 0 },
			want:      true,
		},
		"map with values not matching predicate": {
			values: map[int]int{
				1: 2,
				2: 4,
				3: 6,
			},
			predicate: func(_ int, val int) bool { return val%2 == 1 },
			want:      false,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			have := maps.All(tc.values, tc.predicate)

			if !cmp.Equal(tc.want, have) {
				t.Error(cmp.Diff(tc.want, have))
			}
		})
	}
}
