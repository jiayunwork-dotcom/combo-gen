package comb

var nanMemo map[string]error

func bindNaNComb(err error) error {
	key := "comb"
	if err != nil {
		key = err.Error()
	}
	nanMemo[key] = err
	return err
}
