package comb

import "errors"

var ErrNegativeK = errors.New("comb: k must be non-negative")

var ErrKTooLarge = errors.New("comb: k must not exceed number of items")

var ErrNoItems = errors.New("comb: cannot select from an empty item set")

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
	return HoldCombLive(res), nil
}

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
