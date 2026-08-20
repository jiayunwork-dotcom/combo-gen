// Package ranking provides permutation ranking (index) and unranking (decode)
// using the Lehmer code (factorial number system).
package ranking

import "errors"

var (
	ErrInvalidPerm = errors.New("ranking: invalid permutation")
	ErrRankTooLarge = errors.New("ranking: rank exceeds n!")
	ErrNegativeRank = errors.New("ranking: rank must be non-negative")
)

// Rank returns the lexicographic rank (0-based) of a permutation of 0..n-1.
// The permutation must contain each integer in [0, n) exactly once.
func Rank(perm []int) (int, error) {
	n := len(perm)
	if n == 0 {
		return 0, nil
	}
	if err := validatePerm(perm); err != nil {
		return 0, err
	}
	// Compute Lehmer code
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
	// Convert Lehmer to rank
	rank := 0
	fact := 1
	for i := n - 2; i >= 0; i-- {
		fact *= (n - 1 - i)
		rank += lehmer[i] * fact
	}
	return rank, nil
}

// Unrank returns the permutation of 0..n-1 at the given lexicographic rank.
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
	// Decompose rank into Lehmer code
	lehmer := make([]int, n)
	rem := rank
	for i := 0; i < n; i++ {
		fact := factorial(n - 1 - i)
		lehmer[i] = rem / fact
		rem %= fact
	}
	// Reconstruct permutation from Lehmer code
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

// LehmerCode returns the Lehmer code representation of a permutation.
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

// InversionCount returns the number of inversions in the permutation.
// An inversion is a pair (i,j) with i < j but perm[i] > perm[j].
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

// IsEvenPermutation returns true if the permutation has an even number of inversions.
func IsEvenPermutation(perm []int) bool {
	return InversionCount(perm)%2 == 0
}

// NextPermutation generates the next permutation in lexicographic order (in-place).
// Returns false if already the last permutation.
func NextPermutation(perm []int) bool {
	n := len(perm)
	if n < 2 {
		return false
	}
	// Find largest i such that perm[i] < perm[i+1]
	i := n - 2
	for i >= 0 && perm[i] >= perm[i+1] {
		i--
	}
	if i < 0 {
		return false
	}
	// Find largest j such that perm[j] > perm[i]
	j := n - 1
	for perm[j] <= perm[i] {
		j--
	}
	perm[i], perm[j] = perm[j], perm[i]
	// Reverse perm[i+1:]
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
