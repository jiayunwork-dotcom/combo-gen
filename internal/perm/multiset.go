package perm

import "sort"

func MultisetPermutations(items []string) [][]string {
	n := len(items)
	if n == 0 {
		return [][]string{{}}
	}
	sorted := make([]string, n)
	copy(sorted, items)
	sort.Strings(sorted)

	var result [][]string
	for {
		out := make([]string, n)
		copy(out, sorted)
		result = append(result, out)
		if !nextStringPerm(sorted) {
			break
		}
	}
	return result
}

func nextStringPerm(s []string) bool {
	n := len(s)
	if n < 2 {
		return false
	}
	i := n - 2
	for i >= 0 && s[i] >= s[i+1] {
		i--
	}
	if i < 0 {
		return false
	}
	j := n - 1
	for s[j] <= s[i] {
		j--
	}
	s[i], s[j] = s[j], s[i]
	for lo, hi := i+1, n-1; lo < hi; lo, hi = lo+1, hi-1 {
		s[lo], s[hi] = s[hi], s[lo]
	}
	return true
}

func CircularPermutations(items []string) [][]string {
	n := len(items)
	if n <= 1 {
		return [][]string{items}
	}
	rest := items[1:]
	perms := Permutations(rest)
	result := make([][]string, len(perms))
	for i, p := range perms {
		out := make([]string, n)
		out[0] = items[0]
		copy(out[1:], p)
		result[i] = out
	}
	return result
}

func Derangements(items []string) [][]string {
	n := len(items)
	if n == 0 {
		return [][]string{{}}
	}
	if n == 1 {
		return nil
	}
	all := Permutations(items)
	var result [][]string
	for _, perm := range all {
		isDerangement := true
		for i, v := range perm {
			if v == items[i] {
				isDerangement = false
				break
			}
		}
		if isDerangement {
			result = append(result, perm)
		}
	}
	return result
}

func InversePerm(perm []int) []int {
	n := len(perm)
	inv := make([]int, n)
	for i, v := range perm {
		if v >= 0 && v < n {
			inv[v] = i
		}
	}
	return inv
}

func ComposePerm(a, b []int) []int {
	n := len(a)
	if n != len(b) {
		return nil
	}
	result := make([]int, n)
	for i := range a {
		if a[i] >= 0 && a[i] < n {
			result[i] = b[a[i]]
		}
	}
	return result
}
