package comb

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

func TestCombinationsCount(t *testing.T) {
	got, err := Combinations([]string{"a", "b", "c", "d", "e"}, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 10 {
		t.Fatalf("len = %d, want 10", len(got))
	}
	for _, g := range got {
		if len(g) != 2 {
			t.Fatalf("combination %v has len %d, want 2", g, len(g))
		}
	}
}

func TestCombinationsContent(t *testing.T) {
	got, err := Combinations([]string{"a", "b", "c"}, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []string{"ab", "ac", "bc"}
	if g := flatten(got); !sameStrings(g, want) {
		t.Errorf("combinations = %v, want %v", g, want)
	}
}

func TestCombinationsFullLength(t *testing.T) {
	got, err := Combinations([]string{"a", "b", "c"}, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 1 {
		t.Fatalf("len = %d, want 1", len(got))
	}
	if !sameStrings(got[0], []string{"a", "b", "c"}) {
		t.Errorf("got[0] = %v, want [a b c]", got[0])
	}
}

func TestCombinationsZeroK(t *testing.T) {
	got, err := Combinations([]string{"a", "b"}, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 1 || len(got[0]) != 0 {
		t.Errorf("got = %v, want one empty combination", got)
	}
}

func TestCombinationsNegativeK(t *testing.T) {
	got, err := Combinations([]string{"a", "b"}, -2)
	if err != ErrNegativeK {
		t.Errorf("err = %v, want ErrNegativeK", err)
	}
	if got != nil {
		t.Errorf("got = %v, want nil", got)
	}
}

func TestCombinationsKTooLarge(t *testing.T) {
	got, err := Combinations([]string{"a", "b"}, 5)
	if err != ErrKTooLarge {
		t.Errorf("err = %v, want ErrKTooLarge", err)
	}
	if got != nil {
		t.Errorf("got = %v, want nil", got)
	}
}

func TestCombinationsResultsAreIndependent(t *testing.T) {
	items := []string{"a", "b", "c"}
	got, err := Combinations(items, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got[0][0] = "mutated"
	if items[0] != "a" {
		t.Errorf("input mutated: items[0] = %q", items[0])
	}
	if got[1][0] == "mutated" {
		t.Error("results alias each other")
	}
}

func TestCombinationsWithRepetitionCount(t *testing.T) {
	got, err := CombinationsWithRepetition([]string{"a", "b", "c"}, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []string{"aa", "ab", "ac", "bb", "bc", "cc"}
	if g := flatten(got); !sameStrings(g, want) {
		t.Errorf("combinations = %v, want %v", g, want)
	}
}

func TestCombinationsWithRepetitionAllowsKAboveN(t *testing.T) {
	got, err := CombinationsWithRepetition([]string{"a", "b"}, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []string{"aaa", "aab", "abb", "bbb"}
	if g := flatten(got); !sameStrings(g, want) {
		t.Errorf("combinations = %v, want %v", g, want)
	}
}

func TestCombinationsWithRepetitionZeroK(t *testing.T) {
	got, err := CombinationsWithRepetition(nil, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 1 || len(got[0]) != 0 {
		t.Errorf("got = %v, want one empty combination", got)
	}
}

func TestCombinationsWithRepetitionNoItems(t *testing.T) {
	got, err := CombinationsWithRepetition(nil, 2)
	if err != ErrNoItems {
		t.Errorf("err = %v, want ErrNoItems", err)
	}
	if got != nil {
		t.Errorf("got = %v, want nil", got)
	}
}

func TestCombinationsWithRepetitionNegativeK(t *testing.T) {
	if _, err := CombinationsWithRepetition([]string{"a"}, -1); err != ErrNegativeK {
		t.Errorf("err = %v, want ErrNegativeK", err)
	}
}
