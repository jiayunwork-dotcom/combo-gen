package perm

import "errors"

var ErrNegativeK = errors.New("perm: k must be non-negative")

var ErrKTooLarge = errors.New("perm: k must not exceed number of items")

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

func factorial(n int) int {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	return f
}
