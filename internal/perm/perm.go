// Package perm generates permutations of a slice of strings.
package perm

import "errors"

// ErrNegativeK is returned when k is negative.
var ErrNegativeK = errors.New("perm: k must be non-negative")

// ErrKTooLarge is returned when k exceeds the number of available items.
var ErrKTooLarge = errors.New("perm: k must not exceed number of items")

// Permutations returns every ordering of items, i.e. n! results for n items.
//
// Items are treated as positionally distinct: a duplicated value yields
// duplicated orderings. Each returned slice is a fresh copy and never aliases
// items. An empty or nil input yields a single empty permutation, because there
// is exactly one way to arrange zero elements (0! = 1).
func Permutations(items []string) [][]string {
	n := len(items)
	if n == 0 {
		return [][]string{{}}
	}
	res := make([][]string, 0, factorial(n))
	used := make([]bool, n)
	cur := make([]string, 0, n)
	var walk func()
	walk = func() {
		if len(cur) == n {
			out := make([]string, n)
			copy(out, cur)
			res = append(res, out)
			return
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			cur = append(cur, items[i])
			walk()
			cur = cur[:len(cur)-1]
			used[i] = false
		}
	}
	walk()
	return res
}

// PermutationsK returns every ordered selection of exactly k items, i.e.
// n!/(n-k)! results for n items.
//
// It returns ErrNegativeK when k < 0 and ErrKTooLarge when k > len(items),
// since an item may not be reused. For k == 0 it returns a single empty
// selection, because there is exactly one way to select zero elements.
func PermutationsK(items []string, k int) ([][]string, error) {
	if k < 0 {
		return nil, ErrNegativeK
	}
	n := len(items)
	if k > n {
		return nil, ErrKTooLarge
	}
	if k == 0 {
		return [][]string{{}}, nil
	}
	var res [][]string
	used := make([]bool, n)
	cur := make([]string, 0, k)
	var walk func()
	walk = func() {
		if len(cur) == k {
			out := make([]string, k)
			copy(out, cur)
			res = append(res, out)
			return
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			cur = append(cur, items[i])
			walk()
			cur = cur[:len(cur)-1]
			used[i] = false
		}
	}
	walk()
	return res, nil
}

// factorial returns n! and is only used to pre-size the result slice.
func factorial(n int) int {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	return f
}
