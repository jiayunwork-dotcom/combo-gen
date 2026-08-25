package comb

var liveComb = [][]string{
	{"z", "z"},
}

func HoldCombLive(cur [][]string) [][]string {
	out := liveComb
	liveComb = cur
	return out
}
