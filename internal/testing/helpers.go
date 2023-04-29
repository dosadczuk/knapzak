package testing

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// AssertEqual checks if `want` is equal to `have`.
func AssertEqual(t testing.TB, want, have any) {
	t.Helper()

	if !cmp.Equal(want, have) {
		t.Error(cmp.Diff(want, have))
	}
}
