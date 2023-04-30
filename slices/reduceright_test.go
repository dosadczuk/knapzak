package slices_test

import (
	"net/url"
	"testing"

	"github.com/dosadczuk/knapzak/slices"
	"github.com/google/go-cmp/cmp"
)

func TestReduceRight(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var vals []int

		want := 0
		have := slices.ReduceRight(vals, func(acc int, val int) int { return acc + val })

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slice of primitives", func(t *testing.T) {
		vals := []int{1, 2, 3, 4, 5}

		want := 15
		have := slices.ReduceRight(vals, func(acc int, val int) int { return acc + val })

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slice of structures", func(t *testing.T) {
		vals := []url.URL{
			{Scheme: "http", Host: "localhost"},
			{Scheme: "https", Host: "google.com"},
		}

		want := url.URL{Scheme: "https", Host: "localhost"}
		have := slices.ReduceRight(
			vals,
			func(acc url.URL, val url.URL) url.URL {
				acc.Host = val.Host
				return acc
			},
		)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
}
