// Package subset generates subsets (power set) and related enumerations.
package subset

// PowerSet returns all 2^n subsets of items, including the empty set.
// Each subset is a fresh copy. Order: by binary counting (lexicographic on indices).
func PowerSet(items []string) [][]string {
	n := len(items)
	total := 1 << n
	result := make([][]string, 0, total)
	for mask := 0; mask < total; mask++ {
		var sub []string
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				sub = append(sub, items[i])
			}
		}
		if sub == nil {
			sub = []string{}
		}
		result = append(result, sub)
	}
	return result
}

// SubsetsOfSize returns all subsets of exactly size k.
// Returns nil if k < 0 or k > len(items).
func SubsetsOfSize(items []string, k int) [][]string {
	n := len(items)
	if k < 0 || k > n {
		return nil
	}
	if k == 0 {
		return [][]string{{}}
	}
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
			cur = append(cur, items[i])
			walk(i + 1)
			cur = cur[:len(cur)-1]
		}
	}
	walk(0)
	return result
}

// SubsetsOfSizeRange returns all subsets with size in [minK, maxK].
func SubsetsOfSizeRange(items []string, minK, maxK int) [][]string {
	if minK < 0 {
		minK = 0
	}
	n := len(items)
	if maxK > n {
		maxK = n
	}
	var result [][]string
	for k := minK; k <= maxK; k++ {
		result = append(result, SubsetsOfSize(items, k)...)
	}
	return result
}

// NonEmptySubsets returns all non-empty subsets (power set minus empty set).
func NonEmptySubsets(items []string) [][]string {
	all := PowerSet(items)
	if len(all) == 0 {
		return nil
	}
	// First element is always the empty set
	return all[1:]
}

// Complement returns items not in the given subset. Items order is preserved.
func Complement(items []string, subset []string) []string {
	inSub := make(map[string]bool, len(subset))
	for _, s := range subset {
		inSub[s] = true
	}
	var comp []string
	for _, item := range items {
		if !inSub[item] {
			comp = append(comp, item)
		}
	}
	return comp
}

// SymmetricDifference returns elements in exactly one of a or b.
func SymmetricDifference(a, b []string) []string {
	setA := toSet(a)
	setB := toSet(b)
	var result []string
	for _, v := range a {
		if !setB[v] {
			result = append(result, v)
		}
	}
	for _, v := range b {
		if !setA[v] {
			result = append(result, v)
		}
	}
	return result
}

// Intersection returns elements present in both a and b.
func Intersection(a, b []string) []string {
	setB := toSet(b)
	var result []string
	for _, v := range a {
		if setB[v] {
			result = append(result, v)
		}
	}
	return result
}

// Union returns all distinct elements from a and b (preserving first-seen order).
func Union(a, b []string) []string {
	seen := map[string]bool{}
	var result []string
	for _, v := range a {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	for _, v := range b {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

func toSet(s []string) map[string]bool {
	m := make(map[string]bool, len(s))
	for _, v := range s {
		m[v] = true
	}
	return m
}
