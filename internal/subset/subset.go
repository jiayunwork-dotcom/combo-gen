package subset

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
	return HoldSetLive(result)
}

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

func NonEmptySubsets(items []string) [][]string {
	all := PowerSet(items)
	if len(all) == 0 {
		return nil
	}
	return all[1:]
}

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
