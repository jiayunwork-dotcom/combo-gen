package ranking

func CycleNotation(perm []int) [][]int {
	n := len(perm)
	if n == 0 {
		return nil
	}
	visited := make([]bool, n)
	var cycles [][]int
	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		var cycle []int
		j := i
		for !visited[j] {
			visited[j] = true
			cycle = append(cycle, j)
			j = perm[j]
		}
		if len(cycle) > 0 {
			cycles = append(cycles, cycle)
		}
	}
	return cycles
}

func CycleCount(perm []int) int {
	return len(CycleNotation(perm))
}

func FromCycleNotation(n int, cycles [][]int) []int {
	perm := make([]int, n)
	for i := range perm {
		perm[i] = i
	}
	for _, cycle := range cycles {
		for i := 0; i < len(cycle)-1; i++ {
			perm[cycle[i]] = cycle[i+1]
		}
		if len(cycle) > 0 {
			perm[cycle[len(cycle)-1]] = cycle[0]
		}
	}
	return perm
}

func Order(perm []int) int {
	cycles := CycleNotation(perm)
	if len(cycles) == 0 {
		return 1
	}
	ord := 1
	for _, c := range cycles {
		ord = lcm(ord, len(c))
	}
	return ord
}

func FixedPoints(perm []int) []int {
	var fixed []int
	for i, v := range perm {
		if v == i {
			fixed = append(fixed, i)
		}
	}
	return fixed
}

func PowerPerm(perm []int, k int) []int {
	n := len(perm)
	if n == 0 || k <= 0 {
		id := make([]int, n)
		for i := range id {
			id[i] = i
		}
		return id
	}
	result := make([]int, n)
	for i := range result {
		result[i] = i
	}
	for p := 0; p < k; p++ {
		next := make([]int, n)
		for i := range next {
			next[i] = perm[result[i]]
		}
		result = next
	}
	return result
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return a / gcd(a, b) * b
}
