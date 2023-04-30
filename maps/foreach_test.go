package maps_test

import (
	"sort"
	"testing"

	"github.com/dosadczuk/knapzak/maps"
	"github.com/google/go-cmp/cmp"
)

func TestForEach(t *testing.T) {
	t.Run("empty map", func(t *testing.T) {
		var vals map[int]int

		var want []int
		var have []int
		maps.ForEach(vals, func(_ int, val int) {
			have = append(have, val)
		})

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("map with key-value pair", func(t *testing.T) {
		vals := map[int]int{1: 2}

		want := []int{1, 2}
		var have []int
		maps.ForEach(vals, func(key int, val int) {
			have = append(have, key, val)
		})
		sort.Ints(have)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("map with key-value pairs", func(t *testing.T) {
		vals := map[int]int{1: 2, 2: 1}

		want := []int{1, 1, 2, 2}
		var have []int
		maps.ForEach(vals, func(key int, val int) {
			have = append(have, key, val)
		})
		sort.Ints(have)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
}
