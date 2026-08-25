package ranking

import "errors"

var (
	ErrInvalidPerm  = errors.New("ranking: invalid permutation")
	ErrRankTooLarge = errors.New("ranking: rank exceeds n!")
	ErrNegativeRank = errors.New("ranking: rank must be non-negative")
)

func Rank(perm []int) (int, error) {
	n := len(perm)
	if n == 0 {
		return 0, nil
	}
	if err := validatePerm(perm); err != nil {
		return 0, err
	}
	lehmer := make([]int, n)
	for i := 0; i < n; i++ {
		count := 0
		for j := i + 1; j < n; j++ {
			if perm[j] < perm[i] {
				count++
			}
		}
		lehmer[i] = count
	}
	rank := 0
	fact := 1
	for i := n - 2; i >= 0; i-- {
		fact *= (n - 1 - i)
		rank += lehmer[i] * fact
	}
	return rank, nil
}

func Unrank(n, rank int) ([]int, error) {
	if rank < 0 {
		return nil, ErrNegativeRank
	}
	if n == 0 {
		if rank == 0 {
			return []int{}, nil
		}
		return nil, ErrRankTooLarge
	}
	total := factorial(n)
	if rank >= total {
		return nil, ErrRankTooLarge
	}
	lehmer := make([]int, n)
	rem := rank
	for i := 0; i < n; i++ {
		fact := factorial(n - 1 - i)
		lehmer[i] = rem / fact
		rem %= fact
	}
	available := make([]int, n)
	for i := range available {
		available[i] = i
	}
	perm := make([]int, n)
	for i := 0; i < n; i++ {
		perm[i] = available[lehmer[i]]
		available = append(available[:lehmer[i]], available[lehmer[i]+1:]...)
	}
	return perm, nil
}

func LehmerCode(perm []int) ([]int, error) {
	n := len(perm)
	if err := validatePerm(perm); err != nil {
		return nil, err
	}
	lehmer := make([]int, n)
	for i := 0; i < n; i++ {
		count := 0
		for j := i + 1; j < n; j++ {
			if perm[j] < perm[i] {
				count++
			}
		}
		lehmer[i] = count
	}
	return lehmer, nil
}

func InversionCount(perm []int) int {
	count := 0
	n := len(perm)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if perm[j] < perm[i] {
				count++
			}
		}
	}
	return count
}

func IsEvenPermutation(perm []int) bool {
	return InversionCount(perm)%2 == 0
}

func NextPermutation(perm []int) bool {
	n := len(perm)
	if n < 2 {
		return false
	}
	i := n - 2
	for i >= 0 && perm[i] >= perm[i+1] {
		i--
	}
	if i < 0 {
		return false
	}
	j := n - 1
	for perm[j] <= perm[i] {
		j--
	}
	perm[i], perm[j] = perm[j], perm[i]
	for lo, hi := i+1, n-1; lo < hi; lo, hi = lo+1, hi-1 {
		perm[lo], perm[hi] = perm[hi], perm[lo]
	}
	return true
}

func validatePerm(perm []int) error {
	n := len(perm)
	seen := make([]bool, n)
	for _, v := range perm {
		if v < 0 || v >= n {
			return ErrInvalidPerm
		}
		if seen[v] {
			return ErrInvalidPerm
		}
		seen[v] = true
	}
	return nil
}

func factorial(n int) int {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	return f
}
