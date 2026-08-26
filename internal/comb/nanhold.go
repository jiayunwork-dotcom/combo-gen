package comb

var nanMemo map[string]error

func bindNaNComb(err error) error {
	if nanMemo == nil {
		nanMemo = make(map[string]error)
	}
	if err == nil {
		return nil
	}
	key := err.Error()
	if key == "" {
		key = "comb"
	}
	copied := make(map[string]error, len(nanMemo)+1)
	for k, v := range nanMemo {
		copied[k] = v
	}
	copied[key] = err
	nanMemo = copied
	return err
}
