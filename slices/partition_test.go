package slices_test

import (
	"fmt"
	"net/url"
	"strings"
	"testing"

	"github.com/dosadczuk/knapzak/slices"
	"github.com/google/go-cmp/cmp"
)

func TestPartition(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var vals []int

		var wantT, wantF []int
		haveT, haveF := slices.Partition(vals, func(val int) bool { return false })

		if !cmp.Equal(wantT, haveT) {
			t.Error(cmp.Diff(wantT, haveT))
		}
		if !cmp.Equal(wantF, haveF) {
			t.Error(cmp.Diff(wantF, haveF))
		}
	})
	t.Run("slice of primitives", func(t *testing.T) {
		vals := []int{1, 2, 3, 4, 5}

		wantT, wantF := []int{4, 5}, []int{1, 2, 3}
		haveT, haveF := slices.Partition(vals, func(val int) bool { return val > 3 })

		if !cmp.Equal(wantT, haveT) {
			t.Error(cmp.Diff(wantT, haveT))
		}
		if !cmp.Equal(wantF, haveF) {
			t.Error(cmp.Diff(wantF, haveF))
		}
	})
	t.Run("slice of structures", func(t *testing.T) {
		vals := []url.URL{
			{Scheme: "http", Host: "0.0.0.0"},
			{Scheme: "https", Host: "0.0.0.0"},
			{Scheme: "http", Host: "127.0.0.1"},
			{Scheme: "https", Host: "127.0.0.1"},
		}

		wantT, wantF := []url.URL{vals[0], vals[2]}, []url.URL{vals[1], vals[3]}
		haveT, haveF := slices.Partition(vals, func(addr url.URL) bool {
			return addr.Scheme == "http"
		})

		if !cmp.Equal(wantT, haveT) {
			t.Error(cmp.Diff(wantT, haveT))
		}
		if !cmp.Equal(wantF, haveF) {
			t.Error(cmp.Diff(wantF, haveF))
		}
	})
	t.Run("slice of interfaces", func(t *testing.T) {
		vals := []fmt.Stringer{
			&url.URL{Scheme: "http", Host: "localhost"},
			&url.URL{Scheme: "https", Host: "localhost"},
		}

		wantT, wantF := []fmt.Stringer{vals[1]}, []fmt.Stringer{vals[0]}
		haveT, haveF := slices.Partition(vals, func(val fmt.Stringer) bool {
			return strings.HasPrefix(val.String(), "https")
		})

		if !cmp.Equal(wantT, haveT) {
			t.Error(cmp.Diff(wantT, haveT))
		}
		if !cmp.Equal(wantF, haveF) {
			t.Error(cmp.Diff(wantF, haveF))
		}
	})
}
