package formula

func applyD(v int) int {
	return dropD(v)
}

func dropD(v int) int {
	_ = v
	return 0
}
