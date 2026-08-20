package formula

import "testing"

func TestFibonacci(t *testing.T) {
	tests := []struct{ n, want int }{
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {10, 55},
	}
	for _, tt := range tests {
		if got := Fibonacci(tt.n); got != tt.want {
			t.Errorf("Fib(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestLucas(t *testing.T) {
	tests := []struct{ n, want int }{
		{0, 2}, {1, 1}, {2, 3}, {3, 4}, {4, 7},
	}
	for _, tt := range tests {
		if got := Lucas(tt.n); got != tt.want {
			t.Errorf("Lucas(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestEulerianNumber(t *testing.T) {
	// A(4,1) = 11
	if got := EulerianNumber(4, 1); got != 11 {
		t.Errorf("A(4,1) = %d, want 11", got)
	}
	// A(3,0) = 1
	if got := EulerianNumber(3, 0); got != 1 {
		t.Errorf("A(3,0) = %d, want 1", got)
	}
}

func TestNarayanaNumber(t *testing.T) {
	// N(4,2) = 6
	if got := NarayanaNumber(4, 2); got != 6 {
		t.Errorf("N(4,2) = %d, want 6", got)
	}
}

func TestMotzkinNumber(t *testing.T) {
	tests := []struct{ n, want int }{
		{0, 1}, {1, 1}, {2, 2}, {3, 4}, {4, 9},
	}
	for _, tt := range tests {
		if got := MotzkinNumber(tt.n); got != tt.want {
			t.Errorf("Motzkin(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestPartitionNumber(t *testing.T) {
	tests := []struct{ n, want int }{
		{0, 1}, {1, 1}, {4, 5}, {5, 7}, {10, 42},
	}
	for _, tt := range tests {
		if got := PartitionNumber(tt.n); got != tt.want {
			t.Errorf("p(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}
