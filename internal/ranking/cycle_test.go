package ranking

import "testing"

func TestCycleNotation(t *testing.T) {
	// perm [1,2,0] has one cycle: (0 1 2)
	cycles := CycleNotation([]int{1, 2, 0})
	if len(cycles) != 1 {
		t.Fatalf("expected 1 cycle, got %d", len(cycles))
	}
	if len(cycles[0]) != 3 {
		t.Errorf("cycle length expected 3, got %d", len(cycles[0]))
	}
}

func TestCycleNotation_Identity(t *testing.T) {
	cycles := CycleNotation([]int{0, 1, 2})
	// Identity has 3 fixed-point cycles of length 1
	if len(cycles) != 3 {
		t.Errorf("identity expected 3 cycles, got %d", len(cycles))
	}
}

func TestCycleCount(t *testing.T) {
	if CycleCount([]int{1, 0, 3, 2}) != 2 {
		t.Error("expected 2 cycles")
	}
}

func TestFromCycleNotation(t *testing.T) {
	cycles := [][]int{{0, 1, 2}}
	perm := FromCycleNotation(3, cycles)
	expected := []int{1, 2, 0}
	for i, v := range expected {
		if perm[i] != v {
			t.Errorf("perm[%d]=%d, want %d", i, perm[i], v)
		}
	}
}

func TestOrder(t *testing.T) {
	// (0 1 2) has order 3
	if Order([]int{1, 2, 0}) != 3 {
		t.Error("expected order 3")
	}
	// Identity has order 1
	if Order([]int{0, 1, 2}) != 1 {
		t.Error("identity order should be 1")
	}
}

func TestFixedPoints(t *testing.T) {
	fp := FixedPoints([]int{0, 2, 1, 3})
	// 0 and 3 are fixed
	if len(fp) != 2 || fp[0] != 0 || fp[1] != 3 {
		t.Errorf("fixed points: %v", fp)
	}
}

func TestPowerPerm(t *testing.T) {
	perm := []int{1, 2, 0}
	p2 := PowerPerm(perm, 2)
	// perm^2: 0→1→2, 1→2→0, 2→0→1 → [2,0,1]
	expected := []int{2, 0, 1}
	for i, v := range expected {
		if p2[i] != v {
			t.Errorf("p2[%d]=%d, want %d", i, p2[i], v)
		}
	}
}
