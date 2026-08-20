package ranking

import "testing"

func TestRank_Identity(t *testing.T) {
	// [0,1,2] is the first permutation → rank 0
	r, err := Rank([]int{0, 1, 2})
	if err != nil {
		t.Fatal(err)
	}
	if r != 0 {
		t.Errorf("identity rank expected 0, got %d", r)
	}
}

func TestRank_Last(t *testing.T) {
	// [2,1,0] is the last permutation of 3 → rank 5
	r, err := Rank([]int{2, 1, 0})
	if err != nil {
		t.Fatal(err)
	}
	if r != 5 {
		t.Errorf("last perm rank expected 5, got %d", r)
	}
}

func TestUnrank_Zero(t *testing.T) {
	perm, err := Unrank(3, 0)
	if err != nil {
		t.Fatal(err)
	}
	expected := []int{0, 1, 2}
	for i, v := range expected {
		if perm[i] != v {
			t.Errorf("perm[%d]=%d, want %d", i, perm[i], v)
		}
	}
}

func TestRankUnrank_RoundTrip(t *testing.T) {
	for rank := 0; rank < 24; rank++ { // 4! = 24
		perm, err := Unrank(4, rank)
		if err != nil {
			t.Fatalf("unrank(%d): %v", rank, err)
		}
		got, err := Rank(perm)
		if err != nil {
			t.Fatalf("rank: %v", err)
		}
		if got != rank {
			t.Errorf("round-trip: rank %d → perm %v → rank %d", rank, perm, got)
		}
	}
}

func TestUnrank_TooLarge(t *testing.T) {
	_, err := Unrank(3, 6) // 3! = 6, max rank is 5
	if err != ErrRankTooLarge {
		t.Errorf("expected ErrRankTooLarge, got %v", err)
	}
}

func TestLehmerCode(t *testing.T) {
	code, err := LehmerCode([]int{2, 0, 1})
	if err != nil {
		t.Fatal(err)
	}
	// 2 has 2 elements less after it (0,1); 0 has 0; 1 has 0
	expected := []int{2, 0, 0}
	for i, v := range expected {
		if code[i] != v {
			t.Errorf("lehmer[%d]=%d, want %d", i, code[i], v)
		}
	}
}

func TestInversionCount(t *testing.T) {
	// [2,0,1]: inversions are (2,0), (2,1) → 2
	if InversionCount([]int{2, 0, 1}) != 2 {
		t.Errorf("expected 2 inversions")
	}
}

func TestIsEvenPermutation(t *testing.T) {
	if !IsEvenPermutation([]int{0, 1, 2}) {
		t.Error("identity should be even")
	}
	if IsEvenPermutation([]int{1, 0, 2}) {
		t.Error("single swap should be odd")
	}
}

func TestNextPermutation(t *testing.T) {
	perm := []int{0, 1, 2}
	if !NextPermutation(perm) {
		t.Fatal("should have next")
	}
	expected := []int{0, 2, 1}
	for i, v := range expected {
		if perm[i] != v {
			t.Errorf("next[%d]=%d, want %d", i, perm[i], v)
		}
	}
}

func TestNextPermutation_Last(t *testing.T) {
	perm := []int{2, 1, 0}
	if NextPermutation(perm) {
		t.Error("last permutation should return false")
	}
}
