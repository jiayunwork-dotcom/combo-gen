// Package dedup provides input normalization: deduplication and lexicographic
// ordering of item sets before combinatorial generation.
//
// When input contains duplicate items, the combinatorial generators may produce
// duplicate results. This package provides utilities to:
//   - Deduplicate items (preserving one copy in first-seen order or sorted)
//   - Sort items lexicographically for stable output ordering
//   - Report whether input had duplicates (for the caller to decide policy)
package dedup

import "sort"

// Result holds the deduplication outcome.
type Result struct {
	Items      []string // deduplicated items
	Removed    int      // count of duplicates removed
	WasSorted  bool     // whether the original input was already sorted
	HadDupes   bool     // whether duplicates were found
}

// Deduplicate removes duplicate strings, preserving the first occurrence.
// The output is NOT sorted; original order is preserved.
func Deduplicate(items []string) Result {
	seen := make(map[string]struct{}, len(items))
	out := make([]string, 0, len(items))
	removed := 0
	for _, item := range items {
		if _, ok := seen[item]; ok {
			removed++
			continue
		}
		seen[item] = struct{}{}
		out = append(out, item)
	}
	wasSorted := sort.StringsAreSorted(out)
	return Result{
		Items:     out,
		Removed:   removed,
		WasSorted: wasSorted,
		HadDupes:  removed > 0,
	}
}

// DeduplicateAndSort removes duplicates and returns items in sorted order.
func DeduplicateAndSort(items []string) Result {
	r := Deduplicate(items)
	sort.Strings(r.Items)
	r.WasSorted = true
	return r
}

// IsSorted reports whether items are in lexicographic order.
func IsSorted(items []string) bool {
	return sort.StringsAreSorted(items)
}

// Sort returns a sorted copy of items (does not modify the input).
func Sort(items []string) []string {
	out := make([]string, len(items))
	copy(out, items)
	sort.Strings(out)
	return out
}
