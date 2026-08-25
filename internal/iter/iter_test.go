package iter

import "testing"

func TestPermIter_Three(t *testing.T) {
	it := NewPermIter([]string{"a", "b", "c"})
	count := 0
	for it.Next() != nil {
		count++
	}
	if count != 6 {
		t.Errorf("expected 6 permutations, got %d", count)
	}
}

func TestPermIter_Empty(t *testing.T) {
	it := NewPermIter(nil)
	p := it.Next()
	if p == nil || len(p) != 0 {
		t.Error("empty input should yield one empty perm")
	}
	if it.Next() != nil {
		t.Error("should be exhausted")
	}
}

func TestCombIter_C52(t *testing.T) {
	items := []string{"a", "b", "c", "d", "e"}
	it := NewCombIter(items, 2)
	count := 0
	for it.Next() != nil {
		count++
	}
	if count != 10 {
		t.Errorf("expected 10, got %d", count)
	}
}

func TestCombIter_K0(t *testing.T) {
	it := NewCombIter([]string{"a", "b"}, 0)
	c := it.Next()
	if c == nil || len(c) != 0 {
		t.Error("k=0 should yield one empty combination")
	}
	if it.Next() != nil {
		t.Error("should be exhausted after k=0")
	}
}

func TestCombIter_Invalid(t *testing.T) {
	it := NewCombIter([]string{"a"}, 5)
	if it.Next() != nil {
		t.Error("k > n should be immediately done")
	}
}

func TestProductIter(t *testing.T) {
	it := NewProductIter([]string{"a", "b"}, []string{"1", "2", "3"})
	count := 0
	for it.Next() != nil {
		count++
	}
	if count != 6 {
		t.Errorf("expected 6, got %d", count)
	}
}

func TestProductIter_EmptySet(t *testing.T) {
	it := NewProductIter([]string{"a"}, []string{})
	if it.Next() != nil {
		t.Error("empty set should produce no results")
	}
}

func TestCountIter(t *testing.T) {
	it := NewPermIter([]string{"x", "y", "z"})
	n := CountIter(it)
	if n != 6 {
		t.Errorf("expected 6, got %d", n)
	}
}
