package partition

import "testing"

func TestSetPartitions_Three(t *testing.T) {
	parts := SetPartitions([]string{"a", "b", "c"})
	if len(parts) != 5 {
		t.Errorf("expected 5 set partitions, got %d", len(parts))
	}
}

func TestSetPartitions_Empty(t *testing.T) {
	parts := SetPartitions(nil)
	if len(parts) != 1 {
		t.Error("empty set has 1 partition (the empty partition)")
	}
}

func TestIntegerPartitions_Four(t *testing.T) {
	parts := IntegerPartitions(4)
	if len(parts) != 5 {
		t.Errorf("expected 5 partitions of 4, got %d", len(parts))
	}
	for _, p := range parts {
		sum := 0
		for _, v := range p {
			sum += v
		}
		if sum != 4 {
			t.Errorf("partition %v does not sum to 4", p)
		}
	}
}

func TestIntegerPartitions_One(t *testing.T) {
	parts := IntegerPartitions(1)
	if len(parts) != 1 || parts[0][0] != 1 {
		t.Errorf("p(1) should be [[1]], got %v", parts)
	}
}

func TestIntegerPartitionsK(t *testing.T) {
	parts := IntegerPartitionsK(5, 2)
	if len(parts) != 2 {
		t.Errorf("expected 2 partitions of 5 into 2 parts, got %d", len(parts))
	}
}

func TestPartitionCount(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 1}, {1, 1}, {2, 2}, {3, 3}, {4, 5}, {5, 7}, {10, 42},
	}
	for _, tt := range tests {
		got := PartitionCount(tt.n)
		if got != tt.want {
			t.Errorf("p(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestCompositions_Three(t *testing.T) {
	comps := Compositions(3)
	if len(comps) != 4 {
		t.Errorf("expected 4 compositions of 3, got %d", len(comps))
	}
}

func TestCompositions_Zero(t *testing.T) {
	if Compositions(0) != nil {
		t.Error("compositions of 0 should be nil")
	}
}
