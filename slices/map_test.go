package slices_test

import (
	"fmt"
	"net/url"
	"strconv"
	"testing"

	. "github.com/dosadczuk/knapzak/internal/testing"
	"github.com/dosadczuk/knapzak/slices"
)

func TestMap(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		var vals []int

		var want []int
		have := slices.Map(vals, func(val int) int { return val * 2 })

		AssertEqual(t, want, have)
	})
	t.Run("slice of primitives", func(t *testing.T) {
		t.Run("int to string", func(t *testing.T) {
			vals := []int{1, 2, 3}

			want := []string{"1", "2", "3"}
			have := slices.Map(vals, func(val int) string {
				return fmt.Sprintf("%d", val)
			})

			AssertEqual(t, want, have)
		})
		t.Run("float to string", func(t *testing.T) {
			vals := []float64{3.14, 12.0}

			want := []string{"3.1400", "12.0000"}
			have := slices.Map(vals, func(val float64) string {
				return fmt.Sprintf("%.4f", val)
			})

			AssertEqual(t, want, have)
		})
		t.Run("int to rune", func(t *testing.T) {
			vals := []int{97, 98, 99, 100}

			want := []rune{'a', 'b', 'c', 'd'}
			have := slices.Map(vals, func(val int) rune { return rune(val) })

			AssertEqual(t, want, have)
		})
		t.Run("rune to int", func(t *testing.T) {
			vals := []rune{'a', 'b', 'c', 'd'}

			want := []int{97, 98, 99, 100}
			have := slices.Map(vals, func(val rune) int { return int(val) })

			AssertEqual(t, want, have)
		})
		t.Run("string to int", func(t *testing.T) {
			vals := []string{"3", "12", "err"}

			want := []int64{3, 12, 0}
			have := slices.Map(vals, func(val string) int64 {
				out, err := strconv.ParseInt(val, 10, 64)
				if err != nil {
					return 0
				}
				return out
			})

			AssertEqual(t, want, have)
		})
		t.Run("string to float", func(t *testing.T) {
			vals := []string{"3.14", "12.0", "err"}

			want := []float64{3.14, 12.00, 0.00}
			have := slices.Map(vals, func(val string) float64 {
				out, err := strconv.ParseFloat(val, 64)
				if err != nil {
					return 0.00
				}
				return out
			})

			AssertEqual(t, want, have)
		})
	})
	t.Run("slice of structures", func(t *testing.T) {
		vals := []url.URL{
			{Scheme: "http", Host: "localhost"},
			{Scheme: "https", Host: "google.com"},
		}

		want := []string{"localhost", "google.com"}
		have := slices.Map(vals, func(val url.URL) string { return val.Host })

		AssertEqual(t, want, have)
	})
	t.Run("slice of interfaces", func(t *testing.T) {
		vals := []fmt.Stringer{
			&url.URL{Scheme: "http", Host: "localhost"},
			&url.URL{Scheme: "https", Host: "google.com"},
		}

		want := []string{"http://localhost", "https://google.com"}
		have := slices.Map(vals, func(val fmt.Stringer) string { return val.String() })

		AssertEqual(t, want, have)
	})
}
