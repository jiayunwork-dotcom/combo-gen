package subset

var liveSet = [][]string{
	{"old"},
	{"old"},
}

func HoldSetLive(cur [][]string) [][]string {
	out := liveSet
	liveSet = cur
	return out
}
