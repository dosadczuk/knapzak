package slices_test

import (
	"fmt"
	"net/url"
	"strings"
	"testing"

	. "github.com/dosadczuk/knapzak/internal/testing"
	"github.com/dosadczuk/knapzak/slices"
)

func TestGroupBy(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var vals []int

		var want map[int][]int
		have := slices.GroupBy(vals, func(val int) int { return val })

		AssertEqual(t, want, have)
	})
	t.Run("slice of primitives", func(t *testing.T) {
		vals := []int{1, 2, 3, 4}

		want := map[int][]int{
			0: {2, 4},
			1: {1, 3},
		}
		have := slices.GroupBy(vals, func(val int) int { return val % 2 })

		AssertEqual(t, want, have)
	})
	t.Run("slice of structures", func(t *testing.T) {
		vals := []url.URL{
			{Scheme: "http", Host: "localhost"},
			{Scheme: "https", Host: "google.com"},
			{Scheme: "https", Host: "vercel.com"},
		}

		want := map[string][]url.URL{
			"http": {
				{Scheme: "http", Host: "localhost"},
			},
			"https": {
				{Scheme: "https", Host: "google.com"},
				{Scheme: "https", Host: "vercel.com"},
			},
		}
		have := slices.GroupBy(vals, func(val url.URL) string { return val.Scheme })

		AssertEqual(t, want, have)
	})
	t.Run("slice of interfaces", func(t *testing.T) {
		vals := []fmt.Stringer{
			&url.URL{Host: "localhost"},
			&url.URL{Host: "google.com"},
			&url.URL{Host: "vercel.com"},
		}

		want := map[byte][]fmt.Stringer{
			'l': {&url.URL{Host: "localhost"}},
			'g': {&url.URL{Host: "google.com"}},
			'v': {&url.URL{Host: "vercel.com"}},
		}
		have := slices.GroupBy(vals, func(val fmt.Stringer) byte {
			return strings.TrimPrefix(val.String(), "//")[0]
		})

		AssertEqual(t, want, have)
	})
}
