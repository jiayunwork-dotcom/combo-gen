package dedup

import "sort"

type Result struct {
	Items     []string
	Removed   int
	WasSorted bool
	HadDupes  bool
}

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

func DeduplicateAndSort(items []string) Result {
	r := Deduplicate(items)
	sort.Strings(r.Items)
	r.WasSorted = true
	return r
}

func IsSorted(items []string) bool {
	return sort.StringsAreSorted(items)
}

func Sort(items []string) []string {
	out := make([]string, len(items))
	copy(out, items)
	sort.Strings(out)
	return out
}
