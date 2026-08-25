package formula

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

func EulerianNumber(n, k int) int {
	if n <= 0 || k < 0 || k >= n {
		return 0
	}
	if k == 0 {
		return 1
	}
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

func NarayanaNumber(n, k int) int {
	if n <= 0 || k <= 0 || k > n {
		return 0
	}
	return Binomial(n, k) * Binomial(n, k-1) / n
}

func MotzkinNumber(n int) int {
	if n < 0 {
		return 0
	}
	if n <= 1 {
		return 1
	}
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
