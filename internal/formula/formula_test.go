package formula

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct{ n, want int }{
		{0, 1}, {1, 1}, {5, 120}, {10, 3628800},
	}
	for _, tt := range tests {
		if got := Factorial(tt.n); got != tt.want {
			t.Errorf("Factorial(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestBinomial(t *testing.T) {
	tests := []struct{ n, k, want int }{
		{5, 2, 10}, {10, 3, 120}, {0, 0, 1}, {5, 0, 1}, {5, 5, 1}, {3, 4, 0},
	}
	for _, tt := range tests {
		if got := Binomial(tt.n, tt.k); got != tt.want {
			t.Errorf("C(%d,%d) = %d, want %d", tt.n, tt.k, got, tt.want)
		}
	}
}

func TestMultinomial(t *testing.T) {
	// 6! / (2! * 2! * 2!) = 90
	got := Multinomial(6, []int{2, 2, 2})
	if got != 90 {
		t.Errorf("Multinomial(6,[2,2,2]) = %d, want 90", got)
	}
}

func TestMultinomial_Invalid(t *testing.T) {
	if Multinomial(5, []int{2, 2}) != 0 {
		t.Error("sum != n should return 0")
	}
}

func TestCatalan(t *testing.T) {
	tests := []struct{ n, want int }{
		{0, 1}, {1, 1}, {2, 2}, {3, 5}, {4, 14}, {5, 42},
	}
	for _, tt := range tests {
		if got := Catalan(tt.n); got != tt.want {
			t.Errorf("Catalan(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestStirlingFirst(t *testing.T) {
	// |s(4,2)| = 11
	if got := StirlingFirst(4, 2); got != 11 {
		t.Errorf("StirlingFirst(4,2) = %d, want 11", got)
	}
}

func TestStirlingSecond(t *testing.T) {
	// S(4,2) = 7
	if got := StirlingSecond(4, 2); got != 7 {
		t.Errorf("StirlingSecond(4,2) = %d, want 7", got)
	}
}

func TestBell(t *testing.T) {
	tests := []struct{ n, want int }{
		{0, 1}, {1, 1}, {2, 2}, {3, 5}, {4, 15}, {5, 52},
	}
	for _, tt := range tests {
		if got := Bell(tt.n); got != tt.want {
			t.Errorf("Bell(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestDerangement(t *testing.T) {
	tests := []struct{ n, want int }{
		{0, 1}, {1, 0}, {2, 1}, {3, 2}, {4, 9}, {5, 44},
	}
	for _, tt := range tests {
		if got := Derangement(tt.n); got != tt.want {
			t.Errorf("D(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestFallingFactorial(t *testing.T) {
	// 5 * 4 * 3 = 60
	if got := FallingFactorial(5, 3); got != 60 {
		t.Errorf("FallingFactorial(5,3) = %d, want 60", got)
	}
}

func TestRisingFactorial(t *testing.T) {
	// 3 * 4 * 5 = 60
	if got := RisingFactorial(3, 3); got != 60 {
		t.Errorf("RisingFactorial(3,3) = %d, want 60", got)
	}
}
