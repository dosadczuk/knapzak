package slices_test

import (
	"testing"

	"github.com/dosadczuk/knapzak/slices"
	"github.com/google/go-cmp/cmp"
)

func TestAny(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var vals []int

		want := false
		have := slices.Any(vals, func(val int) bool { return false })

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("empty slice and predicate is nil", func(t *testing.T) {
		var vals []int

		want := false
		have := slices.Any(vals, nil)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slice with values and predicate is nil", func(t *testing.T) {
		vals := []int{1, 2, 3, 4, 5}

		want := true
		have := slices.Any(vals, nil)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slice with values matching predicate", func(t *testing.T) {
		vals := []int{1, 2, 3, 4, 5}

		want := true
		have := slices.Any(vals, func(val int) bool { return val > 0 })

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slice with values not matching predicate", func(t *testing.T) {
		vals := []int{1, 2, 3, 4, 5}

		want := false
		have := slices.Any(vals, func(val int) bool { return val < 0 })

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
}
