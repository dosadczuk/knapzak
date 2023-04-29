package slices_test

import (
	"fmt"
	"net/url"
	"testing"

	. "github.com/dosadczuk/knapzak/internal/testing"
	"github.com/dosadczuk/knapzak/slices"
)

func TestReversed(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var vals []int

		var want []int
		have := slices.Reversed(vals)

		AssertEqual(t, want, have)
	})
	t.Run("slice of primitives", func(t *testing.T) {
		vals := []int{1, 2, 3, 4, 5}

		want := []int{5, 4, 3, 2, 1}
		have := slices.Reversed(vals)

		AssertEqual(t, want, have)
	})
	t.Run("slice of structures", func(t *testing.T) {
		vals := []url.URL{
			{Scheme: "http", Host: "localhost"},
			{Scheme: "https", Host: "google.com"},
		}

		want := []url.URL{
			{Scheme: "https", Host: "google.com"},
			{Scheme: "http", Host: "localhost"},
		}
		have := slices.Reversed(vals)

		AssertEqual(t, want, have)
	})
	t.Run("slice of interfaces", func(t *testing.T) {
		vals := []fmt.Stringer{
			&url.URL{Scheme: "http", Host: "localhost"},
			&url.URL{Scheme: "https", Host: "google.com"},
		}

		want := []fmt.Stringer{
			&url.URL{Scheme: "https", Host: "google.com"},
			&url.URL{Scheme: "http", Host: "localhost"},
		}
		have := slices.Reversed(vals)

		AssertEqual(t, want, have)
	})
}
