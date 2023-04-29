package slices_test

import (
	"fmt"
	"net/url"
	"strings"
	"testing"
	"time"

	. "github.com/dosadczuk/knapzak/internal/testing"
	"github.com/dosadczuk/knapzak/slices"
)

func TestFilterNot(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var vals []int

		var want []int
		have := slices.FilterNot(vals, func(val int) bool { return false })

		AssertEqual(t, want, have)
	})
	t.Run("slice of primitives", func(t *testing.T) {
		vals := []int{1, 2, 3, 4, 5}

		want := []int{1, 2, 3}
		have := slices.FilterNot(vals, func(val int) bool { return val > 3 })

		AssertEqual(t, want, have)
	})
	t.Run("slice of structures", func(t *testing.T) {
		vals := []time.Time{
			time.Now().Add(1 * time.Hour),
			time.Now().Add(2 * time.Hour),
			time.Now().Add(3 * time.Hour),
			time.Now().Add(4 * time.Hour),
			time.Now().Add(5 * time.Hour),
		}

		want := []time.Time{
			vals[0],
			vals[1],
			vals[2],
			vals[3],
			vals[4],
		}
		have := slices.FilterNot(vals, func(val time.Time) bool {
			return val.Before(time.Now())
		})

		AssertEqual(t, want, have)
	})
	t.Run("slice of interfaces", func(t *testing.T) {
		vals := []fmt.Stringer{
			&url.URL{Scheme: "http", Host: "localhost"},
			&url.URL{Scheme: "https", Host: "localhost"},
		}

		want := []fmt.Stringer{vals[0]}
		have := slices.FilterNot(vals, func(val fmt.Stringer) bool {
			return strings.HasPrefix(val.String(), "https")
		})

		AssertEqual(t, want, have)
	})
}
