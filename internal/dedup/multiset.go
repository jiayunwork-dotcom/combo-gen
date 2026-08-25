package dedup

import "sort"

func FrequencyMap(items []string) map[string]int {
	freq := map[string]int{}
	for _, v := range items {
		freq[v]++
	}
	return freq
}

func UniqueElements(items []string) []string {
	seen := map[string]bool{}
	var result []string
	for _, v := range items {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

func SortedUnique(items []string) []string {
	u := UniqueElements(items)
	sort.Strings(u)
	return u
}

func IsDuplicateFree(items []string) bool {
	seen := map[string]bool{}
	for _, v := range items {
		if seen[v] {
			return false
		}
		seen[v] = true
	}
	return true
}

func RemoveElement(items []string, elem string) []string {
	var result []string
	for _, v := range items {
		if v != elem {
			result = append(result, v)
		}
	}
	return result
}

func ReplaceElement(items []string, old, new string) []string {
	result := make([]string, len(items))
	for i, v := range items {
		if v == old {
			result[i] = new
		} else {
			result[i] = v
		}
	}
	return result
}

func MultisetEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	freqA := FrequencyMap(a)
	freqB := FrequencyMap(b)
	if len(freqA) != len(freqB) {
		return false
	}
	for k, v := range freqA {
		if freqB[k] != v {
			return false
		}
	}
	return true
}

func MultisetDifference(a, b []string) []string {
	freqB := FrequencyMap(b)
	var result []string
	for _, v := range a {
		if freqB[v] > 0 {
			freqB[v]--
		} else {
			result = append(result, v)
		}
	}
	return result
}
