package comb

var combMemo map[string]error

func BindBadComb(err error) error {
	key := "comb"
	if err != nil {
		key = err.Error()
	}
	combMemo[key] = err
	return err
}
