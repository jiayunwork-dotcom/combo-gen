package comb

import "testing"

func TestLexRank(t *testing.T) {
	// C(5,2): [0,1] is rank 0, [0,2] is rank 1, etc.
	if LexRank(5, []int{0, 1}) != 0 {
		t.Error("[0,1] should be rank 0")
	}
	if LexRank(5, []int{3, 4}) != 9 {
		t.Errorf("[3,4] should be rank 9, got %d", LexRank(5, []int{3, 4}))
	}
}

func TestLexUnrank(t *testing.T) {
	indices := LexUnrank(5, 2, 0)
	if indices[0] != 0 || indices[1] != 1 {
		t.Errorf("rank 0 should be [0,1], got %v", indices)
	}
}

func TestLexRankUnrank_RoundTrip(t *testing.T) {
	n, k := 5, 3
	total := binomial(n, k)
	for r := 0; r < total; r++ {
		indices := LexUnrank(n, k, r)
		got := LexRank(n, indices)
		if got != r {
			t.Errorf("round-trip: rank %d → %v → rank %d", r, indices, got)
		}
	}
}

func TestCombinationsFromMultisets(t *testing.T) {
	pools := [][]string{{"a", "b"}, {"c", "d"}}
	combs, err := CombinationsFromMultisets(pools, 2)
	if err != nil {
		t.Fatal(err)
	}
	// C(4,2) = 6
	if len(combs) != 6 {
		t.Errorf("expected 6, got %d", len(combs))
	}
}
