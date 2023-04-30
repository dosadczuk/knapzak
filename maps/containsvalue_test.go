package maps_test

import (
	"testing"

	"github.com/dosadczuk/knapzak/maps"
	"github.com/google/go-cmp/cmp"
)

func TestContainsValue(t *testing.T) {
	t.Run("empty map", func(t *testing.T) {
		var vals map[int]int

		want := false
		have := maps.ContainsValue(vals, 1)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("map with matching value", func(t *testing.T) {
		vals := map[int]int{1: 2}

		want := true
		have := maps.ContainsValue(vals, 2)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("map with not matching value", func(t *testing.T) {
		vals := map[int]int{1: 2}

		want := false
		have := maps.ContainsValue(vals, 1)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
}
