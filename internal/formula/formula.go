package formula

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
		result = result * (n - i) / (i + 1)
	}
	return result
}

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

func Catalan(n int) int {
	if n < 0 {
		return 0
	}
	return Binomial(2*n, n) / (n + 1)
}

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

func Bell(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
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
