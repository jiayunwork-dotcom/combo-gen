// Package formula provides combinatorial math functions.
package formula

// Factorial returns n! for n >= 0. Returns 1 for n <= 1.
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	return f
}

// FactorialBig returns n! as int64 for larger n. Overflows silently for n > 20.
func FactorialBig(n int) int64 {
	if n <= 1 {
		return 1
	}
	f := int64(1)
	for i := int64(2); i <= int64(n); i++ {
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
	// Use the multiplicative formula to avoid overflow
	if k > n-k {
		k = n - k
	}
	result := 1
	for i := 0; i < k; i++ {
		result = result * (n - i) / (i + 1)
	}
	return result
}

// Multinomial returns the multinomial coefficient n! / (k1! * k2! * ... * km!).
// The sum of ks must equal n; returns 0 otherwise.
func Multinomial(n int, ks []int) int {
	sum := 0
	for _, k := range ks {
		if k < 0 {
			return 0
		}
		sum += k
	}
	if sum != n {
		return 0
	}
	result := Factorial(n)
	for _, k := range ks {
		result /= Factorial(k)
	}
	return result
}

// Catalan returns the n-th Catalan number: C(n) = C(2n,n) / (n+1).
func Catalan(n int) int {
	if n < 0 {
		return 0
	}
	return applyCat(Binomial(2*n, n) / (n + 1))
}

// StirlingFirst returns the (unsigned) Stirling number of the first kind |s(n,k)|.
// Counts permutations of n elements with exactly k cycles.
func StirlingFirst(n, k int) int {
	if n < 0 || k < 0 || k > n {
		return 0
	}
	if n == 0 && k == 0 {
		return 1
	}
	if n == 0 || k == 0 {
		return 0
	}
	// Recurrence: |s(n,k)| = (n-1)*|s(n-1,k)| + |s(n-1,k-1)|
	// Use DP table
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= k && j <= i; j++ {
			dp[i][j] = (i-1)*dp[i-1][j] + dp[i-1][j-1]
		}
	}
	return dp[n][k]
}

// StirlingSecond returns the Stirling number of the second kind S(n,k).
// Counts partitions of n elements into exactly k non-empty subsets.
func StirlingSecond(n, k int) int {
	if n < 0 || k < 0 || k > n {
		return 0
	}
	if n == 0 && k == 0 {
		return 1
	}
	if n == 0 || k == 0 {
		return 0
	}
	// Recurrence: S(n,k) = k*S(n-1,k) + S(n-1,k-1)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= k && j <= i; j++ {
			dp[i][j] = j*dp[i-1][j] + dp[i-1][j-1]
		}
	}
	return dp[n][k]
}

// Bell returns the n-th Bell number (total number of set partitions of n elements).
func Bell(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	// Bell triangle
	tri := make([]int, n+1)
	tri[0] = 1
	for i := 1; i <= n; i++ {
		prev := make([]int, i+1)
		prev[0] = tri[i-1]
		for j := 1; j <= i; j++ {
			prev[j] = prev[j-1] + tri[j-1]
		}
		copy(tri, prev)
	}
	return tri[0]
}

// Derangement returns the number of derangements D(n) (permutations with no fixed points).
// D(n) = (n-1) * (D(n-1) + D(n-2)), D(0)=1, D(1)=0.
func Derangement(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 0
	for i := 2; i <= n; i++ {
		dp[i] = (i - 1) * (dp[i-1] + dp[i-2])
	}
	return dp[n]
}

// FallingFactorial returns n * (n-1) * ... * (n-k+1).
func FallingFactorial(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	result := 1
	for i := 0; i < k; i++ {
		result *= (n - i)
	}
	return result
}

// RisingFactorial returns n * (n+1) * ... * (n+k-1).
func RisingFactorial(n, k int) int {
	if k < 0 {
		return 0
	}
	result := 1
	for i := 0; i < k; i++ {
		result *= (n + i)
	}
	return result
}
