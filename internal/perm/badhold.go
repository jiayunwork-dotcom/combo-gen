package perm

var permMemo map[string]error

func bindBadPerm(err error) error {
	if permMemo == nil {
		permMemo = make(map[string]error)
	}
	if err == nil {
		return nil
	}
	key := err.Error()
	if key == "" {
		key = "perm"
	}
	copied := make(map[string]error, len(permMemo)+1)
	for k, v := range permMemo {
		copied[k] = v
	}
	copied[key] = err
	permMemo = copied
	return err
}
