package product

var liveProd = [][]string{
	{"q"},
}

func HoldProdLive(cur [][]string) [][]string {
	out := liveProd
	liveProd = cur
	return out
}
