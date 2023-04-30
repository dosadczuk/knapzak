package maps_test

import (
	"testing"

	"github.com/dosadczuk/knapzak/maps"
	"github.com/google/go-cmp/cmp"
)

func TestKeyOf(t *testing.T) {
	t.Run("empty map", func(t *testing.T) {
		var vals map[int]int

		var want int // zero value
		have := maps.KeyOf(vals, 1)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("map with key-value pair", func(t *testing.T) {
		vals := map[int]int{1: 2}

		want := 1
		have := maps.KeyOf(vals, 2)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("map with key-value pairs", func(t *testing.T) {
		vals := map[int]int{1: 2, 2: 1}

		want := 2
		have := maps.KeyOf(vals, 1)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
}
