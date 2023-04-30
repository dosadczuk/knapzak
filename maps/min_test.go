package maps_test

import (
	"fmt"
	"net/url"
	"strings"
	"testing"

	"github.com/dosadczuk/knapzak/maps"
	"github.com/google/go-cmp/cmp"
)

func TestMin(t *testing.T) {
	t.Run("empty map", func(t *testing.T) {
		var vals map[int]int

		var want int // zero value
		have := maps.Min(vals)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("map with key-value pair", func(t *testing.T) {
		vals := map[int]int{1: 2}

		want := 2
		have := maps.Min(vals)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("map with key-value pairs", func(t *testing.T) {
		vals := map[int]int{
			1: 1,
			3: 3,
			2: 4,
		}

		want := 1
		have := maps.Min(vals)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
}

func TestMinFunc(t *testing.T) {
	t.Run("empty map", func(t *testing.T) {
		var vals map[int]int

		var want int // zero value
		have := maps.MinFunc(vals, func(v1, v2 int) int { return v1 - v2 })

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("map with key-value pair", func(t *testing.T) {
		vals := map[int]fmt.Stringer{
			127: &url.URL{Host: "127.0.0.1"},
		}

		want := vals[127]
		have := maps.MinFunc(vals, func(v1, v2 fmt.Stringer) int {
			return strings.Compare(v1.String(), v2.String())
		})

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("map with key-value pairs", func(t *testing.T) {
		vals := map[int]fmt.Stringer{
			255: &url.URL{Host: "255.0.0.0"},
			127: &url.URL{Host: "127.0.0.1"},
		}

		want := vals[127]
		have := maps.MinFunc(vals, func(v1, v2 fmt.Stringer) int {
			return strings.Compare(v1.String(), v2.String())
		})

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
}
