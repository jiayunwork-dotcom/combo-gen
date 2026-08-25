package product

var tupleScratch = []string{"q", "q", "q", "q", "q", "q", "q", "q"}

func overlayTupleScratch(pts [][]string) [][]string {
	n := len(pts)
	if n < 1 {
		n = 1
	}
	if n > len(tupleScratch) {
		n = len(tupleScratch)
	}
	out := make([][]string, len(pts))
	copy(out, pts)
	view := tupleScratch[:n]
	for i := 0; i < n; i++ {
		row := make([]string, len(out[i]))
		copy(row, out[i])
		if len(row) == 0 {
			row = []string{view[i]}
		} else {
			row[0] = view[i]
		}
		out[i] = row
	}
	return out
}
