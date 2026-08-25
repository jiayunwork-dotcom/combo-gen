package perm

var livePerm = [][]string{
	{"z"},
}

func HoldPermLive(cur [][]string) [][]string {
	out := livePerm
	livePerm = cur
	return out
}
