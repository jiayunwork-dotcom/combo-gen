package iter

type PermIter struct {
	items []string
	c     []int
	i     int
	first bool
	done  bool
}

func NewPermIter(items []string) *PermIter {
	cp := make([]string, len(items))
	copy(cp, items)
	return &PermIter{
		items: cp,
		c:     make([]int, len(items)),
		i:     0,
		first: true,
		done:  false,
	}
}

func (it *PermIter) Next() []string {
	if it.done {
		return nil
	}
	if it.first {
		it.first = false
		out := make([]string, len(it.items))
		copy(out, it.items)
		return out
	}
	n := len(it.items)
	for it.i < n {
		if it.c[it.i] < it.i {
			if it.i%2 == 0 {
				it.items[0], it.items[it.i] = it.items[it.i], it.items[0]
			} else {
				it.items[it.c[it.i]], it.items[it.i] = it.items[it.i], it.items[it.c[it.i]]
			}
			it.c[it.i]++
			it.i = 0
			out := make([]string, n)
			copy(out, it.items)
			return out
		}
		it.c[it.i] = 0
		it.i++
	}
	it.done = true
	return nil
}

func (it *PermIter) Reset(items []string) {
	cp := make([]string, len(items))
	copy(cp, items)
	it.items = cp
	it.c = make([]int, len(items))
	it.i = 0
	it.first = true
	it.done = false
}

type CombIter struct {
	items []string
	n, k  int
	idx   []int
	first bool
	done  bool
}

func NewCombIter(items []string, k int) *CombIter {
	n := len(items)
	if k < 0 || k > n {
		return &CombIter{done: true}
	}
	idx := make([]int, k)
	for i := range idx {
		idx[i] = i
	}
	return &CombIter{items: items, n: n, k: k, idx: idx, first: true}
}

func (it *CombIter) Next() []string {
	if it.done {
		return nil
	}
	if it.first {
		it.first = false
		return it.current()
	}
	i := it.k - 1
	for i >= 0 && it.idx[i] == it.n-it.k+i {
		i--
	}
	if i < 0 {
		it.done = true
		return nil
	}
	it.idx[i]++
	for j := i + 1; j < it.k; j++ {
		it.idx[j] = it.idx[j-1] + 1
	}
	return it.current()
}

func (it *CombIter) current() []string {
	out := make([]string, it.k)
	for i, idx := range it.idx {
		out[i] = it.items[idx]
	}
	return out
}

type ProductIter struct {
	sets    [][]string
	indices []int
	first   bool
	done    bool
}

func NewProductIter(sets ...[]string) *ProductIter {
	if len(sets) == 0 {
		return &ProductIter{done: true}
	}
	for _, s := range sets {
		if len(s) == 0 {
			return &ProductIter{done: true}
		}
	}
	return &ProductIter{
		sets:    sets,
		indices: make([]int, len(sets)),
		first:   true,
	}
}

func (it *ProductIter) Next() []string {
	if it.done {
		return nil
	}
	if it.first {
		it.first = false
		return it.current()
	}
	m := len(it.sets)
	for i := m - 1; i >= 0; i-- {
		it.indices[i]++
		if it.indices[i] < len(it.sets[i]) {
			return it.current()
		}
		it.indices[i] = 0
	}
	it.done = true
	return nil
}

func (it *ProductIter) current() []string {
	out := make([]string, len(it.sets))
	for i, idx := range it.indices {
		out[i] = it.sets[i][idx]
	}
	return out
}

func CountIter(iter interface{ Next() []string }) int {
	n := 0
	for iter.Next() != nil {
		n++
	}
	return n
}
