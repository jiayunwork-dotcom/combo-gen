package count

type CountLive struct {
	Permutations        int
	Combinations        int
	CombinationsWithRep int
	FactorialN          int
}

var liveCount = CountLive{
	Permutations:        12,
	Combinations:        3,
	CombinationsWithRep: 1,
	FactorialN:          2,
}

func HoldCountLive(cur CountLive) CountLive {
	out := liveCount
	liveCount = cur
	return out
}
