package maps_test

import (
	"sort"
	"testing"

	"github.com/dosadczuk/knapzak/maps"
	"github.com/google/go-cmp/cmp"
)

func TestEntries(t *testing.T) {
	t.Run("empty map", func(t *testing.T) {
		var vals map[int]int

		var want []maps.Entry[int, int]
		have := maps.Entries(vals)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("map with key-value pair", func(t *testing.T) {
		vals := map[int]int{1: 2}

		want := []maps.Entry[int, int]{{1, 2}}
		have := maps.Entries(vals)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("map with key-value pairs", func(t *testing.T) {
		vals := map[int]int{
			1: 2,
			2: 3,
			3: 4,
		}

		want := []maps.Entry[int, int]{{1, 2}, {2, 3}, {3, 4}}
		have := maps.Entries(vals)
		sort.Slice(have, func(i, j int) bool { return have[i].Key < have[j].Key })

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
}
