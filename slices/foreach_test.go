package slices_test

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	. "github.com/dosadczuk/knapzak/internal/testing"
	"github.com/dosadczuk/knapzak/slices"
)

func TestForEach(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var want []int
		var have []int
		slices.ForEach(want, func(_, val int) {
			have = append(have, val)
		})

		AssertEqual(t, want, have)
	})
	t.Run("slice of primitives", func(t *testing.T) {
		want := []int{1, 2, 3, 4, 5}
		have := make([]int, 0, len(want))
		slices.ForEach(want, func(_, val int) {
			have = append(have, val)
		})

		AssertEqual(t, want, have)
	})
	t.Run("slice of structures", func(t *testing.T) {
		want := []time.Time{
			time.Now().Add(1 * time.Hour),
			time.Now().Add(2 * time.Hour),
			time.Now().Add(3 * time.Hour),
			time.Now().Add(4 * time.Hour),
			time.Now().Add(5 * time.Hour),
		}
		have := make([]time.Time, 0, len(want))
		slices.ForEach(want, func(_ int, val time.Time) {
			have = append(have, val)
		})

		AssertEqual(t, want, have)
	})
	t.Run("slice of interfaces", func(t *testing.T) {
		want := []fmt.Stringer{
			&url.URL{Scheme: "http", Host: "localhost"},
			&url.URL{Scheme: "https", Host: "localhost"},
		}
		have := make([]fmt.Stringer, 0, len(want))
		slices.ForEach(want, func(_ int, val fmt.Stringer) {
			have = append(have, val)
		})

		AssertEqual(t, want, have)
	})
}
