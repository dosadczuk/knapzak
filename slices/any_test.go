package slices_test

import (
	"fmt"
	"net/netip"
	"regexp"
	"testing"
	"time"

	. "github.com/dosadczuk/knapzak/internal/testing"
	"github.com/dosadczuk/knapzak/slices"
)

func TestAny(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var vals []int

		want := false
		have := slices.Any(vals, func(val int) bool { return false })

		AssertEqual(t, want, have)
	})
	t.Run("slice of primitives", func(t *testing.T) {
		vals := []int{1, 2, 3, 4, 5}

		want := true
		have := slices.Any(vals, func(val int) bool { return val > 3 })

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

		want := false
		have := slices.Any(vals, func(val time.Time) bool {
			return val.Before(time.Now())
		})

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

		want := true
		have := slices.Any(vals, func(val fmt.Stringer) bool {
			re := regexp.MustCompile("^\\d+\\.0.0.0$")
			return re.MatchString(val.String())
		})

		AssertEqual(t, want, have)
	})
}
