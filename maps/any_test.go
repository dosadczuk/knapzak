package maps_test

import (
	"testing"

	"github.com/dosadczuk/knapzak/maps"
	"github.com/google/go-cmp/cmp"
)

func TestAny(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var vals map[int]int

		want := false
		have := maps.Any(vals, func(_ int, _ int) bool { return false })

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("empty slice and predicate is nil", func(t *testing.T) {
		var vals map[int]int

		want := false
		have := maps.Any(vals, nil)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slice with values and predicate is nil", func(t *testing.T) {
		vals := map[int]int{
			1: 2,
			2: 4,
			3: 6,
		}

		want := true
		have := maps.Any(vals, nil)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slice with values matching predicate", func(t *testing.T) {
		vals := map[int]int{
			1: 2,
			2: 4,
			3: 6,
		}

		want := true
		have := maps.Any(vals, func(_ int, v int) bool { return v%2 == 0 })

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slice with values not matching predicate", func(t *testing.T) {
		vals := map[int]int{
			1: 2,
			2: 4,
			3: 6,
		}

		want := false
		have := maps.Any(vals, func(_ int, v int) bool { return v%2 == 1 })

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
}
