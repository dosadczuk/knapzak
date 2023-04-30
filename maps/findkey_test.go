package maps_test

import (
	"testing"

	"github.com/dosadczuk/knapzak/maps"
	"github.com/google/go-cmp/cmp"
)

func TestFindKey(t *testing.T) {
	t.Run("empty map", func(t *testing.T) {
		var vals map[int]int

		var wantValue int  // zero value
		var wantFound bool // zero value
		haveValue, haveFound := maps.FindKey(vals, func(_ int, _ int) bool { return false })

		if !cmp.Equal(wantValue, haveValue) {
			t.Error(cmp.Diff(wantValue, haveValue))
		}
		if !cmp.Equal(wantFound, haveFound) {
			t.Error(cmp.Diff(wantFound, haveFound))
		}
	})
	t.Run("map with value matching predicate", func(t *testing.T) {
		vals := map[int]int{1: 2}

		wantValue, wantFound := 1, true
		haveValue, haveFound := maps.FindKey(vals, func(_ int, val int) bool {
			return val%2 == 0
		})

		if !cmp.Equal(wantValue, haveValue) {
			t.Error(cmp.Diff(wantValue, haveValue))
		}
		if !cmp.Equal(wantFound, haveFound) {
			t.Error(cmp.Diff(wantFound, haveFound))
		}
	})
	t.Run("map with value not matching predicate", func(t *testing.T) {
		vals := map[int]int{1: 2, 2: 1}

		wantValue, wantFound := 0, false
		haveValue, haveFound := maps.FindKey(vals, func(key int, val int) bool {
			return key > 2 || val > 2
		})

		if !cmp.Equal(wantValue, haveValue) {
			t.Error(cmp.Diff(wantValue, haveValue))
		}
		if !cmp.Equal(wantFound, haveFound) {
			t.Error(cmp.Diff(wantFound, haveFound))
		}
	})
}
