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

func TestPartition(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var vals []int

		var wantT, wantF []int
		haveT, haveF := slices.Partition(vals, func(val int) bool { return false })

		AssertEqual(t, wantT, haveT)
		AssertEqual(t, wantF, haveF)
	})
	t.Run("slice of primitives", func(t *testing.T) {
		vals := []int{1, 2, 3, 4, 5}

		wantT, wantF := []int{4, 5}, []int{1, 2, 3}
		haveT, haveF := slices.Partition(vals, func(val int) bool { return val > 3 })

		AssertEqual(t, wantT, haveT)
		AssertEqual(t, wantF, haveF)
	})
	t.Run("slice of structures", func(t *testing.T) {
		vals := []time.Time{
			time.Now().Add(1 * time.Hour),
			time.Now().Add(2 * time.Hour),
			time.Now().Add(3 * time.Hour),
			time.Now().Add(4 * time.Hour),
			time.Now().Add(5 * time.Hour),
		}

		wantT, wantF := []time.Time{vals[0], vals[1]}, []time.Time{vals[2], vals[3], vals[4]}
		haveT, haveF := slices.Partition(vals, func(val time.Time) bool {
			return val.Before(time.Now().Add(3 * time.Hour))
		})

		AssertEqual(t, wantT, haveT)
		AssertEqual(t, wantF, haveF)
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

		AssertEqual(t, wantT, haveT)
		AssertEqual(t, wantF, haveF)
	})
}
