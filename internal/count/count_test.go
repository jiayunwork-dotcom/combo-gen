package count

import "testing"

func TestFactorial(t *testing.T) {
	cases := []struct{ n, want int }{
		{0, 1}, {1, 1}, {2, 2}, {3, 6}, {4, 24}, {5, 120}, {6, 720},
	}
	for _, tc := range cases {
		if got := Factorial(tc.n); got != tc.want {
			t.Errorf("Factorial(%d) = %d, want %d", tc.n, got, tc.want)
		}
	}
}

func TestFactorialNegative(t *testing.T) {
	if Factorial(-1) != 0 {
		t.Error("Factorial(-1) should be 0")
	}
}

func TestBinomial(t *testing.T) {
	cases := []struct{ n, k, want int }{
		{5, 2, 10}, {6, 3, 20}, {4, 0, 1}, {4, 4, 1}, {10, 5, 252},
	}
	for _, tc := range cases {
		if got := Binomial(tc.n, tc.k); got != tc.want {
			t.Errorf("Binomial(%d,%d) = %d, want %d", tc.n, tc.k, got, tc.want)
		}
	}
}

func TestBinomialInvalid(t *testing.T) {
	if Binomial(3, -1) != 0 {
		t.Error("Binomial(3,-1) should be 0")
	}
	if Binomial(3, 5) != 0 {
		t.Error("Binomial(3,5) should be 0")
	}
}

func TestCombinationCount(t *testing.T) {
	if CombinationCount(5, 3) != 10 {
		t.Errorf("C(5,3) = %d", CombinationCount(5, 3))
	}
}

func TestCombinationWithRepCount(t *testing.T) {
	if got := CombinationWithRepCount(3, 2); got != 6 {
		t.Errorf("CombWithRep(3,2) = %d, want 6", got)
	}
}

func TestPermutationCount(t *testing.T) {
	if got := PermutationCount(5, 3); got != 60 {
		t.Errorf("P(5,3) = %d, want 60", got)
	}
}

func TestFullPermutationCount(t *testing.T) {
	if got := FullPermutationCount(4); got != 24 {
		t.Errorf("4! = %d, want 24", got)
	}
}

func TestCartesianProductCount(t *testing.T) {
	if got := CartesianProductCount(2, 3, 4); got != 24 {
		t.Errorf("CartProd(2,3,4) = %d, want 24", got)
	}
	if CartesianProductCount() != 0 {
		t.Error("empty should be 0")
	}
}

func TestMultisetPermutationCount(t *testing.T) {
	if got := MultisetPermutationCount([]int{2, 1}); got != 3 {
		t.Errorf("multiset{2,1} = %d, want 3", got)
	}
	if got := MultisetPermutationCount([]int{3, 1}); got != 4 {
		t.Errorf("multiset{3,1} = %d, want 4", got)
	}
}

func TestCountMatchesEnumerationCombinations(t *testing.T) {
	n, k := 6, 3
	expected := CombinationCount(n, k)
	if expected != 20 {
		t.Errorf("C(6,3) = %d, want 20", expected)
	}
}

func TestCountMatchesEnumerationPermK(t *testing.T) {
	n, k := 4, 2
	expected := PermutationCount(n, k)
	if expected != 12 {
		t.Errorf("P(4,2) = %d, want 12", expected)
	}
}
