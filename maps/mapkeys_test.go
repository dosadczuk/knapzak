package maps_test

import (
	"testing"

	"github.com/dosadczuk/knapzak/maps"
	"github.com/google/go-cmp/cmp"
)

func TestMapKeys(t *testing.T) {
	t.Run("empty map", func(t *testing.T) {
		var vals map[int]int

		var want map[int]int
		have := maps.MapKeys(vals, func(_ int, _ int) int { return 0 })

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("map with key-value pair", func(t *testing.T) {
		vals := map[int]int{1: 2}

		want := map[int]int{3: 2}
		have := maps.MapKeys(vals, func(key int, val int) int {
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

		want := map[int]int{
			3: 2,
			6: 4,
			9: 6,
		}
		have := maps.MapKeys(vals, func(key int, val int) int {
			return key + val
		})

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
}
