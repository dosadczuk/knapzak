package slices_test

import (
	"net/url"
	"testing"

	. "github.com/dosadczuk/knapzak/internal/testing"
	"github.com/dosadczuk/knapzak/slices"
)

func TestFold(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var vals []int

		want := 10
		have := slices.Fold(vals, 10, func(acc int, val int) int { return acc + val })

		AssertEqual(t, want, have)
	})
	t.Run("slice of primitives", func(t *testing.T) {
		vals := []int{1, 2, 3, 4, 5}

		want := 35
		have := slices.Fold(vals, 20, func(acc int, val int) int { return acc + val })

		AssertEqual(t, want, have)
	})
	t.Run("slice of structures", func(t *testing.T) {
		vals := []url.URL{
			{Scheme: "http", Host: "localhost"},
			{Scheme: "https", Host: "google.com"},
		}

		want := url.URL{Host: "google.com"}
		have := slices.Fold(vals, url.URL{}, func(acc url.URL, val url.URL) url.URL {
			acc.Host = val.Host
			return acc
		})

		AssertEqual(t, want, have)
	})
}
