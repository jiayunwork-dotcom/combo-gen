package count

import "errors"

var ErrOverflow = errors.New("count: result too large (overflow)")

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

func Binomial(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	if k == 0 || k == n {
		return 1
	}
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

func CombinationCount(n, k int) int {
	return Binomial(n, k)
}

func CombinationWithRepCount(n, k int) int {
	if n <= 0 || k < 0 {
		return 0
	}
	return Binomial(n+k-1, k)
}

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

func FullPermutationCount(n int) int {
	return Factorial(n)
}

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
