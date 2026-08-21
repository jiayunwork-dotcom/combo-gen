// Package comb generates combinations of a slice of strings.
package comb

import "errors"

// ErrNegativeK is returned when k is negative.
var ErrNegativeK = errors.New("comb: k must be non-negative")

// ErrKTooLarge is returned when k exceeds the number of available items.
var ErrKTooLarge = errors.New("comb: k must not exceed number of items")

// ErrNoItems is returned when a non-empty selection is requested from no items.
var ErrNoItems = errors.New("comb: cannot select from an empty item set")

// Combinations returns every unordered selection of exactly k distinct items,
// i.e. C(n, k) results for n items. Items keep their input order within each
// combination and each returned slice is a fresh copy.
//
// It returns ErrNegativeK when k < 0 and ErrKTooLarge when k > len(items).
// For k == 0 it returns a single empty combination, because there is exactly
// one way to select zero elements.
func Combinations(items []string, k int) ([][]string, error) {
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
	cur := make([]string, 0, k)
	var walk func(start int)
	walk = func(start int) {
		if len(cur) == k {
			out := make([]string, k)
			copy(out, cur)
			res = append(res, out)
			return
		}
		for i := start; i < n; i++ {
			cur = append(cur, items[i])
			walk(i + 1)
			cur = cur[:len(cur)-1]
		}
	}
	walk(0)
	return res, nil
}

// CombinationsWithRepetition returns every unordered selection of exactly k
// items where an item may be picked more than once, i.e. C(n+k-1, k) results
// for n items. Because items may repeat, k is allowed to exceed len(items).
//
// It returns ErrNegativeK when k < 0 and ErrNoItems when k > 0 but no items are
// supplied. For k == 0 it returns a single empty combination.
func CombinationsWithRepetition(items []string, k int) ([][]string, error) {
	if k < 0 {
		return nil, ErrNegativeK
	}
	if k == 0 {
		return [][]string{{}}, nil
	}
	n := len(items)
	if n == 0 {
		return nil, ErrNoItems
	}
	var res [][]string
	cur := make([]string, 0, k)
	var walk func(start int)
	walk = func(start int) {
		if len(cur) == k {
			out := make([]string, k)
			copy(out, cur)
			res = append(res, out)
			return
		}
		for i := start; i < n; i++ {
			cur = append(cur, items[i])
			walk(i)
			cur = cur[:len(cur)-1]
		}
	}
	walk(0)
	return res, nil
}
