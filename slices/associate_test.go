package slices_test

import (
	"fmt"
	"net/netip"
	"testing"

	"github.com/dosadczuk/knapzak/slices"
	"github.com/google/go-cmp/cmp"
)

func TestAssociate(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var vals []int

		var want map[int]int
		have := slices.Associate(vals, func(val int) (int, int) { return val, 0 })

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
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

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slice of structures", func(t *testing.T) {
		vals := []netip.Addr{
			netip.AddrFrom4([4]byte{1, 0, 0, 0}),
			netip.AddrFrom4([4]byte{2, 0, 0, 0}),
		}

		want := map[string]bool{
			"1.0.0.0": true,
			"2.0.0.0": true,
		}
		have := slices.Associate(vals, func(val netip.Addr) (string, bool) {
			return val.String(), val.Is4()
		})

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
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

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
}
