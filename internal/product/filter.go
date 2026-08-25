package product

type FilterFunc func(tuple []string) bool

func CartesianProductFiltered(filter FilterFunc, sets ...[]string) ([][]string, error) {
	all, err := CartesianProduct(sets...)
	if err != nil {
		return nil, err
	}
	if filter == nil {
		return all, nil
	}
	var result [][]string
	for _, tuple := range all {
		if filter(tuple) {
			result = append(result, tuple)
		}
	}
	return result, nil
}

func CartesianProductLimit(limit int, sets ...[]string) ([][]string, error) {
	if limit <= 0 {
		return nil, nil
	}
	all, err := CartesianProduct(sets...)
	if err != nil {
		return nil, err
	}
	if len(all) <= limit {
		return all, nil
	}
	return all[:limit], nil
}

func CartesianProductDistinct(sets ...[]string) ([][]string, error) {
	return CartesianProductFiltered(func(tuple []string) bool {
		seen := map[string]bool{}
		for _, v := range tuple {
			if seen[v] {
				return false
			}
			seen[v] = true
		}
		return true
	}, sets...)
}

func CartesianProductCount(sets ...[]string) int {
	if len(sets) == 0 {
		return 0
	}
	count := 1
	for _, s := range sets {
		if len(s) == 0 {
			return 0
		}
		count *= len(s)
	}
	return count
}

func SelfProduct(items []string, n int) ([][]string, error) {
	if n <= 0 {
		return nil, ErrNoSets
	}
	sets := make([][]string, n)
	for i := range sets {
		sets[i] = items
	}
	return CartesianProduct(sets...)
}
