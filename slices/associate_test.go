package slices_test

import (
	"fmt"
	"net/netip"
	"testing"
	"time"

	. "github.com/dosadczuk/knapzak/internal/testing"
	"github.com/dosadczuk/knapzak/slices"
)

func TestAssociate(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var vals []int

		var want map[int]int
		have := slices.Associate(vals, func(val int) (int, int) { return val, 0 })

		AssertEqual(t, want, have)
	})
	t.Run("slice of primitives", func(t *testing.T) {
		vals := []int{1, 2, 3, 4, 5}

		want := map[int]bool{
			1: false,
			2: true,
			3: false,
			4: true,
			5: false,
		}
		have := slices.Associate(vals, func(val int) (int, bool) { return val, val%2 == 0 })

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

		want := map[float64]time.Time{
			1: vals[0],
			2: vals[1],
			3: vals[2],
			4: vals[3],
			5: vals[4],
		}
		have := slices.Associate(vals, func(val time.Time) (float64, time.Time) {
			return val.Sub(time.Now()).Hours(), val
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

		want := map[int]string{
			331: "1.0.0.0",
			332: "2.0.0.0",
			333: "3.0.0.0",
			334: "4.0.0.0",
			335: "5.0.0.0",
		}
		have := slices.Associate(vals, func(val fmt.Stringer) (int, string) {
			var key int
			for _, char := range val.String() {
				key += int(char)
			}
			return key, val.String()
		})

		AssertEqual(t, want, have)
	})
}
