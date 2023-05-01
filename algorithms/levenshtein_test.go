package algorithms_test

import (
	"testing"

	"github.com/dosadczuk/knapzak/algorithms"
	"github.com/google/go-cmp/cmp"
)

func TestLevenshteinDistance(t *testing.T) {
	tt := map[string]struct {
		// input
		src string
		dst string
		// assert
		want int
	}{
		"empty values": {
			src:  "", // zero value
			dst:  "", // zero value
			want: 0,
		},
		"src is longer than dst": {
			src:  "aaa",
			dst:  "a",
			want: 2,
		},
		"dst is longer than src": {
			src:  "a",
			dst:  "aaa",
			want: 2,
		},
		"src and dst are equal": {
			src:  "aba",
			dst:  "aba",
			want: 0,
		},
		"src and dst have the same length but different value": {
			src:  "intention",
			dst:  "execution",
			want: 5,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			have := algorithms.LevenshteinDistance(tc.src, tc.dst)

			if !cmp.Equal(tc.want, have) {
				t.Error(cmp.Diff(tc.want, have))
			}
		})
	}
}

func BenchmarkLevenshteinDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algorithms.LevenshteinDistance("intention", "execution")
	}
}
