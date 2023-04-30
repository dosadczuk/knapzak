package slices_test

import (
	"fmt"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/dosadczuk/knapzak/slices"
	"github.com/google/go-cmp/cmp"
)

func TestFilter(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var vals []int

		var want []int
		have := slices.Filter(vals, func(val int) bool { return false })

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slice of primitives", func(t *testing.T) {
		vals := []int{1, 2, 3, 4, 5}

		want := []int{4, 5}
		have := slices.Filter(vals, func(val int) bool { return val > 3 })

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slice of structures", func(t *testing.T) {
		vals := []time.Time{
			time.Now().Add(1 * time.Hour),
			time.Now().Add(2 * time.Hour),
			time.Now().Add(3 * time.Hour),
			time.Now().Add(4 * time.Hour),
			time.Now().Add(5 * time.Hour),
		}

		var want []time.Time
		have := slices.Filter(vals, func(val time.Time) bool {
			return val.Before(time.Now())
		})

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slice of interfaces", func(t *testing.T) {
		vals := []fmt.Stringer{
			&url.URL{Scheme: "http", Host: "localhost"},
			&url.URL{Scheme: "https", Host: "localhost"},
		}

		want := []fmt.Stringer{vals[1]}
		have := slices.Filter(vals, func(val fmt.Stringer) bool {
			return strings.HasPrefix(val.String(), "https")
		})

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
}
