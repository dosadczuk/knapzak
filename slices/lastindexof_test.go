package slices_test

import (
	"fmt"
	"net/netip"
	"testing"
	"time"

	. "github.com/dosadczuk/knapzak/internal/testing"
	"github.com/dosadczuk/knapzak/slices"
)

func TestLastIndexOf(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var vals []int

		want := -1
		have := slices.LastIndexOf(vals, 0)

		AssertEqual(t, want, have)
	})
	t.Run("slice of primitives", func(t *testing.T) {
		vals := []int{1, 4, 3, 4, 5}

		want := 3
		have := slices.LastIndexOf(vals, 4)

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

		want := -1
		have := slices.LastIndexOf(vals, time.Now())

		AssertEqual(t, want, have)
	})
	t.Run("slice of interfaces", func(t *testing.T) {
		vals := []fmt.Stringer{
			netip.AddrFrom4([4]byte{1, 0, 0, 0}),
			netip.AddrFrom4([4]byte{2, 0, 0, 0}),
			netip.AddrFrom4([4]byte{3, 0, 0, 0}),
			netip.AddrFrom4([4]byte{4, 0, 0, 0}),
			netip.AddrFrom4([4]byte{5, 0, 0, 0}),
		}

		want := 3
		have := slices.LastIndexOf(vals, vals[3])

		AssertEqual(t, want, have)
	})
}
