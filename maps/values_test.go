package maps_test

import (
	"sort"
	"testing"

	"github.com/dosadczuk/knapzak/maps"
	"github.com/google/go-cmp/cmp"
)

func TestValues(t *testing.T) {
	t.Run("empty map", func(t *testing.T) {
		var vals map[int]int

		var want []int
		have := maps.Values(vals)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("map with key-value pair", func(t *testing.T) {
		vals := map[int]int{1: 2}

		want := []int{2}
		have := maps.Values(vals)
		sort.Ints(have)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("map with key-value pairs", func(t *testing.T) {
		vals := map[int]int{1: 2, 2: 1}

		want := []int{1, 2}
		have := maps.Values(vals)
		sort.Ints(have)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
}
