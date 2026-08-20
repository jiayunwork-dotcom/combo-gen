package perm

import "testing"

func TestMultisetPermutations(t *testing.T) {
	// "aab" has 3!/2! = 3 distinct permutations
	perms := MultisetPermutations([]string{"a", "a", "b"})
	if len(perms) != 3 {
		t.Errorf("expected 3, got %d", len(perms))
	}
}

func TestMultisetPermutations_AllSame(t *testing.T) {
	perms := MultisetPermutations([]string{"x", "x", "x"})
	if len(perms) != 1 {
		t.Errorf("all same should give 1, got %d", len(perms))
	}
}

func TestCircularPermutations(t *testing.T) {
	perms := CircularPermutations([]string{"a", "b", "c"})
	// (3-1)! = 2
	if len(perms) != 2 {
		t.Errorf("expected 2 circular perms, got %d", len(perms))
	}
}

func TestDerangements_Three(t *testing.T) {
	derangs := Derangements([]string{"a", "b", "c"})
	// D(3) = 2
	if len(derangs) != 2 {
		t.Errorf("expected 2 derangements, got %d", len(derangs))
	}
}

func TestDerangements_One(t *testing.T) {
	derangs := Derangements([]string{"a"})
	if derangs != nil {
		t.Error("single element has no derangement")
	}
}

func TestInversePerm(t *testing.T) {
	perm := []int{2, 0, 1}
	inv := InversePerm(perm)
	// perm sends 0→2, 1→0, 2→1; inverse: 0→1, 1→2, 2→0
	expected := []int{1, 2, 0}
	for i, v := range expected {
		if inv[i] != v {
			t.Errorf("inv[%d]=%d, want %d", i, inv[i], v)
		}
	}
}

func TestComposePerm(t *testing.T) {
	a := []int{1, 2, 0}
	b := []int{2, 0, 1}
	c := ComposePerm(a, b)
	// c[0] = b[a[0]] = b[1] = 0
	// c[1] = b[a[1]] = b[2] = 1
	// c[2] = b[a[2]] = b[0] = 2
	expected := []int{0, 1, 2}
	for i, v := range expected {
		if c[i] != v {
			t.Errorf("compose[%d]=%d, want %d", i, c[i], v)
		}
	}
}
