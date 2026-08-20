package count

// SurjectiveCount returns the number of surjective (onto) functions from
// an n-element set to a k-element set: k! * S(n,k) where S is Stirling second kind.
func SurjectiveCount(n, k int) int {
	if k <= 0 || k > n {
		return 0
	}
	// Inclusion-exclusion: sum_{i=0}^{k} (-1)^i * C(k,i) * (k-i)^n
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

// MultinomialCount returns n! / (k1! * k2! * ... * km!).
func MultinomialCount(n int, ks []int) int {
	result := factorial(n)
	for _, k := range ks {
		result /= factorial(k)
	}
	return result
}

// LatticePaths returns the number of lattice paths from (0,0) to (m,n)
// using only right (1,0) and up (0,1) steps: C(m+n, m).
func LatticePaths(m, n int) int {
	return binomial(m+n, m)
}

// DyckPaths returns the number of Dyck paths of length 2n: Catalan(n).
func DyckPaths(n int) int {
	if n < 0 {
		return 0
	}
	return binomial(2*n, n) / (n + 1)
}

// BallsInBoxes returns the number of ways to distribute n indistinguishable
// balls into k distinguishable boxes: C(n+k-1, k-1).
func BallsInBoxes(n, k int) int {
	if k <= 0 || n < 0 {
		return 0
	}
	return binomial(n+k-1, k-1)
}

// NecklaceCount returns the number of distinct necklaces with n beads of k colors
// (considering rotational equivalence): (1/n) * sum_{d|n} phi(d) * k^(n/d).
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
