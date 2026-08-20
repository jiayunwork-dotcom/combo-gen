package subset

// ChainMaximal returns all maximal chains in the subset lattice from {} to items.
// A chain is a sequence of subsets where each is contained in the next, differing by one element.
func ChainMaximal(items []string) [][]int {
	n := len(items)
	if n == 0 {
		return [][]int{{0}} // only the empty set
	}
	// A maximal chain has n+1 elements (from {} to full set)
	// Each chain corresponds to a permutation of indices
	perms := permIndices(n)
	chains := make([][]int, len(perms))
	for i, perm := range perms {
		chain := make([]int, n+1)
		mask := 0
		chain[0] = 0
		for j, idx := range perm {
			mask |= (1 << idx)
			chain[j+1] = mask
		}
		chains[i] = chain
	}
	return chains
}

// AntiChain checks if a collection of subsets forms an antichain
// (no subset is contained in another).
func AntiChain(subsets [][]string) bool {
	for i := 0; i < len(subsets); i++ {
		for j := 0; j < len(subsets); j++ {
			if i == j {
				continue
			}
			if isSubset(subsets[i], subsets[j]) {
				return false
			}
		}
	}
	return true
}

// MaxAntiChainSize returns the size of the largest antichain in a set of n elements.
// By Dilworth's theorem, this equals C(n, floor(n/2)).
func MaxAntiChainSize(n int) int {
	if n < 0 {
		return 0
	}
	return binomialInt(n, n/2)
}

// KSubsetLatticeLevel returns all subsets of size exactly k (one level of the lattice).
func KSubsetLatticeLevel(items []string, k int) [][]string {
	return SubsetsOfSize(items, k)
}

// Covers returns true if b covers a in the subset lattice (a ⊂ b and |b| = |a| + 1).
func Covers(a, b []string) bool {
	if len(b) != len(a)+1 {
		return false
	}
	return isSubset(a, b)
}

func isSubset(a, b []string) bool {
	setB := toSet(b)
	for _, v := range a {
		if !setB[v] {
			return false
		}
	}
	return true
}

func permIndices(n int) [][]int {
	if n == 0 {
		return [][]int{{}}
	}
	var result [][]int
	used := make([]bool, n)
	cur := make([]int, 0, n)
	var walk func()
	walk = func() {
		if len(cur) == n {
			out := make([]int, n)
			copy(out, cur)
			result = append(result, out)
			return
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			cur = append(cur, i)
			walk()
			cur = cur[:len(cur)-1]
			used[i] = false
		}
	}
	walk()
	return result
}

func binomialInt(n, k int) int {
	if k < 0 || k > n {
		return 0
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
