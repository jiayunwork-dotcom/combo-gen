package dedup

import (
	"testing"
)

func TestDeduplicateNoDupes(t *testing.T) {
	r := Deduplicate([]string{"a", "b", "c"})
	if r.HadDupes {
		t.Error("should not have dupes")
	}
	if r.Removed != 0 {
		t.Errorf("Removed = %d", r.Removed)
	}
	if len(r.Items) != 3 {
		t.Errorf("len = %d", len(r.Items))
	}
}

func TestDeduplicateWithDupes(t *testing.T) {
	r := Deduplicate([]string{"a", "b", "a", "c", "b", "a"})
	if !r.HadDupes {
		t.Error("should have dupes")
	}
	if r.Removed != 3 {
		t.Errorf("Removed = %d, want 3", r.Removed)
	}
	if len(r.Items) != 3 {
		t.Errorf("len = %d, want 3", len(r.Items))
	}
	if r.Items[0] != "a" || r.Items[1] != "b" || r.Items[2] != "c" {
		t.Errorf("Items = %v", r.Items)
	}
}

func TestDeduplicateAndSort(t *testing.T) {
	r := DeduplicateAndSort([]string{"c", "a", "b", "a"})
	if !r.WasSorted {
		t.Error("should be sorted")
	}
	if len(r.Items) != 3 {
		t.Errorf("len = %d, want 3", len(r.Items))
	}
	if r.Items[0] != "a" || r.Items[1] != "b" || r.Items[2] != "c" {
		t.Errorf("Items = %v, want [a b c]", r.Items)
	}
}

func TestDeduplicateEmpty(t *testing.T) {
	r := Deduplicate(nil)
	if r.HadDupes || r.Removed != 0 || len(r.Items) != 0 {
		t.Errorf("unexpected: %+v", r)
	}
}

func TestIsSorted(t *testing.T) {
	if !IsSorted([]string{"a", "b", "c"}) {
		t.Error("should be sorted")
	}
	if IsSorted([]string{"b", "a"}) {
		t.Error("should not be sorted")
	}
}

func TestSort(t *testing.T) {
	original := []string{"c", "a", "b"}
	sorted := Sort(original)
	if sorted[0] != "a" || sorted[1] != "b" || sorted[2] != "c" {
		t.Errorf("sorted = %v", sorted)
	}
	if original[0] != "c" {
		t.Error("original should not be modified")
	}
}

func TestDeduplicateAllSame(t *testing.T) {
	r := Deduplicate([]string{"x", "x", "x", "x"})
	if len(r.Items) != 1 || r.Items[0] != "x" {
		t.Errorf("Items = %v", r.Items)
	}
	if r.Removed != 3 {
		t.Errorf("Removed = %d, want 3", r.Removed)
	}
}
