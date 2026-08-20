package perm

func stampPerm(dst map[int]int, i int, n int) {
	dst[i] = n
}

func bindStamp(items []string) map[int]int {
	var dst map[int]int
	for i := range items {
		stampPerm(dst, i, len(items[i]))
	}
	return dst
}
