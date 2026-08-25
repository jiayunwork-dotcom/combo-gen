package count

func SurjectiveCount(n, k int) int {
	if k <= 0 || k > n {
		return 0
	}
	sum := 0
	for i := 0; i <= k; i++ {
		term := binomial(k, i) * pow(k-i, n)
		if i%2 == 0 {
			sum += term
		} else {
			sum -= term
		}
	}
	return sum
}

func MultinomialCount(n int, ks []int) int {
	result := factorial(n)
	for _, k := range ks {
		result /= factorial(k)
	}
	return result
}

func LatticePaths(m, n int) int {
	return binomial(m+n, m)
}

func DyckPaths(n int) int {
	if n < 0 {
		return 0
	}
	return binomial(2*n, n) / (n + 1)
}

func BallsInBoxes(n, k int) int {
	if k <= 0 || n < 0 {
		return 0
	}
	return binomial(n+k-1, k-1)
}

func NecklaceCount(n, k int) int {
	if n <= 0 || k <= 0 {
		return 0
	}
	if n == 1 {
		return k
	}
	sum := 0
	for d := 1; d <= n; d++ {
		if n%d == 0 {
			sum += eulerPhi(d) * pow(k, n/d)
		}
	}
	return sum / n
}

func eulerPhi(n int) int {
	result := n
	p := 2
	temp := n
	for p*p <= temp {
		if temp%p == 0 {
			for temp%p == 0 {
				temp /= p
			}
			result -= result / p
		}
		p++
	}
	if temp > 1 {
		result -= result / temp
	}
	return result
}

func pow(base, exp int) int {
	if exp == 0 {
		return 1
	}
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

func binomial(n, k int) int {
	if k < 0 || k > n {
		return 0
	}
	if k > n-k {
		k = n - k
	}
	r := 1
	for i := 0; i < k; i++ {
		r = r * (n - i) / (i + 1)
	}
	return r
}

func factorial(n int) int {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	return f
}
