package slices_test

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	. "github.com/dosadczuk/knapzak/internal/testing"
	"github.com/dosadczuk/knapzak/slices"
)

func TestLast(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var vals []int

		wantValue, wantFound := 0, false
		haveValue, haveFound := slices.Last(vals)

		AssertEqual(t, wantValue, haveValue)
		AssertEqual(t, wantFound, haveFound)
	})
	t.Run("slice of primitives", func(t *testing.T) {
		vals := []int{1, 2, 3, 4, 5}

		wantValue, wantFound := 5, true
		haveValue, haveFound := slices.Last(vals)

		AssertEqual(t, wantValue, haveValue)
		AssertEqual(t, wantFound, haveFound)
	})
	t.Run("slice of structures", func(t *testing.T) {
		vals := []time.Time{
			time.Now().Add(1 * time.Hour),
			time.Now().Add(2 * time.Hour),
			time.Now().Add(3 * time.Hour),
			time.Now().Add(4 * time.Hour),
			time.Now().Add(5 * time.Hour),
		}

		wantValue, wantFound := vals[4], true
		haveValue, haveFound := slices.Last(vals)

		AssertEqual(t, wantValue, haveValue)
		AssertEqual(t, wantFound, haveFound)
	})
	t.Run("slice of interfaces", func(t *testing.T) {
		vals := []fmt.Stringer{
			&url.URL{Scheme: "http", Host: "localhost"},
			&url.URL{Scheme: "https", Host: "localhost"},
		}

		wantValue, wantFound := vals[1], true
		haveValue, haveFound := slices.Last(vals)

		AssertEqual(t, wantValue, haveValue)
		AssertEqual(t, wantFound, haveFound)
	})
}
