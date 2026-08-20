// Package count provides combinatorial counting formulas.
//
// These functions compute the expected count of results for each generator,
// allowing tests to verify that the enumeration is complete (no duplicates
// missed or extra entries generated). The count formulas are the ground truth;
// the generators must produce exactly count results.
//
// Invariants:
//   - Combinations(n, k) = C(n, k) = n! / (k! * (n-k)!)
//   - CombinationsWithRepetition(n, k) = C(n+k-1, k)
//   - Permutations(n) = n!
//   - PermutationsK(n, k) = n! / (n-k)! = P(n, k)
//   - CartesianProduct(sets) = product of |set_i|
//   - DedupedPermutations(n, k, duplicateCounts) uses multinomial formula
package count

import "errors"

// ErrOverflow is returned when the result would exceed a practical limit.
var ErrOverflow = errors.New("count: result too large (overflow)")

// Factorial returns n! for n >= 0.
func Factorial(n int) int {
	if n < 0 {
		return 0
	}
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	return f
}

// Binomial returns C(n, k) = n! / (k! * (n-k)!).
// Returns 0 for invalid inputs (k < 0 or k > n).
func Binomial(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	if k == 0 || k == n {
		return 1
	}
	// optimize: use smaller k
	if k > n-k {
		k = n - k
	}
	result := 1
	for i := 0; i < k; i++ {
		result *= (n - i)
		result /= (i + 1)
	}
	return result
}

// CombinationCount returns C(n, k).
func CombinationCount(n, k int) int {
	return Binomial(n, k)
}

// CombinationWithRepCount returns C(n+k-1, k) for n items choosing k with repetition.
func CombinationWithRepCount(n, k int) int {
	if n <= 0 || k < 0 {
		return 0
	}
	return Binomial(n+k-1, k)
}

// PermutationCount returns P(n, k) = n! / (n-k)!.
func PermutationCount(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	result := 1
	for i := n; i > n-k; i-- {
		result *= i
	}
	return result
}

// FullPermutationCount returns n!.
func FullPermutationCount(n int) int {
	return Factorial(n)
}

// CartesianProductCount returns the product of set sizes.
func CartesianProductCount(sizes ...int) int {
	if len(sizes) == 0 {
		return 0
	}
	result := 1
	for _, s := range sizes {
		if s <= 0 {
			return 0
		}
		result *= s
	}
	return result
}

// MultisetPermutationCount returns n! / (c1! * c2! * ... * cm!) where
// ci are the counts of each distinct element. This gives the number of
// distinct permutations when there are repeated elements.
func MultisetPermutationCount(counts []int) int {
	n := 0
	for _, c := range counts {
		n += c
	}
	result := Factorial(n)
	for _, c := range counts {
		result /= Factorial(c)
	}
	return result
}
