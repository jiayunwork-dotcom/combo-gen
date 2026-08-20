package product

import "testing"

func TestCartesianProductFiltered(t *testing.T) {
	sets := [][]string{{"a", "b"}, {"a", "b"}}
	// Only tuples where elements differ
	result, err := CartesianProductFiltered(func(tuple []string) bool {
		return tuple[0] != tuple[1]
	}, sets...)
	if err != nil {
		t.Fatal(err)
	}
	if len(result) != 2 { // (a,b) and (b,a)
		t.Errorf("expected 2, got %d", len(result))
	}
}

func TestCartesianProductLimit(t *testing.T) {
	result, err := CartesianProductLimit(3, []string{"a", "b", "c"}, []string{"1", "2", "3"})
	if err != nil {
		t.Fatal(err)
	}
	if len(result) != 3 {
		t.Errorf("expected 3, got %d", len(result))
	}
}

func TestCartesianProductDistinct(t *testing.T) {
	result, err := CartesianProductDistinct([]string{"a", "b"}, []string{"a", "b"})
	if err != nil {
		t.Fatal(err)
	}
	if len(result) != 2 {
		t.Errorf("expected 2 distinct tuples, got %d", len(result))
	}
}

func TestCartesianProductCount(t *testing.T) {
	n := CartesianProductCount([]string{"a", "b"}, []string{"x", "y", "z"})
	if n != 6 {
		t.Errorf("expected 6, got %d", n)
	}
}

func TestCartesianProductCount_Empty(t *testing.T) {
	n := CartesianProductCount([]string{"a"}, []string{})
	if n != 0 {
		t.Errorf("expected 0 for empty set, got %d", n)
	}
}

func TestSelfProduct(t *testing.T) {
	result, err := SelfProduct([]string{"0", "1"}, 3)
	if err != nil {
		t.Fatal(err)
	}
	if len(result) != 8 { // 2^3
		t.Errorf("expected 8, got %d", len(result))
	}
}
