package partition

// CompositionsK returns all ordered compositions of n into exactly k positive parts.
func CompositionsK(n, k int) [][]int {
	if n <= 0 || k <= 0 || k > n {
		return nil
	}
	var result [][]int
	cur := make([]int, 0, k)
	var generate func(remaining, parts int)
	generate = func(remaining, parts int) {
		if parts == k {
			if remaining == 0 {
				out := make([]int, len(cur))
				copy(out, cur)
				result = append(result, out)
			}
			return
		}
		partsLeft := k - parts
		maxP := remaining - partsLeft + 1
		for p := 1; p <= maxP; p++ {
			cur = append(cur, p)
			generate(remaining-p, parts+1)
			cur = cur[:len(cur)-1]
		}
	}
	generate(n, 0)
	return result
}

// WeakCompositions returns all weak compositions of n into k non-negative parts.
// (Unlike compositions, parts can be 0.)
func WeakCompositions(n, k int) [][]int {
	if k <= 0 {
		return nil
	}
	if n == 0 {
		zeros := make([]int, k)
		return [][]int{zeros}
	}
	if n < 0 {
		return nil
	}
	var result [][]int
	cur := make([]int, 0, k)
	var generate func(remaining, parts int)
	generate = func(remaining, parts int) {
		if parts == k {
			if remaining == 0 {
				out := make([]int, len(cur))
				copy(out, cur)
				result = append(result, out)
			}
			return
		}
		for p := 0; p <= remaining; p++ {
			cur = append(cur, p)
			generate(remaining-p, parts+1)
			cur = cur[:len(cur)-1]
		}
	}
	generate(n, 0)
	return result
}

// ConjugatePartition returns the conjugate (transpose) of an integer partition.
// The conjugate's i-th part equals the number of parts >= i in the original.
func ConjugatePartition(parts []int) []int {
	if len(parts) == 0 {
		return nil
	}
	maxPart := parts[0]
	conj := make([]int, maxPart)
	for _, p := range parts {
		for i := 0; i < p; i++ {
			conj[i]++
		}
	}
	return conj
}

// IsSelfConjugate checks if a partition equals its own conjugate.
func IsSelfConjugate(parts []int) bool {
	conj := ConjugatePartition(parts)
	if len(conj) != len(parts) {
		return false
	}
	for i := range parts {
		if parts[i] != conj[i] {
			return false
		}
	}
	return true
}

// DistinctPartitions returns all partitions of n into distinct parts.
func DistinctPartitions(n int) [][]int {
	if n <= 0 {
		return nil
	}
	var result [][]int
	cur := make([]int, 0)
	var generate func(remaining, maxPart int)
	generate = func(remaining, maxPart int) {
		if remaining == 0 {
			out := make([]int, len(cur))
			copy(out, cur)
			result = append(result, out)
			return
		}
		for p := min(remaining, maxPart); p >= 1; p-- {
			// Ensure distinct: p must be less than the previous part
			if len(cur) > 0 && p >= cur[len(cur)-1] {
				continue
			}
			cur = append(cur, p)
			generate(remaining-p, p-1)
			cur = cur[:len(cur)-1]
		}
		// Also try starting fresh if cur is empty
		if len(cur) == 0 {
			for p := min(remaining, maxPart); p >= 1; p-- {
				cur = append(cur, p)
				generate(remaining-p, p-1)
				cur = cur[:len(cur)-1]
			}
		}
	}
	// Simpler approach: generate partitions where each part is strictly decreasing
	result = nil
	cur = cur[:0]
	var gen2 func(remaining, maxPart int)
	gen2 = func(remaining, maxPart int) {
		if remaining == 0 {
			out := make([]int, len(cur))
			copy(out, cur)
			result = append(result, out)
			return
		}
		limit := maxPart
		if remaining < limit {
			limit = remaining
		}
		for p := limit; p >= 1; p-- {
			cur = append(cur, p)
			gen2(remaining-p, p-1) // next part must be strictly less
			cur = cur[:len(cur)-1]
		}
	}
	gen2(n, n)
	return result
}
