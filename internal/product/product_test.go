package product

import (
	"strings"
	"testing"
)

func flatten(got [][]string) []string {
	out := make([]string, 0, len(got))
	for _, g := range got {
		out = append(out, strings.Join(g, ""))
	}
	return out
}

func sameStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestCartesianProductTwoSets(t *testing.T) {
	got, err := CartesianProduct([]string{"a", "b"}, []string{"1", "2"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []string{"a1", "a2", "b1", "b2"}
	if g := flatten(got); !sameStrings(g, want) {
		t.Errorf("product = %v, want %v", g, want)
	}
}

func TestCartesianProductThreeSets(t *testing.T) {
	got, err := CartesianProduct([]string{"a", "b"}, []string{"1", "2"}, []string{"x", "y"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 8 {
		t.Fatalf("len = %d, want 8", len(got))
	}
	for _, g := range got {
		if len(g) != 3 {
			t.Fatalf("tuple %v has len %d, want 3", g, len(g))
		}
	}
}

func TestCartesianProductSingleSet(t *testing.T) {
	got, err := CartesianProduct([]string{"a", "b", "c"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []string{"a", "b", "c"}
	if g := flatten(got); !sameStrings(g, want) {
		t.Errorf("product = %v, want %v", g, want)
	}
}

func TestCartesianProductUnevenSizes(t *testing.T) {
	got, err := CartesianProduct([]string{"a"}, []string{"1", "2", "3"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []string{"a1", "a2", "a3"}
	if g := flatten(got); !sameStrings(g, want) {
		t.Errorf("product = %v, want %v", g, want)
	}
}

func TestCartesianProductNoSets(t *testing.T) {
	got, err := CartesianProduct()
	if err != ErrNoSets {
		t.Errorf("err = %v, want ErrNoSets", err)
	}
	if got != nil {
		t.Errorf("got = %v, want nil", got)
	}
}

func TestCartesianProductEmptyFactor(t *testing.T) {
	got, err := CartesianProduct([]string{"a", "b"}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 0 {
		t.Errorf("got = %v, want empty result", got)
	}
}

func TestCartesianProductTuplesAreIndependent(t *testing.T) {
	first := []string{"a", "b"}
	got, err := CartesianProduct(first, []string{"1"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("len = %d, want 2", len(got))
	}
	got[0][0] = "mutated"
	if first[0] != "a" {
		t.Errorf("input mutated: first[0] = %q", first[0])
	}
	if got[1][0] == "mutated" {
		t.Error("tuples alias each other")
	}
}
