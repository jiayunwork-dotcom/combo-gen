package count

import (
	"testing"

	"combo-gen/internal/comb"
	"combo-gen/internal/perm"
	"combo-gen/internal/product"
)

func TestInvariantCombinationsMatchCount(t *testing.T) {
	items := []string{"a", "b", "c", "d", "e"}
	for k := 0; k <= len(items); k++ {
		got, err := comb.Combinations(items, k)
		if err != nil {
			t.Fatalf("k=%d: %v", k, err)
		}
		want := CombinationCount(len(items), k)
		if len(got) != want {
			t.Errorf("C(%d,%d): enum=%d, formula=%d", len(items), k, len(got), want)
		}
	}
}

func TestInvariantCombWithRepMatchCount(t *testing.T) {
	items := []string{"x", "y", "z"}
	for k := 0; k <= 4; k++ {
		got, err := comb.CombinationsWithRepetition(items, k)
		if err != nil {
			t.Fatalf("k=%d: %v", k, err)
		}
		want := CombinationWithRepCount(len(items), k)
		if len(got) != want {
			t.Errorf("CWR(%d,%d): enum=%d, formula=%d", len(items), k, len(got), want)
		}
	}
}

func TestInvariantPermutationsMatchCount(t *testing.T) {
	for n := 0; n <= 5; n++ {
		items := make([]string, n)
		for i := range items {
			items[i] = string(rune('a' + i))
		}
		got := perm.Permutations(items)
		want := FullPermutationCount(n)
		if len(got) != want {
			t.Errorf("P(%d): enum=%d, formula=%d", n, len(got), want)
		}
	}
}

func TestInvariantPermKMatchCount(t *testing.T) {
	items := []string{"a", "b", "c", "d"}
	for k := 0; k <= len(items); k++ {
		got, err := perm.PermutationsK(items, k)
		if err != nil {
			t.Fatalf("k=%d: %v", k, err)
		}
		want := PermutationCount(len(items), k)
		if len(got) != want {
			t.Errorf("P(%d,%d): enum=%d, formula=%d", len(items), k, len(got), want)
		}
	}
}

func TestInvariantCartesianProductMatchCount(t *testing.T) {
	sets := [][]string{
		{"a", "b"},
		{"x", "y", "z"},
		{"1", "2"},
	}
	got, err := product.CartesianProduct(sets...)
	if err != nil {
		t.Fatal(err)
	}
	want := CartesianProductCount(2, 3, 2)
	if len(got) != want {
		t.Errorf("CartProd: enum=%d, formula=%d", len(got), want)
	}
}

func TestInvariantNoDuplicatesInCombinations(t *testing.T) {
	items := []string{"a", "b", "c", "d"}
	got, _ := comb.Combinations(items, 2)
	seen := map[string]bool{}
	for _, combo := range got {
		key := combo[0] + "," + combo[1]
		if seen[key] {
			t.Errorf("duplicate combination: %v", combo)
		}
		seen[key] = true
	}
}

func TestInvariantNoDuplicatesInPermK(t *testing.T) {
	items := []string{"a", "b", "c"}
	got, _ := perm.PermutationsK(items, 2)
	seen := map[string]bool{}
	for _, p := range got {
		key := p[0] + "," + p[1]
		if seen[key] {
			t.Errorf("duplicate permutation: %v", p)
		}
		seen[key] = true
	}
}

func TestInvariantCombinationsLexOrder(t *testing.T) {
	items := []string{"a", "b", "c", "d"}
	got, _ := comb.Combinations(items, 2)
	for _, combo := range got {
		for i := 1; i < len(combo); i++ {
			pos0, pos1 := -1, -1
			for j, item := range items {
				if item == combo[i-1] {
					pos0 = j
				}
				if item == combo[i] {
					pos1 = j
				}
			}
			if pos0 >= pos1 {
				t.Errorf("combination %v not in input order", combo)
			}
		}
	}
}
