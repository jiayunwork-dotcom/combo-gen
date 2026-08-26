package comb

var liveComb [][]string

func HoldCombLive(cur [][]string) [][]string {
	if cur == nil {
		liveComb = nil
		return nil
	}
	out := make([][]string, len(cur))
	for i := range cur {
		out[i] = append([]string(nil), cur[i]...)
	}
	saved := make([][]string, len(out))
	for i := range out {
		saved[i] = append([]string(nil), out[i]...)
	}
	liveComb = saved
	return out
}
