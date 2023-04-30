package maps_test

import (
	"fmt"
	"net/netip"
	"testing"

	"github.com/dosadczuk/knapzak/maps"
	"github.com/google/go-cmp/cmp"
)

func TestEqual(t *testing.T) {
	t.Run("empty maps", func(t *testing.T) {
		var m1 map[int]int
		var m2 map[int]int

		want := true
		have := maps.Equal(m1, m2)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("maps with different size", func(t *testing.T) {
		m1 := map[int]int{1: 2}
		m2 := map[int]int{1: 2, 2: 3}

		want := false
		have := maps.Equal(m1, m2)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("maps with the same keys but different values", func(t *testing.T) {
		m1 := map[int]int{1: 2}
		m2 := map[int]int{1: 3}

		want := false
		have := maps.Equal(m1, m2)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("maps with the same values but different keys", func(t *testing.T) {
		m1 := map[int]int{1: 2}
		m2 := map[int]int{2: 2}

		want := false
		have := maps.Equal(m1, m2)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("maps with the same key-value pairs", func(t *testing.T) {
		m1 := map[int]int{1: 2, 2: 3}
		m2 := map[int]int{1: 2, 2: 3}

		want := true
		have := maps.Equal(m1, m2)

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
}

func TestEqualFunc(t *testing.T) {
	t.Run("empty maps", func(t *testing.T) {
		var m1 map[int]fmt.Stringer
		var m2 map[int]fmt.Stringer

		want := true
		have := maps.EqualFunc(m1, m2, func(_ fmt.Stringer, _ fmt.Stringer) bool { return true })

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("maps with different size", func(t *testing.T) {
		m1 := map[int]fmt.Stringer{
			127: netip.AddrFrom4([4]byte{127, 0, 0, 1}),
		}
		m2 := map[int]fmt.Stringer{
			127: netip.AddrFrom4([4]byte{127, 0, 0, 1}),
			255: netip.AddrFrom4([4]byte{255, 0, 0, 0}),
		}

		want := false
		have := maps.EqualFunc(m1, m2, func(_ fmt.Stringer, _ fmt.Stringer) bool { return true })

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("maps with the same keys but different values", func(t *testing.T) {
		m1 := map[int]fmt.Stringer{
			127: netip.AddrFrom4([4]byte{127, 0, 1, 1}),
			255: netip.AddrFrom4([4]byte{255, 0, 1, 0}),
		}
		m2 := map[int]fmt.Stringer{
			127: netip.AddrFrom4([4]byte{127, 0, 2, 1}),
			255: netip.AddrFrom4([4]byte{255, 0, 2, 0}),
		}

		want := false
		have := maps.EqualFunc(m1, m2, func(v1 fmt.Stringer, v2 fmt.Stringer) bool {
			return v1.String() == v2.String()
		})

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("maps with the same values but different keys", func(t *testing.T) {
		m1 := map[int]fmt.Stringer{
			127: netip.AddrFrom4([4]byte{127, 0, 0, 1}),
			255: netip.AddrFrom4([4]byte{255, 0, 0, 0}),
		}
		m2 := map[int]fmt.Stringer{
			1270: netip.AddrFrom4([4]byte{127, 0, 0, 1}),
			2550: netip.AddrFrom4([4]byte{255, 0, 0, 0}),
		}

		want := false
		have := maps.EqualFunc(m1, m2, func(v1 fmt.Stringer, v2 fmt.Stringer) bool {
			return v1.String() == v2.String()
		})

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
	t.Run("maps with the same key-value pairs", func(t *testing.T) {
		m1 := map[int]fmt.Stringer{
			127: netip.AddrFrom4([4]byte{127, 0, 0, 1}),
			255: netip.AddrFrom4([4]byte{255, 0, 0, 0}),
		}
		m2 := map[int]fmt.Stringer{
			127: netip.AddrFrom4([4]byte{127, 0, 0, 1}),
			255: netip.AddrFrom4([4]byte{255, 0, 0, 0}),
		}

		want := true
		have := maps.EqualFunc(m1, m2, func(v1 fmt.Stringer, v2 fmt.Stringer) bool {
			return v1.String() == v2.String()
		})

		if !cmp.Equal(want, have) {
			t.Error(cmp.Diff(want, have))
		}
	})
}
