package slices_test

import (
	"fmt"
	"net/url"
	"strings"
	"testing"

	. "github.com/dosadczuk/knapzak/internal/testing"
	"github.com/dosadczuk/knapzak/slices"
)

func TestMax(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var vals []int

		want := 0
		have := slices.Max(vals)

		AssertEqual(t, want, have)
	})
	t.Run("slice of primitives", func(t *testing.T) {
		vals := []int{3, 2, 1, 4, 5}

		want := 5
		have := slices.Max(vals)

		AssertEqual(t, want, have)
	})
}

func TestMaxFunc(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var vals []int

		want := 0
		have := slices.MaxFunc(vals, func(a, b int) int { return a - b })

		AssertEqual(t, want, have)
	})
	t.Run("slice of primitives", func(t *testing.T) {
		vals := []int{3, 2, 1, 4, 5}

		want := 5
		have := slices.MaxFunc(
			vals,
			func(a, b int) int {
				if a == b {
					return 0
				}
				if a < b {
					return -1
				} else {
					return +1
				}
			},
		)

		AssertEqual(t, want, have)
	})
	t.Run("slice of structures", func(t *testing.T) {
		vals := []url.URL{
			{Scheme: "http", Host: "localhost"},
			{Scheme: "https", Host: "google.com"},
		}

		want := vals[1]
		have := slices.MaxFunc(
			vals,
			func(a, b url.URL) int {
				return strings.Compare(a.Scheme, b.Scheme)
			},
		)

		AssertEqual(t, want, have)
	})
	t.Run("slice of interfaces", func(t *testing.T) {
		vals := []fmt.Stringer{
			&url.URL{Scheme: "http", Host: "localhost"},
			&url.URL{Scheme: "https", Host: "google.com"},
		}

		want := vals[1]
		have := slices.MaxFunc(
			vals,
			func(a, b fmt.Stringer) int {
				return strings.Compare(a.String(), b.String())
			},
		)

		AssertEqual(t, want, have)
	})
}
