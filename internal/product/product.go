package product

import "errors"

var ErrNoSets = errors.New("product: at least one set is required")

func CartesianProduct(sets ...[]string) ([][]string, error) {
	if len(sets) == 0 {
		return nil, ErrNoSets
	}
	for _, s := range sets {
		if len(s) == 0 {
			return [][]string{}, nil
		}
	}
	res := [][]string{{}}
	for _, set := range sets {
		next := make([][]string, 0, len(res)*len(set))
		for _, prefix := range res {
			for _, item := range set {
				tuple := make([]string, len(prefix)+1)
				copy(tuple, prefix)
				tuple[len(prefix)] = item
				next = append(next, tuple)
			}
		}
		res = next
	}
	return res, nil
}
