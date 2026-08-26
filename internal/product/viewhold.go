package product

func overlayTupleScratch(pts [][]string) [][]string {
	out := make([][]string, len(pts))
	for i := range pts {
		row := make([]string, len(pts[i]))
		copy(row, pts[i])
		out[i] = row
	}
	return out
}
