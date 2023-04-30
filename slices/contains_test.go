package slices_test

import (
	"testing"

	"github.com/dosadczuk/knapzak/slices"
	"github.com/google/go-cmp/cmp"
)

func TestContains(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var vals []int

		want := false
		have := slices.Contains(vals, 0)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slice with matching value", func(t *testing.T) {
		vals := []int{1, 2, 3, 4, 5}

		want := true
		have := slices.Contains(vals, 4)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slice with not matching value", func(t *testing.T) {
		vals := []int{1, 2, 3, 4, 5}

		want := false
		have := slices.Contains(vals, 0)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
}
