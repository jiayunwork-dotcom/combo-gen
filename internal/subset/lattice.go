package subset

func ChainMaximal(items []string) [][]int {
	n := len(items)
	if n == 0 {
		return [][]int{{0}}
	}
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

func MaxAntiChainSize(n int) int {
	if n < 0 {
		return 0
	}
	return binomialInt(n, n/2)
}

func KSubsetLatticeLevel(items []string, k int) [][]string {
	return SubsetsOfSize(items, k)
}

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
