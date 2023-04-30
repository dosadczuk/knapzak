package slices_test

import (
	"fmt"
	"net/netip"
	"testing"

	"github.com/dosadczuk/knapzak/slices"
	"github.com/google/go-cmp/cmp"
)

func TestEqual(t *testing.T) {
	t.Run("empty slices", func(t *testing.T) {
		var v1, v2 []int

		want := true
		have := slices.Equal(v1, v2)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slices with different size", func(t *testing.T) {
		v1 := []int{1, 2}
		v2 := []int{3, 4, 5}

		want := false
		have := slices.Equal(v1, v2)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slices of values in the same order", func(t *testing.T) {
		v1 := []int{1, 2, 3, 4, 5}
		v2 := []int{1, 2, 3, 4, 5}

		want := true
		have := slices.Equal(v1, v2)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slices of values in different order", func(t *testing.T) {
		v1 := []int{1, 2, 3, 4, 5}
		v2 := []int{1, 5, 4, 2, 3}

		want := false
		have := slices.Equal(v1, v2)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
}

func TestEqualFunc(t *testing.T) {
	t.Run("empty slices", func(t *testing.T) {
		var v1, v2 []int

		want := true
		have := slices.EqualFunc(v1, v2, func(v1 int, v2 int) bool { return v1 == v2 })

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slices with different size", func(t *testing.T) {
		v1 := []fmt.Stringer{
			netip.AddrFrom4([4]byte{127, 0, 0, 1}),
			netip.AddrFrom4([4]byte{255, 0, 0, 0}),
		}
		v2 := []fmt.Stringer{
			netip.AddrFrom4([4]byte{127, 0, 0, 1}),
		}

		want := false
		have := slices.EqualFunc(v1, v2, func(v1 fmt.Stringer, v2 fmt.Stringer) bool {
			return v1.String() == v2.String()
		})

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slices of values in the same order", func(t *testing.T) {
		v1 := []fmt.Stringer{
			netip.AddrFrom4([4]byte{127, 0, 0, 1}),
			netip.AddrFrom4([4]byte{255, 0, 0, 0}),
		}
		v2 := []fmt.Stringer{
			netip.AddrFrom4([4]byte{127, 0, 0, 1}),
			netip.AddrFrom4([4]byte{255, 0, 0, 0}),
		}

		want := true
		have := slices.EqualFunc(v1, v2, func(v1 fmt.Stringer, v2 fmt.Stringer) bool {
			return v1.String() == v2.String()
		})

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("slices of values in different order", func(t *testing.T) {
		v1 := []fmt.Stringer{
			netip.AddrFrom4([4]byte{127, 0, 0, 1}),
			netip.AddrFrom4([4]byte{255, 0, 0, 0}),
		}
		v2 := []fmt.Stringer{
			netip.AddrFrom4([4]byte{255, 0, 0, 0}),
			netip.AddrFrom4([4]byte{127, 0, 0, 1}),
		}

		want := false
		have := slices.EqualFunc(v1, v2, func(v1 fmt.Stringer, v2 fmt.Stringer) bool {
			return v1.String() == v2.String()
		})

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
}
