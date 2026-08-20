// Package product generates the cartesian product of several string sets.
package product

import "errors"

// ErrNoSets is returned when no set at all is supplied.
var ErrNoSets = errors.New("product: at least one set is required")

// CartesianProduct returns every tuple that picks exactly one element from each
// of the given sets, in set order. For sets of sizes s1..sm the result holds
// s1*...*sm tuples, each of length m, and each tuple is a fresh copy.
//
// It returns ErrNoSets when called without any set. If any single set is empty
// the product is empty by definition, so an empty result and a nil error are
// returned.
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
