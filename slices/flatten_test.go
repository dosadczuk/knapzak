package slices_test

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/dosadczuk/knapzak/slices"
	"github.com/google/go-cmp/cmp"
)

func TestFlatten(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var vals [][]int

		var want []int
		have := slices.Flatten(vals)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slice of primitives", func(t *testing.T) {
		vals := [][]int{{1}, {1, 2}, {2, 3}}

		want := []int{1, 1, 2, 2, 3}
		have := slices.Flatten(vals)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slice of structures", func(t *testing.T) {
		vals := [][]url.URL{
			{{Host: "google.com"}, {Host: "facebook.com"}},
			{{Host: "localhost"}},
			{{Host: "apple.com"}},
		}

		want := []url.URL{
			{Host: "google.com"},
			{Host: "facebook.com"},
			{Host: "localhost"},
			{Host: "apple.com"},
		}
		have := slices.Flatten(vals)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slice of interfaces", func(t *testing.T) {
		vals := [][]fmt.Stringer{
			{&url.URL{Host: "google.com"}},
			{&url.URL{Host: "localhost"}},
		}
		want := []fmt.Stringer{
			&url.URL{Host: "google.com"},
			&url.URL{Host: "localhost"},
		}
		have := slices.Flatten(vals)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
}
