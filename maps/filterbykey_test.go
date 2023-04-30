package maps_test

import (
	"testing"

	"github.com/dosadczuk/knapzak/maps"
	"github.com/google/go-cmp/cmp"
)

func TestFilterByKey(t *testing.T) {
	t.Run("empty map", func(t *testing.T) {
		var vals map[int]int

		var want map[int]int
		have := maps.FilterByKey(vals, func(key int) bool { return false })

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("map with key-value pair", func(t *testing.T) {
		vals := map[int]int{1: 2}

		want := map[int]int{1: 2}
		have := maps.FilterByKey(vals, func(key int) bool {
			return key%2 == 1
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

		want := map[int]int{1: 2, 3: 6}
		have := maps.FilterByKey(vals, func(key int) bool {
			return key%2 != 0
		})

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
}
