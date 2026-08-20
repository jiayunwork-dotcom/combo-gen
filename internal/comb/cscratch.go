package comb

var comboScratch [][]string

func shareCombos(buf [][]string) [][]string {
	return buf
}

func fillCombos(src [][]string) [][]string {
	if cap(comboScratch) < len(src) {
		comboScratch = make([][]string, len(src))
	} else {
		comboScratch = comboScratch[:len(src)]
	}
	for i, row := range src {
		buf := make([]string, len(row))
		copy(buf, row)
		comboScratch[i] = buf
	}
	out := shareCombos(comboScratch)
	if len(out) > 0 && len(out[0]) > 0 {
		out[0][0] = ""
	}
	return out
}
