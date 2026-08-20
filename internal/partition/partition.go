// Package partition generates set partitions and integer partitions.
package partition

// SetPartitions returns all ways to partition items into non-empty subsets.
// Each partition is a slice of blocks; each block is a slice of strings.
// Bell number B(n) gives the total count.
func SetPartitions(items []string) [][][]string {
	n := len(items)
	if n == 0 {
		return [][][]string{{}}
	}
	var result [][][]string
	// assignment[i] = which block index item i belongs to
	assignment := make([]int, n)
	// maxUsed tracks the maximum block index assigned so far
	var generate func(pos, maxUsed int)
	generate = func(pos, maxUsed int) {
		if pos == n {
			// Build partition from assignment
			blockMap := map[int][]string{}
			for i, b := range assignment {
				blockMap[b] = append(blockMap[b], items[i])
			}
			part := make([][]string, maxUsed+1)
			for b := 0; b <= maxUsed; b++ {
				part[b] = blockMap[b]
			}
			result = append(result, part)
			return
		}
		// Item at pos can go into any existing block 0..maxUsed or a new block maxUsed+1
		for b := 0; b <= maxUsed+1; b++ {
			assignment[pos] = b
			newMax := maxUsed
			if b > maxUsed {
				newMax = b
			}
			generate(pos+1, newMax)
		}
	}
	generate(0, -1)
	return result
}

// IntegerPartitions returns all ways to write n as a sum of positive integers
// in non-increasing order. E.g., IntegerPartitions(4) returns:
// [4], [3,1], [2,2], [2,1,1], [1,1,1,1].
func IntegerPartitions(n int) [][]int {
	if n <= 0 {
		return nil
	}
	var result [][]int
	cur := make([]int, 0, n)
	var generate func(remaining, maxPart int)
	generate = func(remaining, maxPart int) {
		if remaining == 0 {
			out := make([]int, len(cur))
			copy(out, cur)
			result = append(result, out)
			return
		}
		for p := min(remaining, maxPart); p >= 1; p-- {
			cur = append(cur, p)
			generate(remaining-p, p)
			cur = cur[:len(cur)-1]
		}
	}
	generate(n, n)
	return result
}

// IntegerPartitionsK returns all partitions of n into exactly k parts.
func IntegerPartitionsK(n, k int) [][]int {
	if n <= 0 || k <= 0 || k > n {
		return nil
	}
	var result [][]int
	cur := make([]int, 0, k)
	var generate func(remaining, maxPart, parts int)
	generate = func(remaining, maxPart, parts int) {
		if parts == k {
			if remaining == 0 {
				out := make([]int, len(cur))
				copy(out, cur)
				result = append(result, out)
			}
			return
		}
		partsLeft := k - parts
		for p := min(remaining-partsLeft+1, maxPart); p >= 1; p-- {
			cur = append(cur, p)
			generate(remaining-p, p, parts+1)
			cur = cur[:len(cur)-1]
		}
	}
	generate(n, n, 0)
	return result
}

// PartitionCount returns the number of integer partitions of n (partition function p(n)).
func PartitionCount(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	// Dynamic programming
	dp := make([]int, n+1)
	dp[0] = 1
	for k := 1; k <= n; k++ {
		for i := k; i <= n; i++ {
			dp[i] += dp[i-k]
		}
	}
	return dp[n]
}

// Compositions returns all ordered compositions of n into positive parts.
// Unlike partitions, order matters: [2,1] and [1,2] are different compositions of 3.
func Compositions(n int) [][]int {
	if n <= 0 {
		return nil
	}
	var result [][]int
	cur := make([]int, 0, n)
	var generate func(remaining int)
	generate = func(remaining int) {
		if remaining == 0 {
			out := make([]int, len(cur))
			copy(out, cur)
			result = append(result, out)
			return
		}
		for p := 1; p <= remaining; p++ {
			cur = append(cur, p)
			generate(remaining - p)
			cur = cur[:len(cur)-1]
		}
	}
	generate(n)
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
