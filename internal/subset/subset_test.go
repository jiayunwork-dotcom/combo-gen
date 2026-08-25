package subset

import "testing"

func TestPowerSet_Three(t *testing.T) {
	ps := PowerSet([]string{"a", "b", "c"})
	if len(ps) != 8 {
		t.Errorf("expected 8 subsets, got %d", len(ps))
	}
}

func TestPowerSet_Empty(t *testing.T) {
	ps := PowerSet(nil)
	if len(ps) != 1 || len(ps[0]) != 0 {
		t.Error("empty input should yield one empty subset")
	}
}

func TestSubsetsOfSize(t *testing.T) {
	items := []string{"a", "b", "c", "d"}
	subs := SubsetsOfSize(items, 2)
	if len(subs) != 6 {
		t.Errorf("expected 6 subsets of size 2, got %d", len(subs))
	}
}

func TestSubsetsOfSize_Zero(t *testing.T) {
	subs := SubsetsOfSize([]string{"a", "b"}, 0)
	if len(subs) != 1 || len(subs[0]) != 0 {
		t.Error("k=0 should return one empty subset")
	}
}

func TestSubsetsOfSizeRange(t *testing.T) {
	items := []string{"a", "b", "c"}
	subs := SubsetsOfSizeRange(items, 1, 2)
	if len(subs) != 6 {
		t.Errorf("expected 6, got %d", len(subs))
	}
}

func TestNonEmptySubsets(t *testing.T) {
	items := []string{"x", "y"}
	subs := NonEmptySubsets(items)
	if len(subs) != 3 {
		t.Errorf("expected 3, got %d", len(subs))
	}
}

func TestComplement(t *testing.T) {
	items := []string{"a", "b", "c", "d"}
	comp := Complement(items, []string{"b", "d"})
	if len(comp) != 2 || comp[0] != "a" || comp[1] != "c" {
		t.Errorf("complement: %v", comp)
	}
}

func TestSymmetricDifference(t *testing.T) {
	a := []string{"a", "b", "c"}
	b := []string{"b", "c", "d"}
	sd := SymmetricDifference(a, b)
	if len(sd) != 2 {
		t.Errorf("expected 2, got %v", sd)
	}
}

func TestIntersection(t *testing.T) {
	a := []string{"a", "b", "c"}
	b := []string{"b", "c", "d"}
	inter := Intersection(a, b)
	if len(inter) != 2 {
		t.Errorf("expected 2, got %v", inter)
	}
}

func TestUnion(t *testing.T) {
	a := []string{"a", "b"}
	b := []string{"b", "c"}
	u := Union(a, b)
	if len(u) != 3 {
		t.Errorf("expected 3, got %v", u)
	}
}
