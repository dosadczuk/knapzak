package slices_test

import (
	"testing"

	"github.com/dosadczuk/knapzak/slices"
	"github.com/google/go-cmp/cmp"
)

func TestFind(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var vals []int

		wantValue, wantFound := 0, false
		haveValue, haveFound := slices.Find(vals, func(val int) bool { return false })

		if !cmp.Equal(wantValue, haveValue) {
			t.Error(cmp.Diff(wantValue, haveValue))
		}
		if !cmp.Equal(wantFound, haveFound) {
			t.Error(cmp.Diff(wantFound, haveFound))
		}
	})
	t.Run("slice with value matching predicate", func(t *testing.T) {
		vals := []int{1, 2, 3, 4, 5}

		wantValue, wantFound := 2, true
		haveValue, haveFound := slices.Find(vals, func(val int) bool {
			return val%2 == 0
		})

		if !cmp.Equal(wantValue, haveValue) {
			t.Error(cmp.Diff(wantValue, haveValue))
		}
		if !cmp.Equal(wantFound, haveFound) {
			t.Error(cmp.Diff(wantFound, haveFound))
		}
	})
	t.Run("slice with value not matching predicate", func(t *testing.T) {
		vals := []int{1, 2, 3, 4, 5}

		wantValue, wantFound := 0, false
		haveValue, haveFound := slices.Find(vals, func(val int) bool {
			return val < 0
		})

		if !cmp.Equal(wantValue, haveValue) {
			t.Error(cmp.Diff(wantValue, haveValue))
		}
		if !cmp.Equal(wantFound, haveFound) {
			t.Error(cmp.Diff(wantFound, haveFound))
		}
	})
}
