package slices_test

import (
	"fmt"
	"net/url"
	"strings"
	"testing"

	"github.com/dosadczuk/knapzak/slices"
	"github.com/google/go-cmp/cmp"
)

func TestMin(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var vals []int

		want := 0
		have := slices.Min(vals)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slice with value", func(t *testing.T) {
		vals := []int{3}

		want := 3
		have := slices.Min(vals)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slice with values", func(t *testing.T) {
		vals := []int{3, 2, 1, 4, 5}

		want := 1
		have := slices.Min(vals)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
}

func TestMinFunc(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var vals []int

		want := 0
		have := slices.MinFunc(vals, func(a, b int) int { return a - b })

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slice of primitive", func(t *testing.T) {
		vals := []int{3}

		want := 3
		have := slices.MinFunc(
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

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slice of primitives", func(t *testing.T) {
		vals := []int{3, 2, 1, 4, 5}

		want := 1
		have := slices.MinFunc(
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

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slice of structures", func(t *testing.T) {
		vals := []url.URL{
			{Scheme: "http", Host: "localhost"},
			{Scheme: "https", Host: "google.com"},
		}

		want := vals[0]
		have := slices.MinFunc(vals, func(a, b url.URL) int {
			return strings.Compare(a.Scheme, b.Scheme)
		})

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slice of interfaces", func(t *testing.T) {
		vals := []fmt.Stringer{
			&url.URL{Scheme: "http", Host: "localhost"},
			&url.URL{Scheme: "https", Host: "google.com"},
		}

		want := vals[0]
		have := slices.MinFunc(vals, func(a, b fmt.Stringer) int {
			return strings.Compare(a.String(), b.String())
		})

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
}
