package algorithms

import (
	"math"
)

// LevenshteinDistance between two words is the minimum number of
// single-character edits (insertions, deletions or substitutions)
// required to change one word into the other.
//
// Source: https://en.wikipedia.org/wiki/Levenshtein_distance
func LevenshteinDistance(src, dst string) int {
	n, m := len(src), len(dst)

	distance := make([][]int, n+1)
	for i := range distance {
		distance[i] = make([]int, m+1)
	}

	// `src` prefixes can be transformed into empty
	// string by dropping all characters.
	for i := 1; i <= n; i++ {
		distance[i][0] = i
	}
	// `dst` prefixes can be transformed into `src`
	// prefixes by inserting characters.
	for j := 1; j <= m; j++ {
		distance[0][j] = j
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if src[i-1] == dst[j-1] {
				distance[i][j] = distance[i-1][j-1] // skip character
			} else {
				costs := []int{
					distance[i-1][j],   // delete character
					distance[i][j-1],   // insert character
					distance[i-1][j-1], // replace characters
				}

				minCost := math.MaxInt
				for _, cost := range costs {
					if cost < minCost {
						minCost = cost
					}
				}

				distance[i][j] = 1 + minCost
			}
		}
	}

	return distance[n][m]
}
