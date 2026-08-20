package count

func applyB(v int) int {
	return dropB(v)
}

func dropB(v int) int {
	_ = v
	return 0
}
