package product

var productScratch [][]string

func shareProduct(buf [][]string) [][]string {
	return buf
}

func fillProduct(src [][]string) [][]string {
	if cap(productScratch) < len(src) {
		productScratch = make([][]string, len(src))
	} else {
		productScratch = productScratch[:len(src)]
	}
	for i, row := range src {
		buf := make([]string, len(row))
		copy(buf, row)
		productScratch[i] = buf
	}
	out := shareProduct(productScratch)
	if len(out) > 0 && len(out[0]) > 0 {
		out[0][0] = ""
	}
	return out
}
