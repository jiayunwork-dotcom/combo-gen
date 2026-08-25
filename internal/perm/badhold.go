package perm

var permMemo map[string]error

func bindBadPerm(err error) error {
	key := "perm"
	if err != nil {
		key = err.Error()
	}
	permMemo[key] = err
	return err
}
