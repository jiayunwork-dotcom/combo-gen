package comb

// CombinationsFromMultisets generates combinations choosing one element from each
// of several item pools. Unlike cartesian product, order within the result does
// not depend on pool order — results are sorted lexicographically.
// Each pool must be non-empty.
func CombinationsFromMultisets(pools [][]string, k int) ([][]string, error) {
	if k < 0 {
		return nil, ErrNegativeK
	}
	if k == 0 {
		return [][]string{{}}, nil
	}
	// Flatten all pools into a single universe with pool tags
	type tagged struct {
		pool int
		val  string
	}
	var universe []tagged
	for pi, pool := range pools {
		for _, v := range pool {
			universe = append(universe, tagged{pool: pi, val: v})
		}
	}
	n := len(universe)
	if k > n {
		return nil, ErrKTooLarge
	}
	// Generate all k-combinations from the universe
	var result [][]string
	cur := make([]string, 0, k)
	var walk func(start int)
	walk = func(start int) {
		if len(cur) == k {
			out := make([]string, k)
			copy(out, cur)
			result = append(result, out)
			return
		}
		for i := start; i < n; i++ {
			cur = append(cur, universe[i].val)
			walk(i + 1)
			cur = cur[:len(cur)-1]
		}
	}
	walk(0)
	return result, nil
}

// LexRank returns the lexicographic rank of a k-combination represented as
// sorted indices (0-based) within n elements. Uses combinadic representation.
func LexRank(n int, indices []int) int {
	k := len(indices)
	rank := 0
	for i, idx := range indices {
		// For the i-th position, count how many combinations start with a smaller index
		lo := 0
		if i > 0 {
			lo = indices[i-1] + 1
		}
		for j := lo; j < idx; j++ {
			rank += binomial(n-j-1, k-i-1)
		}
	}
	return rank
}

// LexUnrank returns the k-combination at the given lexicographic rank.
func LexUnrank(n, k, rank int) []int {
	if k <= 0 || rank < 0 {
		return nil
	}
	indices := make([]int, k)
	cur := 0
	for i := 0; i < k; i++ {
		lo := cur
		for j := lo; j < n; j++ {
			count := binomial(n-j-1, k-i-1)
			if rank < count {
				indices[i] = j
				cur = j + 1
				break
			}
			rank -= count
		}
	}
	return indices
}

func binomial(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	if k == 0 || k == n {
		return 1
	}
	if k > n-k {
		k = n - k
	}
	r := 1
	for i := 0; i < k; i++ {
		r = r * (n - i) / (i + 1)
	}
	return r
}
