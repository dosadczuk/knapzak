package maps_test

import (
	"sort"
	"testing"

	"github.com/dosadczuk/knapzak/maps"
	"github.com/google/go-cmp/cmp"
)

func TestMap(t *testing.T) {
	t.Run("empty map", func(t *testing.T) {
		var vals map[int]int

		var want []int
		have := maps.Map(vals, func(_ int, _ int) int { return 0 })

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("map with key-value pair", func(t *testing.T) {
		vals := map[int]int{1: 2}

		want := []int{3}
		have := maps.Map(vals, func(key int, val int) int {
			return key + val
		})

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("map with key-value pairs", func(t *testing.T) {
		vals := map[int]int{
			1: 2,
			2: 4,
			3: 6,
		}

		want := []int{3, 6, 9}
		have := maps.Map(vals, func(key int, val int) int {
			return key + val
		})
		sort.Ints(have)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
}
