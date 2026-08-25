package perm

import (
	"sort"
	"strings"
	"testing"
)

func flatten(got [][]string) []string {
	out := make([]string, 0, len(got))
	for _, g := range got {
		out = append(out, strings.Join(g, ""))
	}
	sort.Strings(out)
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

func TestPermutationsCount(t *testing.T) {
	got := Permutations([]string{"a", "b", "c"})
	if len(got) != 6 {
		t.Fatalf("len = %d, want 6", len(got))
	}
	want := []string{"abc", "acb", "bac", "bca", "cab", "cba"}
	if g := flatten(got); !sameStrings(g, want) {
		t.Errorf("orderings = %v, want %v", g, want)
	}
}

func TestPermutationsSingleItem(t *testing.T) {
	got := Permutations([]string{"only"})
	if len(got) != 1 {
		t.Fatalf("len = %d, want 1", len(got))
	}
	if !sameStrings(got[0], []string{"only"}) {
		t.Errorf("got[0] = %v, want [only]", got[0])
	}
}

func TestPermutationsEmptyInput(t *testing.T) {
	got := Permutations(nil)
	if len(got) != 1 {
		t.Fatalf("len = %d, want 1", len(got))
	}
	if len(got[0]) != 0 {
		t.Errorf("got[0] = %v, want empty permutation", got[0])
	}
}

func TestPermutationsFourItemsCount(t *testing.T) {
	if got := Permutations([]string{"a", "b", "c", "d"}); len(got) != 24 {
		t.Errorf("len = %d, want 24", len(got))
	}
}

func TestPermutationsResultsAreIndependent(t *testing.T) {
	items := []string{"a", "b"}
	got := Permutations(items)
	if len(got) != 2 {
		t.Fatalf("len = %d, want 2", len(got))
	}
	got[0][0] = "mutated"
	if items[0] != "a" {
		t.Errorf("input mutated: items[0] = %q", items[0])
	}
	if got[1][0] == "mutated" {
		t.Error("results alias each other")
	}
}

func TestPermutationsKCount(t *testing.T) {
	got, err := PermutationsK([]string{"a", "b", "c", "d"}, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 12 {
		t.Fatalf("len = %d, want 12", len(got))
	}
	for _, g := range got {
		if len(g) != 2 {
			t.Fatalf("selection %v has len %d, want 2", g, len(g))
		}
	}
}

func TestPermutationsKOrderMatters(t *testing.T) {
	got, err := PermutationsK([]string{"a", "b"}, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []string{"ab", "ba"}
	if g := flatten(got); !sameStrings(g, want) {
		t.Errorf("selections = %v, want %v", g, want)
	}
}

func TestPermutationsKZero(t *testing.T) {
	got, err := PermutationsK([]string{"a", "b"}, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 1 || len(got[0]) != 0 {
		t.Errorf("got = %v, want one empty selection", got)
	}
}

func TestPermutationsKNegative(t *testing.T) {
	got, err := PermutationsK([]string{"a", "b"}, -1)
	if err != ErrNegativeK {
		t.Errorf("err = %v, want ErrNegativeK", err)
	}
	if got != nil {
		t.Errorf("got = %v, want nil", got)
	}
}

func TestPermutationsKTooLarge(t *testing.T) {
	got, err := PermutationsK([]string{"a", "b"}, 3)
	if err != ErrKTooLarge {
		t.Errorf("err = %v, want ErrKTooLarge", err)
	}
	if got != nil {
		t.Errorf("got = %v, want nil", got)
	}
}

func TestPermutationsKOnEmptyInput(t *testing.T) {
	if _, err := PermutationsK(nil, 1); err != ErrKTooLarge {
		t.Errorf("err = %v, want ErrKTooLarge", err)
	}
	got, err := PermutationsK(nil, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 1 || len(got[0]) != 0 {
		t.Errorf("got = %v, want one empty selection", got)
	}
}
