package slices_test

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/dosadczuk/knapzak/slices"
	"github.com/google/go-cmp/cmp"
)

func TestReverse(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var want []int
		var have []int
		slices.Reverse(have)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slice of primitives", func(t *testing.T) {
		want := []int{5, 4, 3, 2, 1}
		have := []int{1, 2, 3, 4, 5}
		slices.Reverse(have)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slice of structures", func(t *testing.T) {
		want := []url.URL{
			{Scheme: "https", Host: "google.com"},
			{Scheme: "http", Host: "localhost"},
		}
		have := []url.URL{
			{Scheme: "http", Host: "localhost"},
			{Scheme: "https", Host: "google.com"},
		}
		slices.Reverse(have)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slice of interfaces", func(t *testing.T) {
		want := []fmt.Stringer{
			&url.URL{Scheme: "https", Host: "google.com"},
			&url.URL{Scheme: "http", Host: "localhost"},
		}
		have := []fmt.Stringer{
			&url.URL{Scheme: "http", Host: "localhost"},
			&url.URL{Scheme: "https", Host: "google.com"},
		}
		slices.Reverse(have)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
}
