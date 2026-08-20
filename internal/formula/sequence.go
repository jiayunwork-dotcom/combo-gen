package formula

// Fibonacci returns the n-th Fibonacci number (F(0)=0, F(1)=1).
func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// Lucas returns the n-th Lucas number (L(0)=2, L(1)=1).
func Lucas(n int) int {
	if n == 0 {
		return 2
	}
	if n == 1 {
		return 1
	}
	a, b := 2, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// Tribonacci returns the n-th Tribonacci number (T(0)=0, T(1)=0, T(2)=1).
func Tribonacci(n int) int {
	if n < 2 {
		return 0
	}
	if n == 2 {
		return 1
	}
	a, b, c := 0, 0, 1
	for i := 3; i <= n; i++ {
		a, b, c = b, c, a+b+c
	}
	return c
}

// EulerianNumber returns A(n, k): the number of permutations of 1..n with
// exactly k ascents. Uses the recurrence A(n,k) = (k+1)*A(n-1,k) + (n-k)*A(n-1,k-1).
func EulerianNumber(n, k int) int {
	if n <= 0 || k < 0 || k >= n {
		return 0
	}
	if k == 0 {
		return 1
	}
	// DP
	prev := make([]int, n)
	prev[0] = 1
	for i := 2; i <= n; i++ {
		cur := make([]int, n)
		cur[0] = 1
		for j := 1; j < i; j++ {
			cur[j] = (j+1)*prev[j] + (i-j)*prev[j-1]
		}
		prev = cur
	}
	return prev[k]
}

// NarayanaNumber returns N(n, k): the number of expressions containing n pairs
// of parentheses with k distinct nestings.
// N(n,k) = (1/n) * C(n,k) * C(n,k-1).
func NarayanaNumber(n, k int) int {
	if n <= 0 || k <= 0 || k > n {
		return 0
	}
	return Binomial(n, k) * Binomial(n, k-1) / n
}

// MotzkinNumber returns the n-th Motzkin number.
// M(n) = sum over k of C(n, 2k) * Catalan(k).
func MotzkinNumber(n int) int {
	if n < 0 {
		return 0
	}
	if n <= 1 {
		return 1
	}
	// Recurrence: M(n) = M(n-1) + sum_{k=0}^{n-2} M(k) * M(n-2-k)
	m := make([]int, n+1)
	m[0] = 1
	m[1] = 1
	for i := 2; i <= n; i++ {
		m[i] = m[i-1]
		for k := 0; k <= i-2; k++ {
			m[i] += m[k] * m[i-2-k]
		}
	}
	return m[n]
}

// PartitionNumber returns p(n), the number of integer partitions of n.
func PartitionNumber(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	for k := 1; k <= n; k++ {
		for i := k; i <= n; i++ {
			dp[i] += dp[i-k]
		}
	}
	return dp[n]
}
